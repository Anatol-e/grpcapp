package greet_client

import (
	"github.com/Anatol-e/grpcapp/greet/greetpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func StartApplication() {
	log.Info("Hello, I'm client")
	cc, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	log.Infof("Client was created : %f", c)
}
