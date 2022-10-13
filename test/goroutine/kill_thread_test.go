package goroutine

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestXxxxxxx(t *testing.T) {

	// io.WriteString(os.)
}

func main() {
	do()
}

func do() {
	go initSignal()
	for {
		ioutil.WriteFile("wrrr.log", []byte(fmt.Sprintln(time.Now())), 0644)
		<-time.After(time.Second)
	}
}

func initSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Print("queue exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
