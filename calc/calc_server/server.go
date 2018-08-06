package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/importantcoding/learning_grpc/calc/calcpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("calc was invoked with.... %v\n", req)
	firstNumber := req.GetCalc().GetFirstNumber()
	secondNumber := req.GetCalc().GetSecondNumber()
	totalNumber := req.GetCalc().GetTotalNumber()
	result := firstNumber + secondNumber
	var answer bool
	var res *calcpb.CalcResponse
	if result == totalNumber {
		answer = true
	} else {
		answer = false
	}
	res = &calcpb.CalcResponse{
		Result: result,
		Answer: answer,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
