package server

import (
	"context"
	"fmt"

	"grpc-example/echo"
)

type EchoService struct {
	echo.UnimplementedEchoServer
}

func (EchoService) UnaryEcho(ctx context.Context, in *echo.EchoMessage) (*echo.EchoMessage, error) {
	fmt.Println("server receive:", in.Message)
	return &echo.EchoMessage{
		Message: "Hello from server",
	}, nil
}
