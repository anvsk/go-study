// 管理设备列表和socket连接
package hx6000

import (
	"bytes"
	"strconv"
	"time"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/ctypes"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/handler_mqtt"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util/convert"
)

/****************			从设备接收实时状态数据后回调:分为控制类ACK和实时数据ACK			****************/
func (h *HxSingleConn) AfterReceivePortBytes() {
	defer func() {
		if err := recover(); err != nil {
			svc.Debugln("Hx AfterReceive error:", err)
		}
	}()
	dealLength := 0
	buffer := h.Buffer
	// 遍历处理此次的buffer
	for i := 0; i < buffer.ValidLength-MinPackageLength; {
		//实时状态包
		if buffer.ValidLength-i >= 62 && bytes.Equal(buffer.Buffer[i:i+5], []byte{6, 68, 67, 82, 90}) && (i+62) <= buffer.ValidLength && buffer.Buffer[i+61] == 3 {
			standNo := util.TransferBytes2AsciiCodeString(buffer.Buffer[5:7])
			svc.Infof("Recv:[%d]RealTimeData:%d", standNo, (buffer.ValidLength))
			iCurLength := 62
			realWithBuffer := make([]byte, 62)
			copy(realWithBuffer[:], buffer.Buffer[i:i+62])
			device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
			if device == nil {
				svc.Warnf("接收到%d的数据信息，但设备列表没有", standNo)
			} else {
				// Array.Copy(buffer.Buffer, i, realWithBuffer, 0, 62);
				// fmt.Println("实时数据站号：", util.TransferBytes2AsciiCodeString(realWithBuffer[5:7]))
				// 缓存当前状态
				device.DeviceStatusResponseCache = realWithBuffer
				state := realWithBuffer[24]

				if device.DevHopeStatusCache.HandleFlag {
					// 	//如果状态等于预期状态，则清空标记，同时也上抛数据
					if device.DevHopeStatusCache.HopeDeviceStatus == int(state) || device.DevHopeStatusCache.IsOverTime() {
						device.DevHopeStatusCache.ReSet()
						device.DealDataCommandAck(realWithBuffer)
					}
					// 	//如果状态不符合预期，则不处理任何事情
				} else {
					device.DealDataCommandAck(realWithBuffer)
				}
			}

			i += iCurLength
			dealLength = i
			//工艺上传的启动帧
		} else if buffer.ValidLength-i >= 24 && buffer.Buffer[i] == 6 && buffer.Buffer[i+3] == 82 && buffer.Buffer[i+4] == 80 && (i+24) <= buffer.ValidLength && buffer.Buffer[i+23] == 3 {
			svc.Infof("Recv:工艺上传的启动帧")
			iCurLength := 24
			// int standNo = 0;
			// craftNo := util.ByteArrayToAsciiString(buffer.Buffer, i+1, 2)
			// Logger.Info($"工艺检验，接收到站号{standNo}的上传启动帧报文");
			// standNo = GetCurrSendCraftStationNo(craftNo.ToString());
			standNo := h.CurCraftStanNo
			if standNo == 0 {
				svc.Warnf("接收到%d的工艺上传的启动帧，但没有记录当前工艺站号", standNo)
			} else {
				device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
				if device == nil {
					svc.Warnf("接收到%d的工艺上传的启动帧，但设备列表没有", standNo)
				} else {
					count := 7

					for ; count < 18; count++ {
						if buffer.Buffer[i+count] == 32 {
							break
						}
					}
					craftName := util.ByteArrayToAsciiString(buffer.Buffer, i+7, count-6)
					if len(device.UploadingCraftNodeCache) > 0 {
						rptHeaderNodeIndex := -1
						var rptHeaderNode HxCraftNode
						for k, v := range device.UploadingCraftNodeCache {
							if v.IsRptHeader {
								rptHeaderNodeIndex = k
								rptHeaderNode = v
								break
							}
						}
						if rptHeaderNodeIndex > -1 && rptHeaderNode.VerCraftCode == util.PadLeft(convert.ToStr(device.dinfo.StandNo), 2, "0") && rptHeaderNode.VerCraftName == craftName {
							rptHeaderNode.IsVerify = true
							device.Olock.Lock()
							defer device.Olock.Unlock()
							device.UploadingCraftNodeCache[rptHeaderNodeIndex].IsVerify = true
						}
					}
				}
			}

			i += iCurLength
			dealLength = i
			// 下载工艺启动帧和结束帧
		} else if buffer.ValidLength-i >= 8 && buffer.Buffer[i] == 0x06 && buffer.Buffer[i+3] == 0x57 && buffer.Buffer[i+4] == 0x50 && buffer.ValidLength >= i+8 && buffer.Buffer[i+7] == 0x03 {
			svc.Infof("Recv:下载工艺启动帧和结束帧")
			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();

			i += iCurLength
			dealLength = i
			// 下载工艺工序帧响应
		} else if buffer.ValidLength-i >= 8 && buffer.Buffer[i] == 0x06 && buffer.Buffer[i+3] == 0x57 && buffer.Buffer[i+4] == 0x53 && buffer.ValidLength >= i+8 && buffer.Buffer[i+7] == 0x03 {
			svc.Infof("Recv:下载工艺工序帧响应")

			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
			standNo := convert.ToInt(util.ByteArrayToAsciiString(buffer.Buffer, i+1, 2))
			device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
			if device == nil {
				svc.Warnf("接收到%d的下载工艺工序帧响应，但设备列表没有", standNo)
			} else {
				device.DownCraftNodeResponseEvent = true
			}
			i += iCurLength
			dealLength = i
			//工艺上传的工序帧
		} else if buffer.ValidLength-i >= 24 && buffer.Buffer[i] == 6 && buffer.Buffer[i+1] == 82 && buffer.Buffer[i+2] == 83 && (i+24) <= buffer.ValidLength && buffer.Buffer[i+23] == 3 {
			svc.Infof("Recv:工艺上传的工序帧")
			iCurLength := 24
			standNo := h.CurCraftStanNo
			if standNo == 0 {
				svc.Warnf("接收到%d的工艺上传的启动帧，但没有记录当前工艺站号", standNo)
			} else {
				device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
				if device == nil {
					svc.Warnf("接收到%d的工艺上传的工序帧，但设备列表没有", standNo)
				} else {
					// craftNo := util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2)

					// int standNo = GetCurrSendCraftStationNo(craftNo.ToString());

					// var craftOperationId = util.ByteArrayToAsciiString(buffer.Buffer, i + 5, 2).To<int>();
					craftOperationId := convert.ToInt(util.ByteArrayToAsciiString(buffer.Buffer, i+5, 2))

					// string functionNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 7, 2);
					functionNo := util.ByteArrayToAsciiString(buffer.Buffer, i+7, 2)

					// string p1 = util.ByteArrayToAsciiString(buffer.Buffer, i + 9, 4).To<int>().ToString();
					p1 := util.ByteArrayToAsciiString(buffer.Buffer, i+9, 4)
					p2 := util.ByteArrayToAsciiString(buffer.Buffer, i+13, 4)
					p3 := util.ByteArrayToAsciiString(buffer.Buffer, i+17, 4)

					if len(device.UploadingCraftNodeCache) > 0 {
						rptNodeIndex := -1
						var rptNode HxCraftNode
						for k, v := range device.UploadingCraftNodeCache {
							if v.VerCraftStepIndex == craftOperationId {
								rptNodeIndex = k
								rptNode = v
								break
							}
						}

						if rptNode.VerCraftCode == functionNo &&
							convert.ToInt(rptNode.VerParameterValue1) == convert.ToInt(p1) &&
							convert.ToInt(rptNode.VerParameterValue2) == convert.ToInt(p2) &&
							convert.ToInt(rptNode.VerParameterValue3) == convert.ToInt(p3) {
							device.Olock.Lock()
							defer device.Olock.Unlock()
							device.UploadingCraftNodeCache[rptNodeIndex].IsVerify = true
						}

					}
				}
			}

			// string p2 = util.ByteArrayToAsciiString(buffer.Buffer, i + 13, 4).To<int>().ToString();
			// string p3 = util.ByteArrayToAsciiString(buffer.Buffer, i + 17, 4).To<int>().ToString();
			// Logger.Info($"工艺检验，接收到站号{standNo},工序为{craftOperationId},功能码为{functionNo}，参数为{p1 + p2 + p3}的上传工序帧报文");
			// if (UploadingCraftNodeCache.ContainsKey(standNo))
			// {
			//     var rptNode = UploadingCraftNodeCache[standNo].FirstOrDefault(x => x.VerCraftStepIndex == craftOperationId);
			//     if (rptNode != null && rptNode.VerCraftCode == functionNo && rptNode.VerParameterValue1 == p1 && rptNode.VerParameterValue2 == p2 && rptNode.VerParameterValue3 == p3)
			//         rptNode.IsVerify = true;
			// }
			i += iCurLength
			dealLength = i
			//复位
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WQ" && (i+8 <= buffer.ValidLength) {
			svc.Infof("Recv:复位ACK")

			iCurLength := 8
			// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
			standNo := convert.ToInt(util.ByteArrayToAsciiString(buffer.Buffer, i+1, 2))
			device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
			if device == nil {
				svc.Warnf("接收到%d的复位ACK，但设备列表没有", standNo)
			} else {
				// //先回应控制指令
				device.DealCommandAck(shared.Command_Reset)

				// //再修改实时状态数据
				// DevHopeStatusCache[standNo].Set(56);
				device.DevHopeStatusCache.HopeDeviceStatus = 56
				// Logger.Info($"期望值,{standNo},{56}");
				// h.DealHopeDataCommandAck( DeviceStatus.Awaiting);
				device.DealHopeDataCommandAck(shared.Awaiting)
			}

			i += iCurLength
			dealLength = i
			// 暂停
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WR" && (i+8 <= buffer.ValidLength) {
			svc.Infof("Recv:暂停ACK")
			iCurLength := 8

			standNo := convert.ToInt(util.ByteArrayToAsciiString(buffer.Buffer, i+1, 2))
			device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
			if device == nil {
				svc.Warnf("接收到%d的暂停ACK，但设备列表没有", standNo)
			} else {
				// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
				device.DealCommandAck(shared.Command_Pause)
				// //再修改实时状态数据
				// DevHopeStatusCache[standNo].Set(65);
				device.DevHopeStatusCache.HopeDeviceStatus = 65

				// Logger.Info($"期望值,{standNo},{65}");
				// DealHopeDataCommandAck(standNo, DeviceStatus.Pause);
				device.DealHopeDataCommandAck(shared.Pause)
			}
			i += iCurLength
			dealLength = i
			// 运行
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WT" && (i+8 <= buffer.ValidLength) {
			svc.Infof("Recv:运行ACK")

			iCurLength := 8

			standNo := convert.ToInt(util.ByteArrayToAsciiString(buffer.Buffer, i+1, 2))
			device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
			if device == nil {
				svc.Warnf("接收到%d的暂停ACK，但设备列表没有", standNo)
			} else {
				// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
				// DealCommandAck(standNo, Command.Run);
				device.DealCommandAck(shared.Command_Run)
				// //再修改实时状态数据
				// DevHopeStatusCache[standNo].Set(66);
				device.DevHopeStatusCache.HopeDeviceStatus = 66

				// Logger.Info($"期望值,{standNo},{66}");
				device.DealHopeDataCommandAck(shared.Run)
			}
			i += iCurLength
			dealLength = i
			// 手动
		} else if util.ByteArrayToAsciiString(buffer.Buffer, i+3, 2) == "WM" && (i+8 <= buffer.ValidLength) {
			svc.Infof("Recv:手动ACK")

			iCurLength := 8

			standNo := convert.ToInt(util.ByteArrayToAsciiString(buffer.Buffer, i+1, 2))
			device := Hx6000Manager.GetEntry(int64(convert.ToInt(standNo)))
			if device == nil {
				svc.Warnf("接收到%d的暂停ACK，但设备列表没有", standNo)
			} else {
				// int standNo = util.ByteArrayToAsciiString(buffer.Buffer, i + 1, 2).To<int>();
				// DealCommandAck(standNo, Command.Manual);
				device.DealCommandAck(shared.Command_Manual)
			}
			i += iCurLength
			dealLength = i
		} else {
			// svc.Infof("Recv:没匹配的包")
			i++
		}
	}
	if dealLength > 0 {
		h.Buffer.Clear(dealLength)
	}
	if h.Buffer.ValidLength > MaxPackageLength {
		h.Buffer.Clear(h.Buffer.ValidLength - MaxPackageLength)
	}
	// fmt.Println("ValidLength", h.Buffer.ValidLength)
}

