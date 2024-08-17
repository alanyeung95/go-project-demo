package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	proto "github.com/alanyeung95/GoProjectDemo/api/proto"
)

type GreeterClient struct {
	client proto.GreeterClient
}

func NewGreeterClient(addr string) (*GreeterClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s: %v", addr, err)
	}

	return &GreeterClient{
		client: proto.NewGreeterClient(conn),
	}, nil
}

func (gc *GreeterClient) SayHello(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := gc.client.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		return "", fmt.Errorf("could not greet: %v", err)
	}
	return r.Message, nil
}
