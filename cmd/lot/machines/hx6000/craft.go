package hx6000

import (
	"encoding/json"
	"errors"
	"strconv"

	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/ctypes"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/model"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/util/convert"
)

func (s *HxDevice) BuildUploadCraft(dto model.ControlCommandNotifyForDeviceDTO) (nodes []HxCraftNode, err error) {
	var commandCraft ctypes.CollCommandDownCraft

	bdata, err := json.Marshal(dto)
	if err != nil {
		return
	}

	//json解析到结构体
	err = json.Unmarshal(bdata, &commandCraft)
	if err != nil {
		return
	}
	listModel := commandCraft.Data

	// #region 启动上传帧
	// byte[] byteCraft = new byte[(listModel.CraftStep.Count) * 12 + 10];
	byteCraft := make([]byte, len(listModel.CraftStep)*12+10)
	byteCraft[0] = 5
	bytedevNo := util.GetAsciiStr(util.HandleHxStandNo(int64(s.dinfo.StandNo)), 2, 0)
	// bytedevNo := ConvertHelper.GetAsciiStr(command.StandNo.ToString().PadLeft(2, '0'), 2, 0);
	byteCraft[1] = bytedevNo[0]
	byteCraft[2] = bytedevNo[1]
	byteCraft[3] = 82
	byteCraft[4] = 80

	// 数据包工艺编号使用实际工艺编号 modify by douqishuang 20210810
	// byte[] craftNo = new byte[2];
	iCraftNo, err := strconv.Atoi(listModel.CraftCodes)
	if err != nil {
		return
	}
	// int.TryParse(command.Model.CraftCodes.Substring(command.Model.CraftCodes.IndexOfFirstDigital()), out iCraftNo);
	//int.TryParse(command.Model.CraftCodes.Substring(command.Model.CraftCodes.IndexOf('-') + 1), out iCraftNo);
	// craftNo = ConvertHelper.GetAsciiStr(iCraftNo.ToString("D2"), 2, 0);
	craftNo := util.GetAsciiStr(listModel.CraftCodes, 2, 0)

	byteCraft[5] = craftNo[0]
	byteCraft[6] = craftNo[1]
	//byteCraft[5] = bytedevNo[0];
	//byteCraft[6] = bytedevNo[1];

	strhex := util.PadLeft(util.HexSumVerify(byteCraft, 0, 7), 4, "0")

	// byte[] byteCkeck = new byte[2];
	byteCkeck := util.GetAsciiStr(strhex[len(strhex)-2:], 2, 0)
	// byteCkeck = ConvertHelper.GetAsciiStr(strhex.Substring(strhex.Length - 2, 2), 2, 0);

	byteCraft[7] = byteCkeck[0]
	byteCraft[8] = byteCkeck[1]

	byteCraft[9] = 4
	// byte[] byte_termp = new byte[10];
	byte_termp := make([]byte, 10)
	copy(byte_termp, byteCraft)
	// Array.Copy(byteCraft, 0, byte_termp, 0, 10);

	var node = CreateRptHeaderNode(byte_termp, 0, iCraftNo)
	node.IsVerify = false
	node.IsRptHeader = true
	node.VerStandNo = util.HandleHxStandNo(int64(s.dinfo.StandNo))
	node.VerCraftCode = util.HandleHxStandNo(int64(s.dinfo.StandNo))
	node.VerCraftName = listModel.CraftNames
	node.WorkOrderId = 0
	nodes = append(nodes, node)
	// #endregion

	// #region 工序帧
	Count := 0
	for i1 := 0; i1 < len(listModel.CraftStep); i1++ {
		Count = 10 + i1*12
		byteCraft[Count] = 5
		byteCraft[Count+1] = bytedevNo[0]
		byteCraft[Count+2] = bytedevNo[1]
		byteCraft[Count+3] = 82
		byteCraft[Count+4] = 83
		// 工艺编号使用实际编号
		byteCraft[Count+5] = craftNo[0]
		byteCraft[Count+6] = craftNo[1]
		//byteCraft[Count + 5] = bytedevNo[0];
		//byteCraft[Count + 6] = bytedevNo[1];
		// byte[] bytegxNo = new byte[2];
		// bytegxNo = ConvertHelper.GetAsciiStr((i1 + 1).ToString().PadLeft(2, '0'), 2, 0);
		bytegxNo := util.GetAsciiStr(util.PadLeft(convert.ToStr(i1+1), 2, "0"), 2, 0)

		byteCraft[Count+7] = bytegxNo[0]
		byteCraft[Count+8] = bytegxNo[1]
		// strhex = ConvertHelper.HexSumVerify(byteCraft, Count, 9).PadLeft(4, '0');
		strhex := util.PadLeft(util.HexSumVerify(byteCraft, Count, 9), 4, "0")

		// byteCkeck = new byte[2];
		// byteCkeck = ConvertHelper.GetAsciiStr(strhex.Substring(strhex.Length - 2, 2), 2, 0);
		byteCkeck := util.GetAsciiStr(strhex[len(strhex)-2:], 2, 0)

		byteCraft[Count+9] = byteCkeck[0]
		byteCraft[Count+10] = byteCkeck[1]
		byteCraft[Count+11] = 4

		// byte_termp = new byte[12];
		byte_termp := make([]byte, 12)
		copy(byte_termp, byteCraft[Count:])
		// Array.Copy(byteCraft, Count, byte_termp, 0, 12);

		baseCraftOperations := listModel.CraftStep[i1].CraftOperations
		if len(baseCraftOperations) == 0 {
			continue
		}
		var baseCraftOperation = baseCraftOperations[0]
		var functionNo = util.PadLeft(baseCraftOperation.CraftsDetailCode, 2, "0")
		// byte[] bytefuncNo = new byte[2];//工序指令编号
		var paramList = baseCraftOperation.CraftAttrValues
		plen := len(paramList)
		p1 := "0"
		p2 := "0"
		p3 := "0"
		if paramList != nil {
			if plen > 0 {
				p1 = paramList[0].AttrValue
			}
			if plen > 1 {
				p2 = paramList[1].AttrValue
			}
			if plen > 2 {
				p3 = paramList[2].AttrValue
			}
			// p1 = paramList.Count > 0 ? paramList[0].AttrValue : "0";
			// p2 = paramList.Count > 1 ? paramList[1].AttrValue : "0";
			// p3 = paramList.Count > 2 ? paramList[2].AttrValue : "0";
		}

		node = CreateRptHeaderNode(byte_termp, 0, 0)
		node.IsVerify = false
		node.IsRptHeader = false
		node.IsRptEnd = false
		node.VerStandNo = util.HandleHxStandNo(int64(s.dinfo.StandNo))
		node.VerCraftCode = functionNo
		node.VerCraftName = listModel.CraftNames
		node.WorkOrderId = 0
		node.VerCraftStepIndex = i1 + 1
		node.VerParameterValue1 = p1
		node.VerParameterValue2 = p2
		node.VerParameterValue3 = p3

		nodes = append(nodes, node)

	}
	// #endregion

	return
}

