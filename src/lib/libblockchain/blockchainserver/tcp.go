package blockchainserver

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/panjf2000/gnet/v2"
)

var BLKCHAIN_PORT string

type blockChainServer struct {
	gnet.BuiltinEventEngine
	engine *gnet.Engine
	wg     sync.WaitGroup // ✅ Ensures we wait for shutdown
}

func (s *blockChainServer) OnBoot(engine gnet.Engine) gnet.Action {
	s.engine = &engine
	fmt.Println("BlockChain Server started on :" + BLKCHAIN_PORT)
	return gnet.None
}

func (s *blockChainServer) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Printf("New connection from: %s\n", conn.RemoteAddr())
	return nil, gnet.None
}

func (s *blockChainServer) OnClose(conn gnet.Conn, err error) gnet.Action {
	if err != nil {
		fmt.Printf("Connection closed with error: %v\n", err)
	} else {
		fmt.Printf("Connection closed: %s\n", conn.RemoteAddr())
	}
	return gnet.None
}

func (s *blockChainServer) OnTraffic(conn gnet.Conn) gnet.Action {
	buf, _ := conn.Next(-1)
	fmt.Printf("Received: %s\n", buf)
	conn.AsyncWrite([]byte("Hello from Blockchain Server\n"), nil)
	return gnet.None
}

func BlockchainServer() {
	BLKCHAIN_PORT = os.Getenv("GWT_BLKCHAIN_PORT")
	if BLKCHAIN_PORT == "" {
		log.Fatal("GWT_BLKCHAIN_PORT is not set")
	}

	server := &blockChainServer{}
	server.wg.Add(1) // Block exit until server stops

	// Start the server in a goroutine
	go func() {
		defer server.wg.Done()
		port, _ := strconv.Atoi(BLKCHAIN_PORT)
		for {
			addr := fmt.Sprintf("tcp://:%d", port)
			err := gnet.Run(server, addr, gnet.WithMulticore(true))
			if err != nil && strings.Contains(err.Error(), "address already in use") {
				port++
				continue
			} else if err != nil {
				log.Fatal("TCP Server failed:", err)
			}
			break
		}
	}()

	// Capture SIGINT/SIGTERM for graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig // Wait for interrupt

	fmt.Println("Shutting down blockchain server....")
	if server.engine != nil {
		server.engine.Stop(context.Background())
	}
	server.wg.Wait() // Ensure cleanup before exiting
	fmt.Println("Server shutdown complete.")
}
