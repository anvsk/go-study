// 管理设备列表和socket连接
package tf757

import (
	"time"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/ctypes"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/enum/tf_ack"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/handler_mqtt"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util/convert"
)

/****************			从设备接收实时状态数据后回调:分为控制类ACK和实时数据ACK			****************/
func (s *Tf757Socket) AfterReceive(receiveBytes []byte) {
	if len(receiveBytes) < 15 {
		return
	}
	if receiveBytes[0] == byte(s.dinfo.StandNo) && receiveBytes[1] == 1 {
		s.DealDataCommandAck(receiveBytes)
	} else if receiveBytes[0] == 0x20 && len(receiveBytes) == 16 {
		if receiveBytes[5] != byte(s.dinfo.StandNo) {
			svc.Error("接收站号不一致")
			return
		}
		status := receiveBytes[2]
		if status != 0 {
			svc.Error("收到指令响应状态为失败")
			return
		}
		cmdNo := receiveBytes[1]
		s.DealCommandAck(int(cmdNo))
	} else {
		// 其他格式

	}
}

// ACK组装控制类dto
func (s *Tf757Socket) DealCommandAck(cmdNo int) {
	ack := ""
	switch cmdNo {
	case tf_ack.Run:
		ack = shared.Command_Run
	case tf_ack.Pause, tf_ack.Stop:
		ack = shared.Command_Pause
	case tf_ack.Jump:
		ack = shared.Command_Jump
	case tf_ack.Reset:
		ack = shared.Command_Reset
	case tf_ack.UpdateCraftOnline:
		ack = shared.Command_UpdateCraft
	case tf_ack.SendCraft:
		ack = shared.Command_Craft
	default:
		svc.Errorf("不支持的响应类型%d", cmdNo)
		return
	}
	ackType := convert.ToStr(convert.ToInt(ack) + 1)
	svc.Infof("响应控制指令ACK：ackType:%v", ackType)
	dto := model.ControlCommandReportFromDeviceDTO{
		StandNo:     int64(s.dinfo.StandNo),
		CommandType: ackType,
		DeviceId:    s.dinfo.DeviceId,
		TraceId:     util.UUID(),
		Timestamp:   convert.ToStr(time.Now().Unix()),
	}
	s.AckDtoToMqtt(dto)
}

