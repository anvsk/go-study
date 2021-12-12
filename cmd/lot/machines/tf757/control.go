// 处理socket收发逻辑
package tf757

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util"
	"go.uber.org/zap"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/ctypes"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util/convert"
)

func (s *Tf757Socket) DispatchControlCommand(dto model.ControlCommandNotifyForDeviceDTO) error {
	svc.Warnf("收到控制类指令CommandType:%s", dto.CommandType)
	buf := []byte{byte(s.dinfo.StandNo)}
	switch dto.CommandType {
	case shared.Command_Run:
		buf = append(buf, []byte{16, 0, 0, 0, 0, 0, 1}...)
	case shared.Command_Pause:
		buf = append(buf, []byte{16, 0, 0, 0, 0, 0, 6}...)
	case shared.Command_Stop:
		buf = append(buf, []byte{16, 0, 0, 0, 0, 0, 2}...)
	case shared.Command_Reset:
		buf = append(buf, []byte{16, 0, 0, 0, 0, 0, 3}...)
		s.ResetCmdFlag = true
	case shared.Command_Manual:
		buf = append(buf, []byte{16, 0, 0, 0, 0, 0, 10}...)
	case shared.Command_Auto:
		buf = append(buf, []byte{16, 0, 0, 0, 0, 0, 8}...)
	case shared.Command_Jump:
		buf = append(buf, []byte{16, 0, 0, 0, 0, 0, 4}...)
		var cmd ctypes.CollControlCmd
		bdata, _ := json.Marshal(dto)
		err := json.Unmarshal(bdata, &cmd)
		if err != nil {
			svc.Error("跳步数据反序列化失败", zap.Any("data", dto), zap.Error(err))
			return nil
		}
		// 跳步时根据排产模型工步映射工艺号
		stepNo := s.GetJumpStepNo(cmd.Data.CraftStepIndex)
		stepNoArray := PrePacketExplainRule6(stepNo)
		buf = append(buf, stepNoArray...)
	// case shared.Command_SingleDevice:
	// 	buf = append(buf, []byte{3, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}...)
	case shared.Command_Craft:
		tmpbuf, err := s.DownCraft(dto)
		if err != nil {
			return err
		}
		buf = tmpbuf
	case shared.Command_UpdateCraft:
		tmpbuf, err := s.UpdateCraft(dto)
		if err != nil {
			return err
		}
		buf = tmpbuf
	default:
		return errors.New(fmt.Sprintf("不支持的控制类指令:%s", dto.CommandType))
	}
	if buf != nil {
		s.EnMessageQueue(buf)
	}
	return nil
}

// 下发工艺组装
func (s *Tf757Socket) DownCraft(dto model.ControlCommandNotifyForDeviceDTO) ([]byte, error) {
	s.IsOrderOver = false
	var data ctypes.CollCommandDownCraft
	// 重新解码到下发工艺类型
	convert.AlignStructAndMap(dto, &data)
	// 缓存排产模型
	s.SchedulingExecutionModelCache.CollDownCraftData = *data.Data
	// 文件缓存也加一份
	svc.Service.CacheFile.Set(shared.KeyCollectorTFCraft(s.dinfo.DeviceId), s.SchedulingExecutionModelCache)
	craftModel := data.Data
	res := []byte{}
	res = append(res, byte(s.dinfo.StandNo), 16, 0, 0, 0, 0, 0, 10, 0)
	res = append(res, PrePacketExplainRule5(0)...)
	res = append(res, 1, 0)
	res = append(res, PrePacketExplainRule6(GetStepCount(craftModel.CraftStep))...)
	res = append(res, PrePacketExplainRule3(craftModel.CraftNames, 0)...)
	res = append(res, PrePacketExplainRule6(craftModel.Order.WeightNum)...)
	res = append(res, PrePacketExplainRule5(craftModel.MachineConfig.BathRatio)...)
	res = append(res, PrePacketExplainRule5(craftModel.MachineConfig.BibulousRate)...)
	// stepDetaillist := [][]byte{}
	for _, step := range craftModel.CraftStep {
		for _, operation := range step.CraftOperations {
			operationArray := make([]byte, 14)
			functionType := 0
			if operation.Type != 0 {
				if operation.SideFlag == 1 {
					functionType = 1
				} else {
					functionType = 2
				}
			}
			operationArray[0] = byte(functionType)                              //0主功能/1副功能/2并功能
			operationArray[1] = byte(convert.ToInt(operation.CraftsDetailCode)) //功能编码
			fmt.Println(operationArray)
			for _, attr := range operation.CraftAttrValues {
				switch attr.Sort {
				case 1:
					operationArray[2] = byte(convert.ToInt(attr.AttrValue))
				case 2:
					operationArray[3] = byte(convert.ToInt(attr.AttrValue))
				case 3:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[4] = byte(tmp[1])
					operationArray[5] = byte(tmp[0])
				case 4:
					ival := 0
					// TODO 测试节点
					if strings.Contains(operation.FunctionName, "温控") {
						ival = convert.ToInt(convert.ToFloat(attr.AttrValue) * 10)
					} else {
						ival = convert.ToInt(convert.ToFloat(attr.AttrValue))
					}
					tmp := util.IntToBytesLittle(convert.ToInt(ival))
					operationArray[6] = byte(tmp[1])
					operationArray[7] = byte(tmp[0])
				case 5:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[8] = byte(tmp[1])
					operationArray[9] = byte(tmp[0])
				case 6:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[10] = byte(tmp[1])
					operationArray[11] = byte(tmp[0])
				case 7:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[12] = byte(tmp[1])
					operationArray[13] = byte(tmp[0])
				}
			}
			// stepDetaillist = append(stepDetaillist, operationArray)
			res = append(res, operationArray...)
		}
	}
	// 该代码为复制的C#代码，发现多余已注释
	// stepDetail := ""
	// if len(stepDetaillist) > 0 {
	// 	tempArray := make([]byte, len(stepDetaillist)*14)
	// 	for m := 0; m < len(stepDetaillist); m++ {
	// 		for n := 0; n < len(stepDetaillist[m]); n++ {
	// 			tempArray[14*m+n] = stepDetaillist[m][n]
	// 		}
	// 	}
	// 	stepDetail = strings.ToUpper(hex.EncodeToString(tempArray))
	// }
	// // 工艺步详细信息
	// var stepArray = PrePacketExplainRule4(stepDetail)
	// res = append(res, stepArray...)

	// 加热方式切换
	var warmTypeArray = PrePacketExplainRule1(0)
	res = append(res, warmTypeArray...)

	// 克重
	gramArray := PrePacketExplainRule6(0)
	res = append(res, gramArray...)

	// res = append(res, PrePacketByStdCRC(res)...)

	return res, nil
}

