package tf757

import (
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/ctypes"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/shared"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/enum/tf_run"
	"gitlab.szzhijing.com/quanbu/dyeing/dyeing-edge-go/svc_collector/internal/svc"
)

func (s *Tf757Socket) getDeviceStatus(ts int) int {
	switch ts {
	case tf_run.Awaiting, tf_run.Over, tf_run.Stop:
		return 0
	case tf_run.Prepare:
		return 10
	case tf_run.Running, tf_run.AutoFeed:
		return 20
	case tf_run.Pause, tf_run.Recovery, tf_run.RecoveryPara, tf_run.ManualFeed:
		return 30
	default:
		return 40
	}
}

// func (s *Tf757Socket) getSwitchOutputs() [128]bool {
// 	// List<bool> outList = new List<bool>();
// 	// outList.Add((int)statusArray[1] == 1 ? true : false);//自动
// 	// outList.Add((int)statusArray[2] == 1 ? true : false);//手动
// 	// outList.Add((int)statusArray[3] == 1 ? true : false);
// 	// outList.Add((int)statusArray[4] == 1 ? true : false);
// 	outs := [128]bool{}
// 	for i := 1; i <= 26; i++ {
// 		if s.StatusArray[i] == 1 {
// 			outs[i-1] = true
// 		}
// 	}
// 	return outs
// }

// 从排产模型获取工布
func GetStepCount(step []*ctypes.CollCraftStep) int {
	count := 0
	for _, v := range step {
		if len(v.CraftOperations) > 0 {
			count += len(v.CraftOperations)
		}
	}
	return count
}

func GetSwitchInputs(bytes []byte) [128]bool {
	ba := bytes[169:183]
	result := [128]bool{}
	for i := 0; i < 14; i++ {
		if ba[i] > 0 {
			result[i] = true
		}
	}
	return result
}

func GetSwitchOutputs(bytes []byte) [128]bool {
	ba := bytes[183:197]
	result := [128]bool{}
	for i := 0; i < 14; i++ {
		if ba[i] > 0 {
			result[i] = true
		}
	}
	return result
}

// realtime获取实际的工布
func (s *Tf757Socket) GetCurrRealStepNo(deviceStepNo int) int {
	model := s.SchedulingExecutionModelCache
	realStepNo := 0
	tempCount := 0
	if len(model.CraftStep) > 0 {
		for _, step := range model.CraftStep {
			if deviceStepNo > tempCount {
				if step.CraftOperations != nil {
					tempCount += len(step.CraftOperations)
				}
				if deviceStepNo <= tempCount {
					realStepNo = step.Sort - 1
					break
				}
			}
		}
	} else {
		realStepNo = deviceStepNo
	}
	if realStepNo > 0 {
		return realStepNo
	}
	return 0
}

// 修改工艺的时候	覆盖stepIndex下面的operation
func (s *Tf757Socket) UpdateCraftModel(updateModel *ctypes.CollUpdateCraftAttrCmdData) {
	// 服务重启后内存找不到，从filechache查找
	if s.SchedulingExecutionModelCache.CraftNames == "" || len(s.SchedulingExecutionModelCache.CraftStep) == 0 {
		if err := svc.Service.CacheFile.Get(shared.KeyCollectorTFCraft(s.dinfo.DeviceId), &s.SchedulingExecutionModelCache); err != nil {
			svc.Errorf("tf修改工艺时，从文件缓存获取不到%d", s.dinfo.StandNo)
		}
	}
	for sk, step := range s.SchedulingExecutionModelCache.CraftStep {
		if step.StepID == updateModel.StepID {
			for ok, operation := range step.CraftOperations {
				if operation.CraftOperationID == updateModel.CraftOperation.CraftOperationID {
					s.SchedulingExecutionModelCache.CraftStep[sk].CraftOperations[ok] = updateModel.CraftOperation
					goto OUT
				}
			}
		}
	}
OUT:
	return
}

// 跳步时获取工布
func (s *Tf757Socket) GetJumpStepNo(jumpStepNo int) int {
	model := s.SchedulingExecutionModelCache
	tfStepNo := 0
	tempCount := 0
	jumpStepNo += 1
	if len(model.CraftStep) > 0 {
		for _, step := range model.CraftStep {
			tempCount += 1
			if tempCount >= jumpStepNo {
				tfStepNo += 1
				break
			} else {
				if len(step.CraftOperations) > 0 {
					tfStepNo += len(step.CraftOperations)
				}
			}
		}
	} else {
		tfStepNo = jumpStepNo
	}
	return tfStepNo
}
