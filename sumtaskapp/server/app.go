package main

import (
	"context"
	"github.com/Anatol-e/grpcapp/sumtaskapp/sumpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, request *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	log.Infof("Greet function was invoked with %v", request)
	requestSum := request.GetSum()
	firstNumber := requestSum.GetFirstNumber()
	secondNumber := requestSum.GetSecondNumber()
	result := firstNumber + secondNumber
	res := &sumpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	log.Info("Listening...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
