package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/importantcoding/learning_grpc/calc/calcpb"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("I am client")
		log.Fatalf("Error from grpc.Dial : %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)
	fmt.Printf("Created client: %f", c)
	// doUnary(c)
	doServerStream(c)
}

func doUnary(c calcpb.CalcServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &calcpb.CalcRequest{
		Calc: &calcpb.Calc{
			FirstNumber:  4,
			SecondNumber: 11,
			TotalNumber:  16,
		},
	}
	res, err := c.Calc(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calc RPC: %v", err)
	}
	log.Printf("Response from Calc: %v", res.Answer)
}

func doServerStream(c calcpb.CalcServiceClient) {
	fmt.Println("Starting ServerStream RPC...")
	req := &calcpb.PrimeNumbersDecompRequest{
		Number: 120000,
	}
	res, err := c.PrimeNumbersDecomp(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calc RPC: %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while calling Calc RPC: %v", err)
		}

		log.Printf("Prime Factor: %v", msg)
	}
}
