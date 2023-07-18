package client

import (
	"context"
	"fmt"
	"log"

	"grpc-example/echo"
)

func CallUnaryEcho(c echo.EchoClient) {
	ctx := context.Background()
	in := echo.EchoMessage{
		Message: "Hello from client",
	}

	res, err := c.UnaryEcho(ctx, &in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client receive:", res.Message)
}
