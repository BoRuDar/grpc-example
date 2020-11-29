package main

import (
	"context"
	"log"
	"os"
	"strconv"
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

	if len(os.Args) != 4 {
		log.Fatalf("expected number of arguments: 3")
	}
	a, _ := strconv.ParseFloat(os.Args[2], 32)
	b, _ := strconv.ParseFloat(os.Args[3], 32)
	var op pb.OP
	switch os.Args[1] {
	case "add":
		op = pb.OP_ADD
	case "mul":
		op = pb.OP_MUL
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Calculate(ctx, &pb.Request{Op: op, A: float32(a), B: float32(b)})
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}

	log.Printf("Result: %f", r.GetResult())
}