// ACK组装控制类dto
func (s *HxDevice) DealCommandAck(commandType string) {
	ackType := convert.ToStr(convert.ToInt(commandType) + 1)
	svc.Infof(" 响应控制指令ACK：ackType:%v", ackType)
	dto := model.ControlCommandReportFromDeviceDTO{
		StandNo:     int64(s.dinfo.StandNo),
		CommandType: ackType,
		DeviceId:    s.dinfo.DeviceId,
		TraceId:     util.UUID(),
		Timestamp:   convert.ToStr(time.Now().Unix()),
	}
	s.AckDtoToMqtt(dto)
}

// TODO 本方法要结合业务app搞、之前的实时数据整体大结构拿不到
func (s *HxDevice) DealHopeDataCommandAck(hopeStatus int) {

	// var device = DeviceList.FirstOrDefault(x => x.StandNo == standNo);
	// if (device == null)
	//     return;

	// if (!DeviceHandlers.ContainsKey(standNo))
	//     return;

	// //处理结束工步
	// var olddevData = GlobalContext.DeviceDataHandler.GetDeviceData(device.Id);
	// var newDeviceData = olddevData.Clone();
	// newDeviceData.DeviceTime = DateTime.Now;
	// newDeviceData.DeviceStatus = hopeStatus;
	newDeviceData := s.DeviceDataCache
	newDeviceData.DeviceStatus = convert.ToStr(hopeStatus)
	// if(hopeStatus == DeviceStatus.Awaiting)
	// {
	//     newDeviceData.WorkOrderId = 0 ;
	//     newDeviceData.WorkOrderNo = string.Empty;
	//     newDeviceData.CraftSumTime = 0;
	//     newDeviceData.CraftStepId = 0;
	//     newDeviceData.MasterOperationId = 0;
	//     newDeviceData.CraftsDetilCode = 0;
	//     newDeviceData.MainCraftsDetilName = string.Empty;
	//     newDeviceData.UsedTime = 0;
	// }

	// DataCommandAck ack = new DataCommandAck(newDeviceData) { StandNo = standNo, DeviceId = device.Id };

	// if (olddevData != null)
	// {
	//     if (olddevData.DeviceStatus == DeviceStatus.Run && ack.DeviceData.DeviceStatus == DeviceStatus.Awaiting)
	//     {
	//         ack.DeviceData.CraftsDetilCode = olddevData.CraftsDetilCode;
	//         ack.DeviceData.CurrentStepIndexByTx = olddevData.CurrentStepIndexByTx;
	//         ack.DeviceData.IsEndFlag = true;
	//     }
	// }

	// HardwareCommunicationHandler.DataCommandAck(ack);
	// Console.WriteLine(ack.ToSafeString());
}

