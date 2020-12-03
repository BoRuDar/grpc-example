package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "github.com/BoRuDar/grpc-example/internal/models/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const port = ":50051"

type server struct {
	pb.UnimplementedCalcServer
}

func (s *server) Echo(stream pb.Calc_EchoServer) error {
	go func() {
		for i := 0; true; i++ {
			srvMsg := &pb.Msg{Text: fmt.Sprintf("Text from server: %d", i)}

			if err := stream.Send(srvMsg); err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Printf("Got from client: %s\n", in.Text)
	}
}

func (s *server) Calculate(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Result: 0}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterCalcServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
