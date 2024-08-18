package grpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"

	pb "github.com/alanyeung95/GoProjectDemo/api/proto/helloworld"
)

type GreeterClient struct {
	client pb.GreeterClient
}

func NewGreeterClient(addr string) (*GreeterClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s: %v", addr, err)
	}

	return &GreeterClient{
		client: pb.NewGreeterClient(conn),
	}, nil
}

func (gc *GreeterClient) SayHello(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := gc.client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return "", fmt.Errorf("could not greet: %v", err)
	}
	return r.Message, nil
}

/*
func (gc *GreeterClient) LotsOfGreetings() (string, error) {
	// Create a stream for sending multiple greeting requests
	stream, err := gc.client.LotsOfGreetings(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to create stream: %v", err)
	}

	// Iterate over the slice of names and send each as a request
	for i := 0; i < 1000; i++ {
		req := &pb.HelloRequest{Name: string(i)}
		if err := stream.Send(req); err != nil {
			// Close the stream in case of an error
			stream.CloseSend()
			return "", fmt.Errorf("failed to send request: %v", err)
		}
	}

	// Close the stream and receive the final response
	reply, err := stream.CloseAndRecv()
	if err != nil {
		return "", fmt.Errorf("failed to receive response: %v", err)
	}

	return reply.Message, nil
}
*/

func (gc *GreeterClient) SayHelloThousandTimes(name string) ([]string, error) {
	var count = 10000
	responses := make(chan string, count)

	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			// create new context for isolation of timeout
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			response, err := gc.client.SayHello(ctx, &pb.HelloRequest{Name: name})
			if err != nil {
				log.Printf("Error greeting %s: %v", name, err)
				return
			}
			responses <- response.Message
		}(name)
	}

	wg.Wait()
	close(responses)

	var aggregatedResponses []string
	for response := range responses {
		aggregatedResponses = append(aggregatedResponses, response)
	}

	return aggregatedResponses, nil
}
