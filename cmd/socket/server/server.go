package server

import (
	"fmt"
	"log"
	"net"
)

func StartServer(addr chan string) {

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go serveReceive(conn)
		go servePush(conn)
	}
}

func serveReceive(conn net.Conn) {
	// sending := new(sync.Mutex) // make sure to send a complete response
	// wg := new(sync.WaitGroup)  // wait until all request are handled
	defer conn.Close()
	var buf [1024]byte
	for {
		// 接收数据
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("Welcome client"))
		if err2 != nil {
			return
		}
	}
}

func servePush(conn net.Conn) {
	i := 0
	for {
		i++
		conn.Write([]byte(fmt.Sprintf("[__%d__]", i)))
		// <-time.After(100 * time.Millisecond)
	}

}
