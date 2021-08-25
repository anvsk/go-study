package client

import (
	"fmt"
	"net"
	"os"
)

func Dial(addr string) (*net.TCPConn, error) {
	// 绑定
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	checkError(err)
	// 连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	return conn, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
