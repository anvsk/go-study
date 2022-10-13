package main

import (
	"io"
	"log"
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
	setLogToFile()
	go initSignal()
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				log.Println("====", ii, "=====", time.Now())
				<-time.After(time.Second)
			}
		}(i)
	}
	for {
	}
}

func setLogToFile() {
	writer, _ := os.OpenFile("wrr.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	log.SetOutput(io.MultiWriter([]io.Writer{
		writer,
		os.Stdout,
	}...))
}

func initSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGSTOP, syscall.SIGTSTP, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	files, _ := os.Open("wrrr.log")
	defer files.Close()
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Print("===queue exit")
			os.Exit(0)
			return
		case syscall.SIGSTOP, syscall.SIGTSTP:
			log.Print("===SIGSTOP")
			os.Exit(0)
		case syscall.SIGKILL:
			// log.Println()
			log.Println("===SIGKILL", time.Now())

		// TODO reload
		default:
			return
		}
	}
}
