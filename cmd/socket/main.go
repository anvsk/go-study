// socket编程
package main

import (
	"fmt"
	"go-ticket/cmd/socket/client"
	"go-ticket/cmd/socket/server"
	"os"
)

func main() {
	addr := make(chan string)
	go server.StartServer(addr)

	conn, err := client.Dial(<-addr)
	checkError(err)

	var buf [700]byte

	rAddr := conn.RemoteAddr()
	for {
		// 发送
		// n, err := conn.Write([]byte("Hello server"))
		// checkError(err)
		// 接收
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println("Reply form server", rAddr.String(), string(buf[0:n]))
		// time.Sleep(5 * time.Second)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
