// 处理socket收发逻辑
package hx6000

import (
	"errors"
	"fmt"
	"time"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util"
)

var RetryTimes int = 3

func (s *HxDevice) DispatchControlCommand(dto model.ControlCommandNotifyForDeviceDTO) error {
	svc.Warnf(" 收到控制类指令CommandType:%s", dto.CommandType)
	// var buf []byte
	byteBuffer := make([]byte, 8)
	if dto.StandNo > 99 {
		return errors.New(fmt.Sprintf("航星设备站号不能大于99!!!,当前站号:%d", dto.StandNo))
	}
	bytedevNo := util.GetAsciiStr(util.HandleHxStandNo(int64(s.dinfo.StandNo)), 2, 0)
	switch dto.CommandType {
	case shared.Command_SingleDevice: //实时状态信息
		byteBuffer = append(byteBuffer, []byte{0, 0}...)
		byteBuffer[0] = 5
		byteBuffer[1] = 68
		byteBuffer[2] = 67
		byteBuffer[3] = 82
		byteBuffer[4] = 90
		byteBuffer[5] = bytedevNo[0]
		byteBuffer[6] = bytedevNo[1]
		byteCheck := util.GetHxSumVerify(byteBuffer, 0, 7)
		byteBuffer[7] = byteCheck[0]
		byteBuffer[8] = byteCheck[1]
		byteBuffer[9] = 4

	case shared.Command_Run: //运行
		byteBuffer[0] = 5
		byteBuffer[1] = bytedevNo[0]
		byteBuffer[2] = bytedevNo[1]
		byteBuffer[3] = 87
		byteBuffer[4] = 84
		byteCheck := util.GetHxSumVerify(byteBuffer, 0, 5)
		byteBuffer[5] = byteCheck[0]
		byteBuffer[6] = byteCheck[1]
		byteBuffer[7] = 4
		s.IsCommandRunAck = false

	case shared.Command_Pause: //暂停
		byteBuffer[0] = 5
		byteBuffer[1] = bytedevNo[0]
		byteBuffer[2] = bytedevNo[1]
		byteBuffer[3] = 87
		byteBuffer[4] = 82
		byteCheck := util.GetHxSumVerify(byteBuffer, 0, 5)
		byteBuffer[5] = byteCheck[0]
		byteBuffer[6] = byteCheck[1]
		byteBuffer[7] = 4

	case shared.Command_Manual: //手动
		byteBuffer[0] = 5
		byteBuffer[1] = bytedevNo[0]
		byteBuffer[2] = bytedevNo[1]
		byteBuffer[3] = 87
		byteBuffer[4] = 0x4D
		byteCheck := util.GetHxSumVerify(byteBuffer, 0, 5)
		byteBuffer[5] = byteCheck[0]
		byteBuffer[6] = byteCheck[1]
		byteBuffer[7] = 4

	case shared.Command_Reset: //复位
		byteBuffer[0] = 5
		byteBuffer[1] = bytedevNo[0]
		byteBuffer[2] = bytedevNo[1]
		byteBuffer[3] = 87
		byteBuffer[4] = 81
		byteCheck := util.GetHxSumVerify(byteBuffer, 0, 5)
		byteBuffer[5] = byteCheck[0]
		byteBuffer[6] = byteCheck[1]
		byteBuffer[7] = 4
	case shared.Command_Craft:
		err := s.SendCraftCommand(dto)
		if err != nil {
			return err
		}
		return nil
	}
	util.DebugHex(byteBuffer)
	if byteBuffer != nil {
		s.EnMessageQueue(byteBuffer)
	}
	return nil
}

/******
	航星下发工艺逻辑
	1. 构建下发工艺data	HxCraftNode
	2. 构建上传工艺data	HxCraftNode
	3. 新开线程、先下发工艺、然后监听工步ACK
	4. 下发成功Ack
********/
func (s *HxDevice) SendCraftCommand(dto model.ControlCommandNotifyForDeviceDTO) (err error) {
	s.CurCraftStanNo = s.dinfo.StandNo
	down, err := s.BuildDownCraft(dto)
	if err != nil {
		return
	}
	up, err := s.BuildUploadCraft(dto)
	if err != nil {
		return
	}
	s.UploadingCraftNodeCache = up
	go s.VerificationProcess(down, up)
	return
}

