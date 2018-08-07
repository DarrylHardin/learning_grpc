package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

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

func (*server) PrimeNumbersDecomp(req *calcpb.PrimeNumbersDecompRequest, stream calcpb.CalcService_PrimeNumbersDecompServer) error {
	fmt.Printf("PrimeNumbersDecomp was invoked with %v\n", req)
	firstNumber := req.GetNumber()
	var k int32
	k = 2
	n := firstNumber
	for n > 1 {
		if n%k == 0 { // if k evenly divides into N
			var resStream *calcpb.PrimeNumbersDecompResponse
			resStream = &calcpb.PrimeNumbersDecompResponse{
				Answer: k,
			}
			stream.Send(resStream)
			time.Sleep(1000 * time.Millisecond)
			n = n / k // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}

	}
	return nil
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