// ACK组装实时数据类dto
func (s *Tf757Socket) DealDataCommandAck(bytes []byte) {
	// C#版逻辑
	// 判断结束工步
	// if (oldData.DeviceStatus == DeviceStatus.Run &&
	// 	(ack.DeviceData.DeviceStatus == DeviceStatus.Awaiting || ack.DeviceData.DeviceStatus == DeviceStatus.Ready))
	// {
	// 	//ack.DeviceData.IsEndFlag = true;
	// 	if (!IsOrderOver.ContainsKey(standNo))
	// 		IsOrderOver.Add(standNo, true);
	// 	else
	// 		IsOrderOver[standNo] = true;

	// 	ack.DeviceData.DeviceStatus = DeviceStatus.Awaiting;
	// }
	// if (oldData.DeviceStatus == DeviceStatus.Pause &&
	// 	(ack.DeviceData.DeviceStatus == DeviceStatus.Awaiting || ack.DeviceData.DeviceStatus == DeviceStatus.Ready))
	// {
	// 	//ack.DeviceData.IsEndFlag = true;
	// 	if (!IsOrderOver.ContainsKey(standNo))
	// 		IsOrderOver.Add(standNo, true);
	// 	else
	// 		IsOrderOver[standNo] = true;
	// 	ack.DeviceData.DeviceStatus = DeviceStatus.Awaiting;
	// }
	ack := s.ExplainDataCommandAck(bytes)
	ackData := ack.Data.(ctypes.DeviceData_dtl)
	if s.DeviceStatusCache == shared.Run && (ackData.DeviceStatus == shared.Awaiting || ackData.DeviceStatus == shared.Ready) {
		ackData.DeviceStatus = shared.Awaiting
		s.IsOrderOver = true
	}
	if s.DeviceStatusCache == shared.Pause && (ackData.DeviceStatus == shared.Awaiting || ackData.DeviceStatus == shared.Ready) {
		ackData.DeviceStatus = shared.Awaiting
		s.IsOrderOver = true
	}
	ack.Data = ackData
	s.DeviceStatusCache = ackData.DeviceStatus
	s.AckDtoToMqtt(ack)
}
func (s *Tf757Socket) ExplainDataCommandAck(bytes []byte) model.ControlCommandReportFromDeviceDTO {
	// log.Debug("处理实时数据... ...")
	dto := model.ControlCommandReportFromDeviceDTO{
		StandNo:     int64(s.dinfo.StandNo),
		CommandType: shared.Command_SingleDevice,
		DeviceId:    s.dinfo.DeviceId,
		TraceId:     util.UUID(),
		Timestamp:   convert.ToStr(time.Now().Unix()),
	}
	tfRunStatus := bytes[4]
	dstatus := s.getDeviceStatus(int(tfRunStatus))
	hand := 1
	if bytes[2] > 0x13 {
		hand = 0
	}
	curStepIndex := int(UnPackByExplainRule1(bytes, 7, 8))
	if dstatus == shared.Awaiting || dstatus == shared.Ready {
		curStepIndex = 0
	} else {
		curStepIndex = s.GetCurrRealStepNo(curStepIndex)
	}
	deviceData := ctypes.DeviceData_dtl{
		RunStatus:      convert.ToStr(tfRunStatus),
		DeviceStatus:   dstatus,
		HandAutoStatus: hand,
		CraftStepIndex: convert.ToStr(curStepIndex),
		MainPumpSpeed:  float64((int(bytes[269]))),
	}
	deviceData.RealMainWater = UnPackByExplainRule1(bytes, 95, 96)        //T7_MainCylinderWaterLevel	主缸水位	95	96	1
	deviceData.RealTem = UnPackByExplainRule1(bytes, 81, 82) / 10         //T7_MainCylinderTemperature	主缸温度	81	82	3
	deviceData.AuxiRealMainWater = UnPackByExplainRule1(bytes, 101, 102)  //T7_AssistantCylinderWaterLevel	副缸水位	101	102	1
	deviceData.AuxiRealTem = UnPackByExplainRule3(bytes, 87, 88)          //T7_AssistantCylinderTemperature	副缸温度	87	88	3
	deviceData.FeedMater1Tem = UnPackByExplainRule3(bytes, 83, 84)        //T7_CylinderTemperature1	料1缸温度	83	84	3
	deviceData.FeedMater1Water = UnPackByExplainRule1(bytes, 97, 98)      //T7_CylinderWaterLevel1	料1缸水位	97	98	1
	deviceData.FeedMater2Tem = UnPackByExplainRule3(bytes, 85, 86)        //T7_CylinderTemperature2	料2缸温度	85	86	3
	deviceData.FeedMater2Water = UnPackByExplainRule1(bytes, 99, 100)     //T7_CylinderWaterLevel2	料2缸水位	99	100	1
	deviceData.TotalWaterVolume = UnPackByExplainRule22(bytes, 9, 16)     //T7_MainFlowMeter	累计主缸用水	9	16	22
	deviceData.TotalQuantityVolume = UnPackByExplainRule22(bytes, 17, 24) //T7_GasMeter	累计用气	17	24	22
	deviceData.TotalSteamVolume = UnPackByExplainRule22(bytes, 57, 64)    //T7_ElectricEnergyMeter	累计用电	57	64	22

	deviceData.Fanspeeds = [8]int{
		UnPackByExplainRule18(bytes, 151, 152),
		UnPackByExplainRule18(bytes, 153, 154),
		UnPackByExplainRule18(bytes, 155, 156),
		UnPackByExplainRule18(bytes, 157, 158),
		0, 0, 0, 0,
	}
	deviceData.Tyrespeeds = [8]int{
		UnPackByExplainRule18(bytes, 261, 262),
		UnPackByExplainRule18(bytes, 263, 264),
		UnPackByExplainRule18(bytes, 265, 266),
		UnPackByExplainRule18(bytes, 267, 268),
		0, 0, 0, 0,
	}
	deviceData.TotalCirSpeedCloths = [8]int{
		int(UnPackByExplainRule1(bytes, 235, 236)),
		int(UnPackByExplainRule1(bytes, 239, 240)),
		int(UnPackByExplainRule1(bytes, 243, 244)),
		int(UnPackByExplainRule1(bytes, 247, 248)),
		0, 0, 0, 0,
	}
	deviceData.SwitchInputs = GetSwitchInputs(bytes)
	deviceData.SwitchOutputs = GetSwitchOutputs(bytes)
	dto.Data = deviceData
	go s.UpAlarm(bytes)
	return dto
}

// 上报报警数据
func (s *Tf757Socket) UpAlarm(bytes []byte) {
	if time.Since(s.LastUpAlarmTime) > 10*time.Second {
		alarmSli := bytes[249:259]
		errIndex := 0
		dto := ctypes.AlarmData{}
		dtoMap := map[string]ctypes.AlarmItem{}
		for _, v := range alarmSli {
			for i := 0; i < 8; i++ {
				if (int(v) >> i & 1) == 1 {
					tipCode := convert.ToStr(errIndex)
					dtoMap[tipCode] = ctypes.AlarmItem{
						TipCode:  convert.ToStr(tipCode),
						TipMsg:   "报警",
						TipType:  10,
						Priority: 1,
						WarnTime: time.Now().UTC(),
					}
				}
				errIndex++
			}
		}
		dto.List = dtoMap
		dto.TipType = 10
		res := model.ControlCommandReportFromDeviceDTO{
			StandNo:     int64(s.dinfo.StandNo),
			CommandType: shared.Command_Alarm,
			DeviceId:    s.dinfo.DeviceId,
			TraceId:     util.UUID(),
			Timestamp:   convert.ToStr(time.Now().Unix()),
			Data:        dto,
		}
		s.AckDtoToMqtt(res)
		s.LastUpAlarmTime = time.Now()
	}
}

func (s *Tf757Socket) AckDtoToMqtt(dto model.ControlCommandReportFromDeviceDTO) {
	handler_mqtt.EnQueueMsg(dto)
	// msg, _ := handler_mqtt.BuildMqttMsg(dto)
	// err := svc.Service.Mqtt.Send(msg)
	// if err != nil {
	// 	svc.Warnf("实时数据响应ACK发送到mqtt失败；%v", err)
	// 	return
	// }
}
