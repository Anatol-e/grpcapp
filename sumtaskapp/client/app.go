package main

import (
	"context"
	"github.com/Anatol-e/grpcapp/sumtaskapp/sumpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Info("Hello, I'm client")
	cc, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)

	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			FirstNumber:  10,
			SecondNumber: 20,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC : %v", err)
	}
	log.Infof("Response from Greet : %v", res.SumResult)
}
