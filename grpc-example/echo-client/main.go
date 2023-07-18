package main

import (
	"log"

	"google.golang.org/grpc"

	"grpc-example/echo"
	"grpc-example/echo-client/client"
	"grpc-example/etcd"
)

var (
	// addr        = flag.String("addr", "localhost:20010", "")
	serviceName = "echo-server"
)

func main() {
	// flag.Parse()
	addr, err := etcd.CusServiceDiscovery(serviceName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("get service: %s, address: %s", serviceName, addr)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := echo.NewEchoClient(conn)
	client.CallUnaryEcho(c)
}
