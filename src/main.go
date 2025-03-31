package main

import (
	"root/src/lib/libblockchain/blockchainserver"
	"root/src/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	go blockchainserver.BlockchainServer()
	server.BootHTTPServer()
}
