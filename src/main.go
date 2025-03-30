package main

import (
	"root/src/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	go server.BootTCPServer()
	server.BootHTTPServer()
}
