package blockchainserver

import (
	"fmt"
	"log"
	"os"

	"github.com/panjf2000/gnet/v2"
)

var BLKCHAIN_PORT string

type blockChainServer struct {
	gnet.BuiltinEventEngine
}

func (s *blockChainServer) OnBoot(engine gnet.Engine) (action gnet.Action) {
	fmt.Println("BlockChain Server started on :" + BLKCHAIN_PORT + "\n")
	return
}

func (s *blockChainServer) OnAction(conn gnet.Conn) (action *gnet.Action) {
	buf, _ := conn.Next(-1)
	fmt.Printf("Received: %s\n", &buf)
	conn.AsyncWrite([]byte("Hello from Blockchain Server"+"\n"), nil)
	return
}

func BlockchainServer() {
	BLKCHAIN_PORT = os.Getenv("BLKCHAIN_PORT")
	server := new(blockChainServer)
	if err := gnet.Run(server, "tcp://:"+BLKCHAIN_PORT, gnet.WithMulticore(true)); err != nil {
		log.Fatal("TCP Server failed:", err)
	}
}
