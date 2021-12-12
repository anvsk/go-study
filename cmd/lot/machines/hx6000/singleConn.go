// 此文件用于工厂只有一个设备串口的情况
// 仅维护一个连接，用于所有机器通信
package hx6000

import (
	"errors"
	"sync"
	"time"

	"github.com/tarm/serial"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
)

// 请求设备状态指令
var HxRealTimeDataRequestCommand = []byte{0x05, 0x44, 0x43, 0x52, 0x5A, 0x31, 0x36, 0x39, 0x46, 0x04}
var MinPackageLength = 8

var MaxPackageLength = 62

// 请求周期
var IntervalSendDuation time.Duration = 1000

type HxSingleConn struct {

	// 数据传输socket
	Conn *serial.Port

	// 控制指令发送队列
	SendChan chan []byte

	// 请求设备状态发送队列
	IntervalRequestChan chan struct{}

	// ctx 连接状态
	StopChan chan struct{}

	// 上次接收的时间
	LastReceiveTime time.Time

	// 初始状态
	InitFlag bool

	// 操作锁
	Olock sync.Mutex

	// 接收字节流缓存
	Buffer *HxBuffer

	// 当前正在下发工艺的设备站号
	CurCraftStanNo int
}

func NewSingleConn() (s *HxSingleConn, err error) {
	deviceName := svc.Service.Config.HxConfig.DevFile
	if deviceName == "" {
		err = errors.New("设备地址为空，请在配置文件[hx-dev_file]配置正确的设备地址")
		return
	}
	options := &serial.Config{
		Name:        deviceName,
		Baud:        38400,
		ReadTimeout: time.Second * 20,
		Size:        byte(8),
		StopBits:    serial.Stop1,
		Parity:      serial.ParityNone,
	}
	conn, err := serial.OpenPort(options)
	if err != nil {
		return
	}
	s = &HxSingleConn{
		Conn:                conn,
		Buffer:              NewHxBuffer(),
		SendChan:            make(chan []byte, 30),
		IntervalRequestChan: make(chan struct{}, 100),
		StopChan:            make(chan struct{}),
		LastReceiveTime:     time.Now(),
		InitFlag:            true,
		Olock:               sync.Mutex{},
		CurCraftStanNo:      0,
	}
	// // 注册接收回调goroutine
	go s.registerReceiveHandler()
	// // 注册发送消息goroutine
	go s.registerWriteMessage()

	go s.registerIntervalRequest()

	return
}

func (s *HxSingleConn) registerIntervalRequest() {
	if svc.Service.Config.HxConfig.IntervalSendDuration > 0 {
		IntervalSendDuation = time.Duration(svc.Service.Config.HxConfig.IntervalSendDuration)
	}
	for {
		select {
		case <-s.StopChan:
			return
		default:
			s.IntervalRequestChan <- struct{}{}
		}
		<-time.After(IntervalSendDuation * time.Millisecond)
	}
}

func (s *HxSingleConn) registerReceiveHandler() {
	for {
		select {
		case <-s.StopChan:
			svc.Debug("--registerReceiveHandler breaked!")
			return
		default:
			n, err := s.Conn.Read(s.Buffer.Buffer[s.Buffer.ValidLength:])
			if err != nil {
				s.CloseConn("registerReceiveHandler")
				return
			}
			s.Buffer.ValidLength += n
			s.LastReceiveTime = time.Now()
			// svc.Debugf("从设备接收实时状态数据%d字节", n)
			if s.Buffer.ValidLength < MinPackageLength {
				continue
			}
			s.AfterReceivePortBytes()
			s.InitFlag = false
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (s *HxSingleConn) registerWriteMessage() {
	for {
		select {
		case <-s.StopChan:
			svc.Debug("--registerWriteMessage breaked!")
			return
		default:
			// 优先发送控制指令、其次发送请求状态
			select {
			case buf := <-s.SendChan:
				// log.Debug(" 开始传输控制指令%v", buf)
				s.SendMessage(buf)
				// 不停顿会无效
				time.Sleep(200 * time.Millisecond)
			default:
				select {
				case <-s.IntervalRequestChan:
					s.SendMessage(HxRealTimeDataRequestCommand)
				default:
				}
			}
		}
	}
}

func (s *HxSingleConn) CloseConn(args ...interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			svc.Warnf("CloseConn err :%v", err)
		}
	}()
	if len(args) > 0 {
		svc.Infof("Call CloseConn From : %s", args[0].(string))
	}
	if s.Conn == nil {
		return nil
	}
	s.Olock.Lock()
	defer s.Olock.Unlock()
	// 已经关闭则跳过、重复关闭会panic
	select {
	case <-s.StopChan:
		return nil
	default:
	}
	svc.Infof("关闭连接[%d]%s")
	close(s.StopChan)
	close(s.SendChan)
	s.Conn.Close()
	return nil
}

func (s *HxSingleConn) EnMessageQueue(buf []byte) error {
	// 防止关闭管道后写入panic
	defer func() {
		if err := recover(); err != nil {
			svc.Warnf("send_beats err :%v", err)
		}
	}()
	select {
	case <-s.StopChan:
		return nil
	default:
		s.SendChan <- buf
	}
	return nil
}

func (s *HxSingleConn) SendMessage(buf []byte) {
	if _, err := s.Conn.Write(buf); err != nil {
		svc.Warnf(" SendMessage err :%v", err)
	}
}

// 发送状态获取包
func (s *HxSingleConn) SendStatusGet() {
	select {
	case <-s.StopChan:
		return
	default:
		s.IntervalRequestChan <- struct{}{}
	}
}
