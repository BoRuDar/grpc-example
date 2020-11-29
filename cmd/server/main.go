package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/BoRuDar/grpc-example/internal/models/api"

	"google.golang.org/grpc"
)

const port = ":50051"

type server struct {
	pb.UnimplementedCalcServer
}

func (s *server) Calculate(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %+v", in)

	var result float32

	switch in.Op {
	case pb.OP_ADD:
		result = in.A + in.B

	case pb.OP_MUL:
		result = in.A * in.B

	default:
		return nil, fmt.Errorf("OP[%v] is not implemented", in.Op)
	}

	return &pb.Response{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalcServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
