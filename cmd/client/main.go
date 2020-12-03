package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/BoRuDar/grpc-example/internal/models"
	pb "github.com/BoRuDar/grpc-example/internal/models/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(models.DefaultAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalcClient(conn)

	stream, err := c.Echo(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	go func() {
		for i := 0; true; i++ {
			msg := fmt.Sprintf("Text from client: %d", i)

			err = stream.Send(&pb.Msg{Text: msg})
			if err != nil {
				log.Fatalf("could not send: %v", err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Got from client: %s\n", in.Text)
	}
}
