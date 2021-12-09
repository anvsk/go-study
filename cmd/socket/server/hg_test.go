package server

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

var Conn net.Conn

func TestHgNativeService(*testing.T) {
	svc, err := net.Listen("tcp", ":60009")
	if err != nil {
		panic(err)
	}
	println("开启成功服务")
	connected := make(chan net.Conn, 1)
	go func() {
		<-time.After(time.Second)
		conn, err := net.Dial("tcp", ":60009")
		if err != nil {
			panic(err)
		}
		println("连接成功服务")

		// for {
		// 	rand.Seed(time.Now().UnixNano())
		// 	conn.Write(make([]byte, rand.Intn(10)))
		// 	time.Sleep(time.Second)
		// }
		// for {
		for i := 0; i < 1; i++ {

			// time.Sleep(time.Second)
			buf := []byte{4, 1, 1, 1, 1, 2, 1, 1, 3, 1, 1, 1, 1}

			// type1
			for _, v := range buf {
				time.Sleep(time.Second)
				conn.Write([]byte{v})
			}

			// type2
			// conn.Write(buf)

		}
	}()

	for {
		fmt.Println("准备连接")
		con, err := svc.Accept()
		if err != nil {
			panic(err)
		}
		println("-----")
		fmt.Println("新连接:", con.RemoteAddr())
		connected <- con
		result := bytes.NewBuffer(nil)
		scanner := bufio.NewScanner(result)
		scanner.Split(packetSlitFunc2)
		go func() {
			var buftmp [10240]byte
			for {
				// println("开始接收")
				n, err := con.Read(buftmp[0:])
				// fmt.Println("接收到:", n, buftmp[0:n])
				n, err = result.Write(buftmp[:n])
				if err != nil {
					panic("result.Write!")
				}
				if err != nil {
					if err == io.EOF {
						continue
					} else {
						fmt.Println("read err:", err)
						break
					}
				} else {
					for scanner.Scan() {
						fmt.Println("解析出:", len(scanner.Bytes()), scanner.Bytes())
					}
					fmt.Println("readed!")
				}

				result.Reset()
			}
		}()
	}
}

func packetSlitFunc2(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	// if !atEOF && len(data) >= 1 {
	// 	var l int16
	// 	// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^16)
	// 	// binary.Read(bytes.NewReader(data[:1]), binary.LittleEndian, &l)
	// 	l = int16(data[:1][0])
	// 	println("分析长度:", l)
	// 	pl := int(l) + 1
	// 	if pl <= len(data) {
	// 		return pl, data[:pl], nil
	// 	}
	// }
	return int(data[0]) + 1, data[:1], nil
	// return
}

func packetSlitFunc3(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	if !atEOF && len(data) >= 1 {
		var l int16
		// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^16)
		// binary.Read(bytes.NewReader(data[:1]), binary.LittleEndian, &l)
		l = int16(data[0])
		println("分析长度", l)
		pl := int(l) + 1
		if pl <= len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}

func TestA(*testing.T) {
	// An artificial input source.
	const input = "412343678"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
