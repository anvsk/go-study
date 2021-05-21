// 分布式生成唯一ID、参考雪花算法
// 原理：毫秒级时间戳[范围:2^41]+机器码[2^10]+序号[2^12]组成
package util

import (
    "strconv"

    "github.com/holdno/snowFlakeByGo"
)

var globalIDWorker *snowFlakeByGo.Worker

// InitIDWorker 初始化ID生成器
func InitIDWorker(cluster int64) {
    var (
        err error
    )
    globalIDWorker, err = snowFlakeByGo.NewWorker(cluster)
    if err != nil {
        panic(err)
    }
}

// MakeOrderID 生成订单
func GetStrID() string {
    return strconv.FormatInt(globalIDWorker.GetId(), 10)
}
