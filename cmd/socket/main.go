// socket编程
package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"go-study/cmd/socket/server"
	"net"
	"os"
)

var LEN = 4

func main() {

	// go TestHgNativeService()
	// time.Sleep(time.Second)
	// conn, err := client.Dial(":6000")
	// checkError(err)
	go server.StartServer()
	// var buf [700]byte

	// rAddr := conn.RemoteAddr()
	// for {
	// 	time.Sleep(time.Second)
	// 	println("send...")
	// 	// 发送
	// 	// conn.Write([]byte("Hello server"))
	// 	conn.Write([]byte{4, 1, 1, 1, 1, 2, 1, 1, 3, 1, 1, 1})
	// 	conn.Write([]byte{5})
	// 	conn.Write([]byte{1, 1, 1})
	// 	conn.Write([]byte{1, 1, 6})
	// 	conn.Write([]byte{0, 0, 0, 0})
	// 	conn.Write([]byte{0, 0})
	// 	break
	// 	// checkError(err)
	// 	// 接收
	// 	// n, err := conn.Read(buf[0:])
	// 	// checkError(err)
	// 	// fmt.Println("Reply form server", rAddr.String(), buf[0:n])
	// 	// time.Sleep(5 * time.Second)
	// }
	for {
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func TestHgNativeService() {
	svc, _ := net.Listen("tcp", ":6000")
	for {
		println("comming")
		con, err := svc.Accept()
		if err != nil {
			panic("lllll")
		}
		fmt.Println("新连接:", con.RemoteAddr())
		result := bytes.NewBuffer(nil)

		go func() {
			for {
				println("get...")
				var buf [1024000000]byte
				n, err := con.Read(buf[:])
				if err != nil {
					break
				}
				result.Write(buf[0:n])
				scanner := bufio.NewScanner(result)
				scanner.Split(packetSlitFunc3)
				for scanner.Scan() {
					fmt.Println("recv:", scanner.Bytes()[:9])
				}
				result.Reset()
				// fmt.Println(buf[:n])
				// if n < 8 {
				// 	fmt.Println("continue...")
				// 	continue
				// }
				// fmt.Println(n, buf[3:5])
			}
		}()
	}
}

func packetSlitFunc3(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	if !atEOF && len(data) >= LEN {
		var l int
		// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^16)
		// binary.Read(bytes.NewReader(data[:LEN]), binary.LittleEndian, &l)
		binary.Read(bytes.NewReader(data[:LEN]), binary.LittleEndian, &l)
		// l = int16(data[0])
		println("分析长度", l)
		pl := (int(l) + LEN)
		if pl <= len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}
