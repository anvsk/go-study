package job

import (
	"log"
	"time"
)

type JobPing struct {
	I int
}

func (j *JobPing) Run() {
	log.Println(time.Now())
}
