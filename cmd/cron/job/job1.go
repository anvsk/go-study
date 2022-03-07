package job

import (
	"log"
	"time"
)

type Job1 struct {
	I int
}

func (j *Job1) Run() {
	defer func() { j.I++ }()
	log.Println("Job1 running...", j.I)
	if j.I%5 == 0 {
		time.Sleep(5 * time.Second)
	}
	if j.I%10 == 0 {
		aa()
	}
}

func aa() {
	panic("xxxx")
}
