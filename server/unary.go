package main

import (
	"context"

	pb "github.com/Jose-Tom/go-grpc-demo/proto"
)

func (s *helloServer) SayHelloUnary(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
