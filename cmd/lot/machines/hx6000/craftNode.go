package hx6000

type HxCraftNode struct {
	WorkOrderId        int64
	CraftNodeBytes     []byte
	IsVerify           bool
	IsRptHeader        bool
	IsRptEnd           bool
	VerStandNo         string
	VerCraftName       string
	VerCraftCode       string
	VerCraftStepIndex  int
	VerParameterValue1 string
	VerParameterValue2 string
	VerParameterValue3 string
	CraftNo            int
}

func CreateRptHeaderNode(bytes []byte, workOrderId int64, CraftNo int) HxCraftNode {
	return HxCraftNode{
		CraftNodeBytes: bytes,
		WorkOrderId:    workOrderId,
		IsRptHeader:    true,
		CraftNo:        CraftNo,
	}
}

func CreateNode(bytes []byte, workOrderId int64, step int) HxCraftNode {
	return HxCraftNode{
		CraftNodeBytes:    bytes,
		WorkOrderId:       workOrderId,
		VerCraftStepIndex: step,
	}
}

func CreateRptEndNode(bytes []byte, workOrderId int64) HxCraftNode {
	return HxCraftNode{
		CraftNodeBytes: bytes,
		WorkOrderId:    workOrderId,
		IsRptEnd:       true,
	}
}
