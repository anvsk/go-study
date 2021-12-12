// 实现通用接口方法
package tf757

import (
	"fmt"
	"sync"
	"time"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
	"go.uber.org/zap"
)

var Tf757Manager *Tf757
var ReConnectGoroutineStart sync.Once

type Tf757 struct {
	deviceList    []model.DeviceInfo
	dmap          map[int64]*Tf757Socket
	logger        *zap.Logger
	modelName     string
	dlock         sync.Mutex
	ConnectingMap sync.Map
}

func InitTF757Manager() *Tf757 {
	Tf757Manager = &Tf757{
		dmap:          make(map[int64]*Tf757Socket),
		logger:        svc.Service.ZapLog,
		dlock:         sync.Mutex{},
		ConnectingMap: sync.Map{},
	}
	return Tf757Manager
}

func (h *Tf757) GetSupportDeviceSource() string {
	return shared.TF757
}

func (h *Tf757) GetEntry(standno int64) *Tf757Socket {
	return h.dmap[standno]
}

func (h *Tf757) ReceiveNotifyDeviceList(deviceList []model.DeviceInfo) {
	svc.Info("天赋收到设备列表")
	// svc.Info("---")
	// log.InfoDump(h.GetSupportDeviceSource(), "devicetype")
	// svc.Info("---")

	txDeviceNums := 0
	h.deviceList = deviceList
	// 找出新增的建立连接
	for _, v := range deviceList {
		// if v.DeviceId == "143039954632961227789" {
		// 	log.InfoDump(v, "vvvvv")
		// }
		if v.MachineType != "MODEL_TYPE:TFT757" || v.IpAddress == "" {
			continue
		}
		fmt.Println("收到天赋设备信息：")
		fmt.Println(v)
		// TODO 走不到
		txDeviceNums++
		if _, exist := h.dmap[int64(v.StandNo)]; !exist {
			svc.Infof("天赋下位机开始连接StandNo:[%d]IpAddress[%s]", v.StandNo, v.IpAddress)
			go NewSock(v)
		}
	}
	// 找出删除的断开链接
	for _, v := range h.dmap {
		isset := false
		for _, vv := range deviceList {
			if vv.StandNo == v.dinfo.StandNo {
				isset = true
				break
			}
		}
		if !isset {
			svc.Infof("更新设备列表后删除设备：%s,ip:%s", v.dinfo.DeviceId, v.dinfo.IpAddress)
			v.CloseConn("ReceiveNotifyDeviceList 列表更新删除设备 ")
		}
	}
	// 有设备的情况下、开启 检查重连线程
	if txDeviceNums > 0 {
		ReConnectGoroutineStart.Do(func() {
			time.Sleep(20 * time.Second)
			svc.Info("开启重连检测...")
			go h.ReConnectGoroutine()
		})
	}
}

// 检测没连接上的设备、实现重试、场景：网线插拔
func (h *Tf757) ReConnectGoroutine() {
	interval := 60000 * time.Millisecond
	if svc.Service.Config.TfConfig.IntervalCheckConn > 0 {
		interval = time.Duration(svc.Service.Config.TfConfig.IntervalCheckConn * int(time.Millisecond))
	}
	svc.Infof("重连检测时间间隔:%v", interval)

	for {
		h.dlock.Lock()
		for _, v := range h.deviceList {
			if v.MachineType != h.GetSupportDeviceSource() || v.IpAddress == "" {
				continue
			}
			if _, isset := h.dmap[int64(v.StandNo)]; !isset {
				svc.Info(fmt.Sprintf("%s开始重连... ...", v.IpAddress))
				go NewSock(v)
			}
		}
		h.dlock.Unlock()
		<-time.After(interval)
	}
}

// 从mqtt接收控制指令回调
func (h *Tf757) ReceiveControlCommandNotify(devInfo model.DeviceInfo, dto model.ControlCommandNotifyForDeviceDTO) {
	s, isset := h.dmap[int64(devInfo.StandNo)]
	if !isset {
		svc.Errorf("设备socket不存在%v", devInfo)
		return
	}
	if err := s.DispatchControlCommand(dto); err != nil {
		svc.Info("执行控制指令error:", zap.Error(err))
	}
}

// OnTcpMessage里作为server服务时解码上报实时数据、天富用不到
func (h *Tf757) DecodeReportData(data []byte) (dto model.ControlCommandReportFromDeviceDTO, err error) {
	return
}

func (h *Tf757) EncodeForControlCommandNotify(d model.ControlCommandNotifyForDeviceDTO) (buf []byte, err error) {
	return
}

func (h *Tf757) DeleteSock(standNo int64) {
	h.dlock.Lock()
	delete(h.dmap, standNo)
	h.dlock.Unlock()
}

func (h *Tf757) AddSock(s *Tf757Socket) {
	h.dlock.Lock()
	defer h.dlock.Unlock()
	if _, isset := h.dmap[int64(s.dinfo.StandNo)]; isset {
		s.CloseConn("Adapter AddSock 覆盖旧连接 ")
		return
	}
	h.dmap[int64(s.dinfo.StandNo)] = s
}