func (s *HxDevice) BuildDownCraft(dto model.ControlCommandNotifyForDeviceDTO) (nodes []HxCraftNode, err error) {
	var commandCraft ctypes.CollCommandDownCraft

	bdata, err := json.Marshal(dto)
	if err != nil {
		return
	}

	//json解析到结构体
	err = json.Unmarshal(bdata, &commandCraft)
	if err != nil {
		return
	}
	// 缓存排产模型
	s.SchedulingExecutionModelCache.CollDownCraftData = *commandCraft.Data
	// 文件缓存也加一份
	svc.Service.CacheFile.Set(shared.KeyCollectorHxCraft(s.dinfo.DeviceId), s.SchedulingExecutionModelCache)

	listModel := commandCraft.Data

	// #region 开始帧
	byteName := util.Transfer2AsciiCodeByte(listModel.CraftNames)
	namelength := len(byteName)

	// byte[] byteCraft = new byte[48 + listModel.CraftStep.Count * 26];
	byteCraft := make([]byte, 48+len(listModel.CraftStep)*26)
	byteCraft[0] = 5
	// byte[] bytedevNo = new byte[2];
	// bytedevNo = ConvertHelper.GetAsciiStr(command.StandNo.ToString().PadLeft(2, '0'), 2, 0);
	bytedevNo := util.GetAsciiStr(util.HandleHxStandNo(int64(s.dinfo.StandNo)), 2, 0)
	byteCraft[1] = bytedevNo[0]
	byteCraft[2] = bytedevNo[1]
	byteCraft[3] = 87
	byteCraft[4] = 80

	// 数据包工艺编号使用实际工艺编号 modify by douqishuang 20210810
	// int.TryParse(command.Model.CraftCodes.Substring(command.Model.CraftCodes.IndexOfFirstDigital()), out iCraftNo);
	iCraftNo, err := strconv.Atoi(listModel.CraftCodes)
	if err != nil {
		svc.Warnf("工艺编号错误，将用默认99")
		// return
	} else {
		iCraftNo = 99
	}
	if iCraftNo > 99 {
		err = errors.New("工艺编号超出长度")
		return
	}
	//int.TryParse(command.Model.CraftCodes.Substring(command.Model.CraftCodes.IndexOf('-') + 1), out iCraftNo);
	// craftNo = ConvertHelper.GetAsciiStr(iCraftNo.ToString("D2"), 2, 0);
	//craftNo = ConvertHelper.GetAsciiStr(command.Model.CraftCodes.Substring(command.Model.CraftCodes.IndexOf('-') + 1), 2, 0);
	craftNo := util.GetAsciiStr(listModel.CraftCodes, 2, 0)
	byteCraft[5] = craftNo[0]
	byteCraft[6] = craftNo[1]
	for i := 0; i < len(byteName); i++ {
		byteCraft[7+i] = byteName[i]
	}
	spaceNum := 11 - namelength
	for j := 0; j < spaceNum; j++ {
		byteCraft[7+j+namelength] = 32
	}
	bytegyNum := make([]byte, 2)
	bytegyNum = util.GetAsciiStr(convert.ToStr(len(listModel.CraftStep)), 2, 0)
	byteCraft[18] = bytegyNum[0]
	byteCraft[19] = bytegyNum[1]
	byteCraft[20] = 49

	strhex := util.PadLeft(util.HexSumVerify(byteCraft, 0, 21), 4, "0")
	byteCkeck := util.GetAsciiStr(strhex[len(strhex)-2:], 2, 0)
	byteCraft[21] = byteCkeck[0]
	byteCraft[22] = byteCkeck[1]

	byteCraft[23] = 4

	// byte[] byte_termp= new byte[24];
	byte_termp := make([]byte, 24)
	// TODO
	// Array.Copy(byteCraft, 0, byte_termp, 0, 24);
	copy(byte_termp, byteCraft)
	nodes = append(nodes, CreateRptHeaderNode(byte_termp, 0, iCraftNo))
	// #endregion

	// #region 工序帧
	Count := 0
	for i1 := 0; i1 < len(listModel.CraftStep); i1++ {
		Count = 24 + 26*i1
		byteCraft[Count] = 5
		byteCraft[Count+1] = bytedevNo[0]
		byteCraft[Count+2] = bytedevNo[1]
		byteCraft[Count+3] = 87
		byteCraft[Count+4] = 83

		byteCraft[Count+5] = craftNo[0]
		byteCraft[Count+6] = craftNo[1]
		// byte[] bytegxNo = new byte[2];
		// var baseCraftOperations = listModel.CraftStep[i1].BaseCraftOperations.Where(x => x.CraftStepId > 0).OrderBy(x => x.Sort).ToList();
		baseCraftOperations := listModel.CraftStep[i1].CraftOperations
		if len(baseCraftOperations) == 0 {
			continue
		}

		// bytegxNo = ConvertHelper.GetAsciiStr((i1 + 1).ToString().PadLeft(2, '0'), 2, 0);
		bytegxNo := util.GetAsciiStr(util.PadLeft(convert.ToStr(i1+1), 2, "0"), 2, 0)
		byteCraft[Count+7] = bytegxNo[0]
		byteCraft[Count+8] = bytegxNo[1]
		// var baseCraftOperation = baseCraftOperations.FirstOrDefault();
		baseCraftOperation := baseCraftOperations[0]
		pt := baseCraftOperation.CraftsDetailCode
		bytefuncNo := util.GetAsciiStr(util.PadLeft(pt, 2, "0"), 2, 0)
		byteCraft[Count+9] = bytefuncNo[0]
		byteCraft[Count+10] = bytefuncNo[1]

		var paramList = baseCraftOperation.CraftAttrValues
		plen := len(paramList)
		p1 := 0
		p2 := 0
		p3 := 0
		if paramList != nil {
			if plen > 0 {
				p1 = convert.ToInt(paramList[0].AttrValue)
			}
			if plen > 1 {
				p2 = convert.ToInt(paramList[1].AttrValue)
			}
			if plen > 2 {
				p3 = convert.ToInt(paramList[2].AttrValue)
			}
			// p1 = paramList.Count > 0 ? paramList[0].AttrValue.To<int>() : 0;
			// p2 = paramList.Count > 1 ? paramList[1].AttrValue.To<int>() : 0;
			// p3 = paramList.Count > 2 ? paramList[2].AttrValue.To<int>() : 0;
		}
		byte1 := util.GetAsciiStr(util.PadLeft(convert.ToStr(p1), 4, "0"), 4, 0)
		byteCraft[Count+11] = byte1[0]
		byteCraft[Count+12] = byte1[1]
		byteCraft[Count+13] = byte1[2]
		byteCraft[Count+14] = byte1[3]
		byte2 := util.GetAsciiStr(util.PadLeft(convert.ToStr(p2), 4, "0"), 4, 0)
		byteCraft[Count+15] = byte2[0]
		byteCraft[Count+16] = byte2[1]
		byteCraft[Count+17] = byte2[2]
		byteCraft[Count+18] = byte2[3]
		byte3 := util.GetAsciiStr(util.PadLeft(convert.ToStr(p3), 4, "0"), 4, 0)
		byteCraft[Count+19] = byte3[0]
		byteCraft[Count+20] = byte3[1]
		byteCraft[Count+21] = byte3[2]
		byteCraft[Count+22] = byte3[3]

		strhex = util.PadLeft(util.HexSumVerify(byteCraft, Count, 23), 4, "0")
		byteCkeck := util.GetAsciiStr(strhex[len(strhex)-2:], 2, 0)
		// byteCkeck = ConvertHelper.GetAsciiStr(strhex.Substring(strhex.Length - 2, 2), 2, 0);

		byteCraft[Count+23] = byteCkeck[0]
		byteCraft[Count+24] = byteCkeck[1]
		byteCraft[Count+25] = 4

		// byte_termp = new byte[26];
		byte_termp := make([]byte, 26)
		copy(byte_termp, byteCraft[Count:])
		// Array.Copy(byteCraft, Count, byte_termp, 0,26);
		// nodes.Add(HxCraftNode.CreateNode(byte_termp, command.Model.WorkOrderId, i1));
		nodes = append(nodes, CreateNode(byte_termp, 0, i1))
	}
	// #endregion

	// #region 结束帧
	Count = Count + 26
	byteCraft[Count] = 5
	byteCraft[Count+1] = bytedevNo[0]
	byteCraft[Count+2] = bytedevNo[1]
	byteCraft[Count+3] = 87
	byteCraft[Count+4] = 80

	byteCraft[Count+5] = craftNo[0]
	byteCraft[Count+6] = craftNo[1]
	//byteCraft[Count + 5] = bytedevNo[0];
	//byteCraft[Count + 6] = bytedevNo[1];
	for i := 0; i < namelength; i++ {
		byteCraft[Count+7+i] = byteName[i]
	}
	for j := 0; j < spaceNum; j++ {
		byteCraft[Count+7+j+namelength] = 32
	}
	byteCraft[Count+18] = 48
	byteCraft[Count+19] = 49
	byteCraft[Count+20] = 48

	strhex = util.PadLeft(util.HexSumVerify(byteCraft, Count, 21), 4, "0")
	byteCkeck = util.GetAsciiStr(strhex[len(strhex)-2:], 2, 0)
	// byteCkeck := ConvertHelper.GetAsciiStr(strhex.Substring(strhex.Length - 2, 2), 2, 0);

	byteCraft[Count+21] = byteCkeck[0]
	byteCraft[Count+22] = byteCkeck[1]

	byteCraft[Count+23] = 4

	// byte_termp = new byte[24];
	// Array.Copy(byteCraft, Count, byte_termp, 0,24);
	// nodes.Add(HxCraftNode.CreateRptEndNode(byte_termp, command.Model.WorkOrderId));
	byte_termp_end := make([]byte, 24)
	copy(byte_termp_end, byteCraft[Count:])
	nodes = append(nodes, CreateRptEndNode(byte_termp_end, 0))
	// #endregion

	return
}
