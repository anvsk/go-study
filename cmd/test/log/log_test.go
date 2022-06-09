package log

import (
	// "log"

	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"

	"testing"
)

func TestArguDefer2(*testing.T) {
	// log.SetReportCaller(true)
	// log.SetLevel(5)
	// // fr:=io.Writer
	// log.Println("prinln")
	// log.Printf("prinf")

	// logger := logrus.New()
	// logger.Formatter = &logrus.JSONFormatter{}

	// Use logrus for standard log output
	// Note that `log` here references stdlib's log
	// Not logrus imported under the name `log`.

	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("test-log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("log file open error")
	}
	//同时写文件和流
	log.SetOutput(io.MultiWriter([]io.Writer{file, os.Stdout}...))
	for i := 0; i < 10; i++ {
		go func() {
			for {
				log.Info(time.Now())
				time.Sleep(50 * time.Millisecond)
			}
		}()
	}
	for {
	}
}

// `WithLinkName` 为最新的日志建立软连接
// `WithRotationTime` 设置日志分割的时间，隔多久分割一次
// `WithMaxAge 和 WithRotationCount二者只能设置一个
// `WithMaxAge` 设置文件清理前的最长保存时间
// `WithRotationCount` 设置文件清理前最多保存的个数
func TestXxx(t *testing.T) {
	// json格式化
	// log.SetFormatter(&log.JSONFormatter{})

	// 下面demo配置日志每隔 2s 轮转一个新文件，保留最近 10s 的日志文件，多余的自动清理掉。
	// writer, err := rotatelogs.New(
	// "%Y%m%d%H%M%S.log",
	// rotatelogs.WithMaxAge(time.Duration(10)*time.Second),
	// rotatelogs.WithRotationTime(time.Duration(2)*time.Second),
	// )
	writer, err := rotatelogs.New(
		"tmp3/%Y%m%d%H%M%S.log",
		rotatelogs.WithLinkName("tmp3/ruanlink.log"),
		rotatelogs.WithMaxAge(time.Duration(1)*time.Minute),
		rotatelogs.WithRotationTime(time.Duration(1)*time.Second),
	)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter([]io.Writer{writer, os.Stdout}...))
	for i := 0; i < 10; i++ {
		go func() {
			for {
				log.Info(time.Now())
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}
	for {
	}
}
