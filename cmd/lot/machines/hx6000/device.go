// 每台机器结构题对象
package hx6000

import (
	"sync"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
)

type HxDevice struct {

	// 所有设备绑定一个连接
	*HxSingleConn

	// 设备信息
	dinfo model.DeviceInfo

	// 数据传输socket
	// Conn *serial.Port

	// 定时请求状态队列
	IntervalRequestChan chan struct{}

	// 运行是否回复
	IsCommandRunAck bool

	// 上次状态信息[没用到]
	// TODO 缓存上次状态信息 CraftsDetilCode、DeviceStatus、CurrentStepIndexByTx
	DeviceDataCache DeviceDataCache

	// 初始状态
	InitFlag bool

	// 操作锁
	Olock sync.Mutex

	// 实时数据响应缓存
	DeviceStatusResponseCache []byte

	// 期望值缓存[没用到]
	DevHopeStatusCache HxHopeDevStatus

	// 上传工艺缓存
	UploadingCraftNodeCache []HxCraftNode

	// 工艺发送状态缓存
	HxCraftSendStatusCache HxCraftSendStatusCache

	// 下载工艺响应
	DownCraftNodeResponseEvent bool

	// 排产模型缓存
	SchedulingExecutionModelCache model.SchedulingExecutionModel
}

func NewDevice(info model.DeviceInfo) (s *HxDevice, err error) {
	if _, ok := Hx6000Manager.ConnectingMap.Load(info.StandNo); ok {
		return
	}
	Hx6000Manager.ConnectingMap.Store(info.StandNo, true)
	defer Hx6000Manager.ConnectingMap.Delete(info.StandNo)

	s = &HxDevice{
		HxSingleConn:        HxSingleConnInstance,
		dinfo:               info,
		IntervalRequestChan: make(chan struct{}),
		Olock:               sync.Mutex{},
	}
	Hx6000Manager.AddSock(s)

	svc.Infof("注册设备 StandNo:%d!", s.dinfo.StandNo)
	return
}
