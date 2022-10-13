package demo1

import (
	"fmt"
	"go-study/cmd/cron/job"
	"log"
	"runtime"

	"github.com/robfig/cron/v3"
)

// 测试WithChain里的各种warpper
func Demo1() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			// cron.Recover(cron.DefaultLogger),
			// cron.SkipIfStillRunning(cron.DefaultLogger),
			RecoverAndSkipIfStillRunning(cron.DefaultLogger),
		),
	)
	log.Println("init cron")
	// c.AddJob("* * * * * *", &job.Job1{I: 1})
	c.AddJob("@every 1m", &job.JobPing{})

	c.Start()

}

// 自定义warpper、在recover时继续跑、原生方法recover时chan没释放，SkipIfStillRunning后下一个进程无法开始

func RecoverAndSkipIfStillRunning(logger cron.Logger) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var ch = make(chan struct{}, 1)
		ch <- struct{}{}
		return cron.FuncJob(func() {
			var v struct{}
			defer func() {
				if r := recover(); r != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					logger.Error(err, "panic", "stack", "...\n"+string(buf))
					ch <- v
				}
			}()
			select {
			case v = <-ch:
				j.Run()
				ch <- v
			default:
				logger.Info("skip")
			}
		})
	}
}
