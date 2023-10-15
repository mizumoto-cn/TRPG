package main

import (
	"log"
	"net"

	"github.com/mizumoto-cn/TRPcG/testing/message"

	"github.com/mizumoto-cn/TRPcG"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	server := TRPcG.NewServer()
	server.RegisterName("ArithService", new(message.ArithService))
	server.Serve(lis)
}
