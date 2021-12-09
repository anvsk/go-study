package server

import (
	"fmt"
	"log"
	"net"
	"time"
)

var realtimedata = []byte("000100000065010362000000010000000000000000000000000001000000000000000000010000000000000000000000000000000000000000000000000000000500000000000002c800000a1501be00000000000000000000000000000000000000000000000000000000")
var realtime_resp = []byte{0x00, 0x01, 0x00, 0x00, 0x00, 0x65, 0x01, 0x03, 0x62, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xc8, 0x00, 0x00, 0x0a, 0x15, 0x01, 0xbe, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

func StartServer() {

	l, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		fmt.Println(time.Now(), conn.RemoteAddr(), "connected... ...")
		go serveReceive(conn)
		go servePush(conn)
	}
}

func serveReceive(conn net.Conn) {
	// sending := new(sync.Mutex) // make sure to send a complete response
	// wg := new(sync.WaitGroup)  // wait until all request are handled
	defer conn.Close()
	var buf [1024]byte
	realtimereq := []byte{0, 1, 0, 0, 0, 6, 1, 3, 0, 0, 0, 49}
	for {
		// 接收数据
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), buf[0:n])

		// rec := string(buf[0:n])
		if string(buf[0:n]) == string(realtimereq) {
			conn.Write(realtime_resp)
		} else {
			fmt.Println("other--Receive from client", rAddr.String(), string(buf[0:n]))
			// <-time.After(200 * time.Millisecond)
			conn.Write(realtime_resp)
		}

	}
}

func servePush(conn net.Conn) {
	i := 0
	for {
		fmt.Println("pushed")
		i++
		conn.Write([]byte(fmt.Sprintf("%d", i)))
		<-time.After(1000 * time.Millisecond)
	}

}
