package diy

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
	"testing"
)

func TestLower(t *testing.T) {
	// fmt.Println(strings.ToLower("0x5C0401e81Bc07Ca70fAD469b451682c0d747Ef1c"))

	// s, err := strconv.ParseInt("0xaf4251", 16, 64)
	// if err != nil {
	// 	panic(err)
	// }
	// s := "TZ=Asia   09:08 \n 444\r 8908a90"
	// fmt.Println(strings.Fields(s)[2])
	// t.
	var err error
	defer func() {
		if err != nil {
			t.Log(err.Error())
		} else {
			t.Log("hahaah")
		}
	}()
	// if true {
	// 	_, err = makeerr()
	// 	// _ = nn
	// 	// err = errtmp
	// }
	if _, err = makeerr(); err != nil {

	}

}

func TestS(*testing.T) {
	s := StackToString(0, 9)
	fmt.Println(s)
}

func StackToString(startStack, count int) string {
	var sb strings.Builder
	//函数堆栈信息
	for i := startStack; i < startStack+count; i++ {
		pc, file, line, ok := runtime.Caller(i)
		//
		pcName := runtime.FuncForPC(pc).Name()

		var str string
		if ok { //文件名 行数
			str = fmt.Sprintf("%s()\n%s:%d", pcName, file, line)
		} else {
			str = "None"
		}
		// 折叠??
		if str != "None" {
			if i > startStack {
				sb.WriteString("\n")
			}

			sb.WriteString(str)
		} else {
			break
		}
	}
	return sb.String()
}

func TestLowerError(t *testing.T) {
	if err := LendStatic(); err != nil {
		t.Log(err.Error())
	} else {
		t.Log("no error")
	}
}

func makeerr() (n int, err error) {
	// err = errors.New("iiiiiiii")
	return
}

func LendStatic() (err error) {
	defer func() {
		if err := recover(); err != nil {
			// log.Println(err)
			err = errors.New(fmt.Sprintln(err))
			log.Println("come")
			return
		}
	}()
	_ = err
	err = errors.New("xxxx")
	panic("ssss")
	return
}
