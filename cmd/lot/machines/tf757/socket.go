// 每台机器建立socket连接
package tf757

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/pieterclaerhout/go-log"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
)

// 请求设备状态指令
var RealTimeDataRequestCommandBase = []byte{3, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// 请求周期
var IntervalSendDuation time.Duration = 500

type Tf757Socket struct {
	// 设备信息
	dinfo model.DeviceInfo

	// 数据传输socket
	Conn net.Conn

	// 控制指令发送队列
	SendChan chan []byte

	// 定时请求状态队列
	IntervalRequestChan chan struct{}

	// 标记是否有下发复位指令
	ResetCmdFlag bool

	// ctx 连接状态
	StopChan chan struct{}

	// 上次接收的时间
	LastReceiveTime time.Time

	// 上次上报报警
	LastUpAlarmTime time.Time

	// 上次状态信息
	// StatusArray [50]int

	// 初始状态
	InitFlag bool

	// 操作锁
	Olock sync.Mutex

	// 排产模型缓存
	SchedulingExecutionModelCache model.SchedulingExecutionModel

	// 设备状态缓存
	DeviceStatusCache int

	// 工单结束
	IsOrderOver bool

	// 定时请求bytes
	RealTimeDataRequestBytes []byte
}

func NewSock(info model.DeviceInfo) (s *Tf757Socket, err error) {
	if _, ok := Tf757Manager.ConnectingMap.Load(info.StandNo); ok {
		return
	}
	Tf757Manager.ConnectingMap.Store(info.StandNo, true)
	defer Tf757Manager.ConnectingMap.Delete(info.StandNo)
	conn, err := registerConn(info)
	if err != nil {
		svc.Warnf("天富设备 %s连接失败:%v", info.IpAddress, err.Error())
		return
	}
	s = &Tf757Socket{
		Conn:                conn,
		dinfo:               info,
		SendChan:            make(chan []byte, 10),
		IntervalRequestChan: make(chan struct{}),
		StopChan:            make(chan struct{}),
		LastReceiveTime:     time.Now(),
		InitFlag:            true,
		Olock:               sync.Mutex{},
		IsOrderOver:         true,
	}
	Tf757Manager.AddSock(s)
	s.InitRealTimeDataRequestCommand()
	// 注册接收回调goroutine
	go s.registerReceiveHandler()
	// 注册检测连接状态goroutine
	go s.registerCheckConn()
	// 注册定时请求状态goroutine
	go s.registerIntervalRequest()
	// 注册发送消息goroutine
	go s.registerWriteMessage()

	svc.Infof("天富 建立连接 %s->%s DEVICE:%s StandNo:%d!", s.Conn.LocalAddr(), s.dinfo.IpAddress, s.dinfo.DeviceId, s.dinfo.StandNo)
	return
}

func (h *Tf757Socket) InitRealTimeDataRequestCommand() {
	buf := []byte{byte(h.dinfo.StandNo)}
	buf = append(buf, RealTimeDataRequestCommandBase...)
	h.RealTimeDataRequestBytes = buf
}

func (h *Tf757Socket) Info(msg interface{}) {
	svc.Info(fmt.Sprintf("Model:%s|StandNo:%d|%s", h.dinfo.DeviceName, h.dinfo.StandNo, msg))
}

func registerConn(dinfo model.DeviceInfo) (conn net.Conn, err error) {
	if dinfo.IpAddress == "" {
		return nil, errors.New("ip地址为空")
	}
	port := "3000"
	if dinfo.Port != "" {
		port = dinfo.Port
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", dinfo.IpAddress, port))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s下位机IP（%s）解析失败:%s", dinfo.DeviceId, dinfo.IpAddress, err.Error()))
	}
	// 本地多网卡时指定ip
	localTcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:", svc.Service.Config.TfConfig.Address))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("配置的天富网卡IP（%s）解析失败:%s", svc.Service.Config.TfConfig.Address, err.Error()))
	}
	dialer := net.Dialer{Timeout: 10 * time.Second, LocalAddr: localTcpAddr}
	for i := 0; i < 1; i++ {
		if i != 0 {
			svc.Warnf("%d %v:%v次重连。。。 。。。", dinfo.StandNo, dinfo.DeviceId, i)
		}
		conn, err = dialer.Dial("tcp", tcpAddr.String())
		if err != nil {
			// svc.Warnf("天富设备deviceID[%v]连接失败[%d]:%v", dinfo.DeviceId, dinfo.StandNo, err.Error())
			continue
		} else {
			return conn, nil
		}
	}
	return
}

