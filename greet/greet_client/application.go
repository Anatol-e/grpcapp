package greet_client

import (
	"context"
	"github.com/Anatol-e/grpcapp/greet/greetpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
)

func StartApplication() {
	// No SSL
	// cc, err := grpc.Dial(":8080", grpc.WithInsecure())

	// SSL
	certFile := "ssl/ca.crt"
	creds, sslError := credentials.NewClientTLSFromFile(certFile, "localhost")
	if sslError != nil {
		log.Fatalf("Failing loading certificates : %v", sslError)
	}
	cc, err := grpc.Dial(":8080", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
	doErrorUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Anatol",
			LastName:  "Kozhukhar",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC : %v", err)
	}
	log.Infof("Response from Greet : %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Anatol",
			LastName:  "Kozhukhar",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes rpc : %v", err)
	}

	var msg *greetpb.GreetManyTimesResponse
	for ; err != io.EOF; msg, err = resStream.Recv() {
		if err != nil {
			log.Fatalf("error while reading stream : %v", err)
			break
		}
		log.Infof("Response from GreetManyTimes : %v ", msg)
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetLongRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Anatol",
			LastName:  "Kozhukhar",
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("can't make long greet stream with err : %v ", err)
	}
	for i := 0; i < 10; i++ {
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("can't get response %v ", err)
	}

	log.Infof("Result Client Streaming : %v", res.Result)
}

func doErrorUnary(c greetpb.GreetServiceClient) {
	log.Info("Starting do error unary rpc")

	// correct call
	res, err := c.SquareRoot(context.Background(), &greetpb.SquareRootRequest{
		Number: 25,
	})
	if err != nil {
		log.Infof("response finished with an error %v", err)
	} else {
		log.Infof("Response result : %v ", res.NumberRoot)
	}

	// error call
	res, err = c.SquareRoot(context.Background(), &greetpb.SquareRootRequest{
		Number: -25,
	})
	if err != nil {
		log.Infof("response finished with an error %v", err)
	} else {
		log.Infof("Response result : %v ", res.NumberRoot)
	}
}
