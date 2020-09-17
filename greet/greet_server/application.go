package greet_server

import (
	"github.com/Anatol-e/grpcapp/greet/greetpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type server struct {
}

func StartApplication() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	log.Info("Listening...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