// 修改工艺组装
func (s *Tf757Socket) UpdateCraft(dto model.ControlCommandNotifyForDeviceDTO) ([]byte, error) {
	s.IsOrderOver = false
	var data ctypes.CollUpdateCraftAttrCmd
	// 重新解码到下发工艺类型
	convert.AlignStructAndMap(dto, &data)
	// 覆盖排产模型cache
	s.UpdateCraftModel(data.Data)
	craftModel := s.SchedulingExecutionModelCache
	// 文件缓存也加一份
	svc.Service.CacheFile.Set(shared.KeyCollectorTFCraft(s.dinfo.DeviceId), s.SchedulingExecutionModelCache)
	res := []byte{}
	res = append(res, byte(s.dinfo.StandNo), 16, 0, 0, 0, 0, 0, 16, 0)
	res = append(res, PrePacketExplainRule5(0)...)
	res = append(res, 1, 0)
	res = append(res, PrePacketExplainRule6(GetStepCount(craftModel.CraftStep))...)
	res = append(res, PrePacketExplainRule3(craftModel.CraftNames, 0)...)
	res = append(res, PrePacketExplainRule6(craftModel.Order.WeightNum)...)
	res = append(res, PrePacketExplainRule5(craftModel.MachineConfig.BathRatio)...)
	res = append(res, PrePacketExplainRule5(craftModel.MachineConfig.BibulousRate)...)
	// stepDetaillist := [][]byte{}
	for _, step := range craftModel.CraftStep {
		for _, operation := range step.CraftOperations {
			operationArray := make([]byte, 14)
			functionType := 0
			if operation.Type != 0 {
				if operation.SideFlag == 1 {
					functionType = 1
				} else {
					functionType = 2
				}
			}
			operationArray[0] = byte(functionType)                              //0主功能/1副功能/2并功能
			operationArray[1] = byte(convert.ToInt(operation.CraftsDetailCode)) //功能编码
			// fmt.Println(operationArray)
			for _, attr := range operation.CraftAttrValues {
				switch attr.Sort {
				case 1:
					operationArray[2] = byte(convert.ToInt(attr.AttrValue))
				case 2:
					operationArray[3] = byte(convert.ToInt(attr.AttrValue))
				case 3:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[4] = byte(tmp[1])
					operationArray[5] = byte(tmp[0])
				case 4:
					ival := 0
					// TODO 测试节点
					if strings.Contains(operation.FunctionName, "温控") {
						ival = convert.ToInt(convert.ToFloat(attr.AttrValue) * 10)
					} else {
						ival = convert.ToInt(convert.ToFloat(attr.AttrValue))
					}
					tmp := util.IntToBytesLittle(convert.ToInt(ival))
					operationArray[6] = byte(tmp[1])
					operationArray[7] = byte(tmp[0])
				case 5:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[8] = byte(tmp[1])
					operationArray[9] = byte(tmp[0])
				case 6:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[10] = byte(tmp[1])
					operationArray[11] = byte(tmp[0])
				case 7:
					tmp := util.IntToBytesLittle(convert.ToInt(attr.AttrValue))
					operationArray[12] = byte(tmp[1])
					operationArray[13] = byte(tmp[0])
				}
			}
			// stepDetaillist = append(stepDetaillist, operationArray)
			res = append(res, operationArray...)
		}
	}
	// 该代码为复制的C#代码，发现多余已注释
	// stepDetail := ""
	// if len(stepDetaillist) > 0 {
	// 	tempArray := make([]byte, len(stepDetaillist)*14)
	// 	for m := 0; m < len(stepDetaillist); m++ {
	// 		for n := 0; n < len(stepDetaillist[m]); n++ {
	// 			tempArray[14*m+n] = stepDetaillist[m][n]
	// 		}
	// 	}
	// 	stepDetail = strings.ToUpper(hex.EncodeToString(tempArray))
	// }
	// // 工艺步详细信息
	// var stepArray = PrePacketExplainRule4(stepDetail)
	// res = append(res, stepArray...)

	// 加热方式切换
	// var warmTypeArray = PrePacketExplainRule1(0)
	// res = append(res, warmTypeArray...)

	// 克重
	// gramArray := PrePacketExplainRule6(0)
	// res = append(res, gramArray...)

	// res = append(res, PrePacketByStdCRC(res)...)

	return res, nil
}
