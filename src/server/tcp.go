package server

import (
	"fmt"
	"log"
	"os"

	"github.com/panjf2000/gnet/v2"
)

var TCP_PORT = os.Getenv("TCP_PORT")

type tcpServer struct {
	gnet.BuiltinEventEngine
}

func (s *tcpServer) OnBoot(engine gnet.Engine) (action gnet.Action) {
	fmt.Println("TCP server started on :" + TCP_PORT + "\n")
	return
}

func (s *tcpServer) OnAction(conn gnet.Conn) (action *gnet.Action) {
	buf, _ := conn.Next(-1)
	fmt.Printf("Received: %s\n", &buf)
	conn.AsyncWrite([]byte("Hello from gnet tcp server\n"), nil)
	return
}

func BootTCPServer() {
	server := new(tcpServer)
	if err := gnet.Run(server, "tcp://:"+TCP_PORT, gnet.WithMulticore(true)); err != nil {
		log.Fatal("TCP Server failed:", err)
	}
}