// ACK组装实时数据类dto
func (s *HxDevice) DealDataCommandAck(receviedData []byte) {
	// log.Debug(" 处理实时数据... ...")
	dto := model.ControlCommandReportFromDeviceDTO{
		StandNo:     int64(s.dinfo.StandNo),
		CommandType: shared.Command_SingleDevice,
		DeviceId:    s.dinfo.DeviceId,
		TraceId:     util.UUID(),
		Timestamp:   convert.ToStr(time.Now().Unix()),
	}

	deviceData := ctypes.DeviceData_dtl{}

	szd := util.ByteArrayToAsciiString(receviedData, 24, 1)
	if szd == "B" {
		deviceData.RunStatus = "1"
		deviceData.DeviceStatus = shared.Run
		deviceData.HandAutoStatus = 1
	} else if szd == "A" {
		deviceData.RunStatus = "3"
		deviceData.DeviceStatus = shared.Pause
		deviceData.HandAutoStatus = 1
	} else if szd == "9" { //手自动
		deviceData.RunStatus = "2"
		deviceData.DeviceStatus = shared.Run
		deviceData.HandAutoStatus = 0
	} else if szd == "0" {
		return
	} else {
		deviceData.RunStatus = "5"
		deviceData.DeviceStatus = shared.Awaiting
		deviceData.HandAutoStatus = 1
	}

	if deviceData.RunStatus == "5" {
		deviceData.CraftsCode = ""
		deviceData.CraftStepIndex = "0"
	} else {
		//工艺号
		tmp, _ := strconv.ParseInt(util.ByteArrayToAsciiString(receviedData, 31, 2), 16, 0)
		deviceData.CraftsCode = convert.ToStr(tmp)
		//工步号
		craftsDetilCode, _ := strconv.ParseInt(util.ByteArrayToAsciiString(receviedData, 33, 2), 16, 0)
		deviceData.CraftStepIndex = convert.ToStr(craftsDetilCode - 1)
		deviceData.OriStepIndex = convert.ToStr(craftsDetilCode)
	}

	//主缸液位
	deviceData.RealMainWater = float64(convert.ToInt(util.ByteArrayToAsciiString(receviedData, 15, 4)))
	//主缸温度
	deviceData.RealTem = convert.ToFloat(util.ByteArrayToAsciiString(receviedData, 7, 4)) / 10
	//副缸液位
	deviceData.AuxiRealMainWater = float64(convert.ToInt(util.ByteArrayToAsciiString(receviedData, 19, 4)))
	//副缸温度
	deviceData.AuxiRealTem = convert.ToFloat(util.ByteArrayToAsciiString(receviedData, 11, 4)) / 10

	deviceData.SwitchOutputs = GetSwitchOutputs(receviedData)
	dto.Data = deviceData
	s.AckDtoToMqtt(dto)
}