func (s *HxDevice) VerificationProcess(craftNodes, uploadBytes []HxCraftNode) {
	if s.HxCraftSendStatusCache.SendStepNo == 0 {
		s.HxCraftSendStatusCache = HxCraftSendStatusCache{
			CraftNo:       craftNodes[0].CraftNo,
			SendStepNo:    1,
			BeginSendTime: time.Now(),
		}
	}
	craftNodesIndex := 0
	for _, sendbyte := range craftNodes {
		stepFlag := false
		for i := 0; i < RetryTimes; i++ { //重试三次
			s.EnMessageQueue(sendbyte.CraftNodeBytes)
			if sendbyte.IsRptHeader {
				stepFlag = true
				break
			}
			svc.Debugln("【下发】第", craftNodesIndex, "步、第", i+1, "次发送")
			ticker := time.NewTimer(time.Second)
			for {
				select {
				case <-ticker.C:
					goto OUT1
				default:
					if s.DownCraftNodeResponseEvent {
						stepFlag = true
						goto OUT1
					}
				}
			}
		OUT1:
			if stepFlag {
				break
			}
			// 等待超时就发个状态获取包
			s.SendStatusGet()
		}
		craftNodesIndex++
		if !stepFlag {
			svc.Debugln(sendbyte.VerCraftStepIndex+1, "工步下发失败")
			return
		}
		s.HxCraftSendStatusCache.DownCraftStepIndex = sendbyte.VerCraftStepIndex
		time.Sleep(100 * time.Millisecond)
	}
	// --- 上传工艺

	s.HxCraftSendStatusCache.SendStepNo = 2
	uploadStep := 0
	for uindex, node := range uploadBytes {

		if node.IsRptHeader || node.IsRptEnd {
			node.IsVerify = true
			continue
		}
		uploadStep++
		isOK := false
		for i := 0; i < RetryTimes; i++ { //重试三次
			svc.Debugln("【上传】第", uploadStep, "步、第", i, "次发送")
			s.EnMessageQueue(node.CraftNodeBytes)
			ticker := time.NewTimer(time.Second)
			for {
				select {
				case <-ticker.C:
					goto OUT2
				default:
					if s.UploadingCraftNodeCache[uindex].IsVerify {
						isOK = true
						goto OUT2
					}
				}
				time.Sleep(20 * time.Millisecond)
			}
		OUT2:
			if isOK {
				break
			}
			s.SendStatusGet()
		}
		if isOK {
			s.HxCraftSendStatusCache.VerCraftStepIndex = node.VerCraftStepIndex
			time.Sleep(100 * time.Millisecond)
			svc.Debugln("验证【上传{", node.VerCraftStepIndex, "}步工艺】成功")
		} else {
			svc.Debugln("验证【上传{", node.VerCraftStepIndex, "}步工艺】失败")
			// return
		}
	}

	s.HxCraftSendStatusCache.SendStepNo = 3

	go s.DealCommandAck("1070")

	//确保工艺成下发立即发送运行指令
	// s.DispatchControlCommand(model.ControlCommandNotifyForDeviceDTO{
	// 	DeviceId:    s.dinfo.DeviceId,
	// 	StandNo:     int64(s.dinfo.StandNo),
	// 	CommandType: "1010",
	// })
	// isOK := false
	// for i := 0; i < RetryTimes; i++ { //重试三次
	// 	svc.Debugln("下发工艺后发送运行指令次数:", i)
	// 	s.DispatchControlCommand(model.ControlCommandNotifyForDeviceDTO{
	// 		DeviceId:    s.dinfo.DeviceId,
	// 		StandNo:     int64(s.dinfo.StandNo),
	// 		CommandType: "1010",
	// 	})
	// 	ticker := time.NewTimer(time.Second)
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			goto OUT3
	// 		default:
	// 			if s.IsCommandRunAck {
	// 				isOK = true
	// 				goto OUT3
	// 			}
	// 		}
	// 		time.Sleep(20 * time.Millisecond)
	// 	}
	// OUT3:
	// 	if isOK {
	// 		break
	// 	}
	// }
	// 下发完成发送一个状态查询包
	s.SendStatusGet()
	// 发送运行状态信息
	// SendDeviceRunStatus(standNo);
}
