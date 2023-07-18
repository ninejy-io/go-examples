package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-example/echo"
	"grpc-example/echo-server/server"
	"grpc-example/etcd"
)

var (
	serviceName = "echo-server"
	serviceHost = "localhost"
	port        = flag.Int("port", 20010, "")
)

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	echo.RegisterEchoServer(s, server.EchoService{})

	etcd.CusServiceRegister(serviceName, fmt.Sprintf("%s:%d", serviceHost, *port))

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