func (s *HxDevice) AckDtoToMqtt(dto model.ControlCommandReportFromDeviceDTO) {
	handler_mqtt.EnQueueMsg(dto)
	// msg, _ := handler_mqtt.BuildMqttMsg(dto)
	// err := svc.Service.Mqtt.Send(msg)
	// if err != nil {
	// 	svc.Warnf("实时数据响应ACK发送到mqtt失败；%v", err)
	// 	return
	// }
}

func GetSwitchOutputs(statusArray []byte) (outList [128]bool) {
	outInfo := util.ByteArrayToAsciiString(statusArray, 25, 4)
	sb1 := ""
	sb1 += util.PadLeft(hxSwitchTansfer(string(outInfo[2])), 4, "0")
	sb1 += util.PadLeft(hxSwitchTansfer(string(outInfo[3])), 4, "0")
	sb1 += util.PadLeft(hxSwitchTansfer(string(outInfo[0])), 4, "0")
	sb1 += util.PadLeft(hxSwitchTansfer(string(outInfo[1])), 4, "0")
	i := 0
	for n := 16; n > 0; n-- {
		if string(sb1[n-1]) == "1" {
			outList[i] = true
		}
		i++
	}
	return outList
}

func hxSwitchTansfer(s string) string {
	tmp, _ := strconv.ParseInt(s, 16, 0)
	return strconv.FormatInt(int64(tmp), 2)
}
