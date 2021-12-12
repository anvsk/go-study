// 实现通用接口方法
package hx6000

import (
	"fmt"
	"sync"
	"time"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
	"go.uber.org/zap"
)

var Hx6000Manager *Hx
var ReConnectGoroutineStart sync.Once
var HxSingleConnOnce sync.Once
var HxSingleConnInstance *HxSingleConn

type Hx struct {
	deviceList    []model.DeviceInfo
	dmap          map[int64]*HxDevice
	logger        *zap.Logger
	modelName     string
	dlock         sync.Mutex
	ConnectingMap sync.Map
}

func InitHxManager() *Hx {
	Hx6000Manager = &Hx{
		dmap:          make(map[int64]*HxDevice),
		logger:        svc.Service.ZapLog,
		dlock:         sync.Mutex{},
		ConnectingMap: sync.Map{},
	}
	return Hx6000Manager
}

func (h *Hx) GetSupportDeviceSource() string {
	return shared.HXE6000
}

func (h *Hx) GetEntry(standno int64) *HxDevice {
	if socket, exsit := h.dmap[standno]; exsit {
		return socket
	}
	return nil
}

func (h *Hx) ReceiveNotifyDeviceList(deviceList []model.DeviceInfo) {
	svc.Info("收到下发设备列表")
	txDeviceNums := 0
	h.deviceList = deviceList
	// 找出新增的建立连接
	for _, v := range deviceList {
		if v.MachineType != h.GetSupportDeviceSource() {
			continue
		}
		txDeviceNums++
		// 初始化连接，只会触发一次
		HxSingleConnOnce.Do(func() {
			conn, err := NewSingleConn()
			if err != nil {
				svc.Error("航星初始化连接失败%v", zap.Error(err))
				return
			}
			HxSingleConnInstance = conn
		})
		if _, exist := h.dmap[int64(v.StandNo)]; !exist {
			svc.Infof("%s开始连接", v.Port)
			go NewDevice(v)
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
			svc.Infof("更新设备列表后删除设备：%s,ip:%s", v.dinfo.DeviceId, v.dinfo.Port)
			v.CloseConn("更新列表后删除设备")
		}
	}
	// 有  设备的情况下、开启 检查重连线程
	if txDeviceNums > 0 {
		ReConnectGoroutineStart.Do(func() {
			time.Sleep(20 * time.Second)
			svc.Info("  开启重连检测...")
			// go h.ReConnectGoroutine()
		})
	}
}

// 检测没连接上的设备、实现重试、场景：网线插拔
func (h *Hx) ReConnectGoroutine() {
	interval := 60000 * time.Millisecond
	if svc.Service.Config.HxConfig.IntervalCheckConn > 0 {
		interval = time.Duration(svc.Service.Config.HxConfig.IntervalCheckConn * int(time.Millisecond))
	}
	svc.Infof("重连检测时间间隔:%v", interval)

	for {
		h.dlock.Lock()
		for _, v := range h.deviceList {
			if v.MachineType != h.GetSupportDeviceSource() {
				continue
			}
			if _, isset := h.dmap[int64(v.StandNo)]; !isset {
				svc.Info(fmt.Sprintf("%s开始重连... ...", v.Port))
				go NewDevice(v)
			}
		}
		h.dlock.Unlock()
		<-time.After(interval)
	}
}

// 从mqtt接收控制指令回调
func (h *Hx) ReceiveControlCommandNotify(devInfo model.DeviceInfo, dto model.ControlCommandNotifyForDeviceDTO) {
	s, isset := h.dmap[int64(devInfo.StandNo)]
	if !isset {
		svc.Errorf("设备socket不存在%v", devInfo)
		return
	}
	if err := s.DispatchControlCommand(dto); err != nil {
		svc.Error(fmt.Sprintln("执行控制指令错误", dto.CommandType, dto.StandNo, err))
	}
}

// OnTcpMessage里作为server服务时解码上报实时数据、  用不到
func (h *Hx) DecodeReportData(data []byte) (dto model.ControlCommandReportFromDeviceDTO, err error) {
	return
}

func (h *Hx) EncodeForControlCommandNotify(d model.ControlCommandNotifyForDeviceDTO) (buf []byte, err error) {
	return
}

func (h *Hx) DeleteSock(standNo int64) {
	h.dlock.Lock()
	delete(h.dmap, standNo)
	h.dlock.Unlock()
}

func (h *Hx) AddSock(s *HxDevice) {
	h.dlock.Lock()
	defer h.dlock.Unlock()
	if _, isset := h.dmap[int64(s.dinfo.StandNo)]; isset {
		s.CloseConn("替换旧链接")
		return
	}
	h.dmap[int64(s.dinfo.StandNo)] = s
}
