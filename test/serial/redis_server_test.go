package com_serial

import (
	"fmt"
	"net"
	"testing"
)

// 测试redis的序列化协议
func TestMain(t *testing.T) {
	l, _ := net.Listen("tcp", "127.0.0.1:63790")
	fmt.Println("wait connect...")
	for {
		conn, _ := l.Accept()
		go func(conn net.Conn) {
			fmt.Println("get one connect!", conn.LocalAddr().String())
			buf := make([]byte, 1024)
			for {
				_, err := conn.Read(buf)
				if err != nil {
					fmt.Println("close conn:", err)
					return
				}
				fmt.Println("read:", string(buf))
				conn.Write([]byte("+OK\r\n"))
			}
		}(conn)
	}

}
