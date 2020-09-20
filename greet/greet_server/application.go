package greet_server

import (
	"context"
	"fmt"
	"github.com/Anatol-e/grpcapp/greet/greetpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"math"
	"net"
	"strconv"
)

type server struct {
}

func (s *server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Infof("Greet function was invoked with %v", request)
	firstname := request.GetGreeting().GetFirstName()
	result := "Hello " + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	var msg *greetpb.GreetLongRequest
	var err error
	result := ""
	for ; err != io.EOF; msg, err = stream.Recv() {
		if err != nil {
			log.Fatalf("error when receiving stream %v ", err)
			return err
		}
		result += msg.GetGreeting().GetFirstName() + "\n"
	}
	return stream.SendAndClose(&greetpb.GreetLongResponse{
		Result: result,
	})
}

func (s *server) GreetManyTimes(
	request *greetpb.GreetManyTimesRequest,
	stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := request.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		response := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(response)
	}
	return nil
}

func (s *server) SquareRoot(ctx context.Context, req *greetpb.SquareRootRequest) (*greetpb.SquareRootResponse, error) {
	log.Info("Receiving data SquareRoot")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number : %v", number),
		)
	}
	return &greetpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
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
