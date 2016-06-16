package network

import (
	"log"
)

type HttpServer struct{
	Addr string
}

func (server *HttpServer) Start() {
	server.init()
}

func (server *HttpServer)  init() {
	log.Println("Http init")
}