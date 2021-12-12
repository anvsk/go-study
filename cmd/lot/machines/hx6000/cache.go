package hx6000

import "time"

type HxHopeDevStatus struct {
	HopeDeviceStatus int       //设备期望状态
	FlagTime         time.Time //标记时间
	HandleFlag       bool      //处理控制命令的下发情况
}

func (r HxHopeDevStatus) IsOverTime() bool {
	return time.Since(r.FlagTime) >= 15*time.Second
}

func (r HxHopeDevStatus) ReSet() {
	r.HandleFlag = false
}

type HxCraftSendStatusCache struct {
	CraftNo            int       //工艺编号
	BeginSendTime      time.Time //开始下发时间
	SendStepNo         int       //下发到第几个步骤
	DownCraftStepIndex int       //工艺已下发到第几个工步
	VerCraftStepIndex  int       //已验证到第几个工步
}

type DeviceDataCache struct {
	CraftsDetilCode      string
	DeviceStatus         string
	CurrentStepIndexByTx string // oriStepIndex 原始状态
}