func (s *Tf757Socket) registerReceiveHandler() {
	var buf [1024]byte
	for {
		select {
		case <-s.StopChan:
			log.Debug(s.dinfo.DeviceId, "--registerReceiveHandler breaked!")
			return
		default:
			n, err := s.Conn.Read(buf[0:])
			if err != nil {
				svc.Warnf("%v-接收错误:%v", s.dinfo.DeviceId, err)
				s.CloseConn("registerReceiveHandler")
				return
			} else {
				s.LastReceiveTime = time.Now()
				// svc.Debugf("从设备接收实时状态数据%d字节", buf[:n])
				// svc.Debugf("从设备接收实时状态数据%d字节", n)
				s.AfterReceive(buf[:n])
				log.Debug("接收到实时状态数据！")
				s.InitFlag = false
			}
		}
	}
}

func (s *Tf757Socket) registerCheckConn() {
	for {
		select {
		case <-s.StopChan:
			log.Debug(s.dinfo.DeviceId, "--registerCheckConn breaked!")
			return
		default:
			if time.Since(s.LastReceiveTime) > 20*time.Second {
				log.Debug(s.dinfo.DeviceId, "--registerCheckConn 20s没有收到下位机消息、将断开连接")
				s.CloseConn("registerCheckConn")
				return
			}
		}
		<-time.After(5 * time.Second)
	}
}
func (s *Tf757Socket) registerIntervalRequest() {
	if svc.Service.Config.TfConfig.IntervalSendDuration > 0 {
		IntervalSendDuation = time.Duration(svc.Service.Config.TfConfig.IntervalSendDuration)
	}
	for {
		select {
		case <-s.StopChan:
			log.Debug(s.dinfo.DeviceId, "--registerIntervalRequest breaked!")
			return
		default:
			s.IntervalRequestChan <- struct{}{}
			// log.Debug("发送请求状态")
		}
		<-time.After(IntervalSendDuation * time.Millisecond)
	}
}

func (s *Tf757Socket) registerWriteMessage() {
	for {
		select {
		case <-s.StopChan:
			log.Debug(s.dinfo.DeviceId, "--registerWriteMessage breaked!")
			return
		default:
			// 优先发送控制指令、其次发送请求状态
			select {
			case buf := <-s.SendChan:
				log.Debug("开始传输控制指令")
				s.SendMessage(buf)
			default:
				select {
				case <-s.IntervalRequestChan:
					s.SendMessage(s.RealTimeDataRequestBytes)
				default:
				}
			}
		}
		<-time.After(100 * time.Millisecond)
	}
}

func (s *Tf757Socket) CloseConn(args ...interface{}) error {
	if len(args) > 0 {
		svc.Infof("Call CloseConn From : %s", args[0].(string))
	}
	defer func() {
		if err := recover(); err != nil {
			svc.Warnf("CloseConn err :%v", err)
		}
	}()
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
	svc.Infof("天富关闭连接[%d]%s", s.dinfo.StandNo, s.dinfo.IpAddress)
	close(s.StopChan)
	close(s.SendChan)
	s.Conn.Close()
	Tf757Manager.DeleteSock(int64(s.dinfo.StandNo))
	return nil
}

func (s *Tf757Socket) EnMessageQueue(buf []byte) error {
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

func (s *Tf757Socket) SendMessage(buf []byte) {
	// 发送前统一添加校验位
	buf = append(buf, PrePacketByStdCRC(buf)...)
	if bytes.Compare(buf[1:4], []byte{3, 0, 1}) != 0 {
		fmt.Println("发送buf", buf)
	}
	if _, err := s.Conn.Write(buf); err != nil {
		svc.Warnf("天富 SendMessage err :%v", err)
	}
}
