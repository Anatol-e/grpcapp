package main

import (
	"context"
	"github.com/Anatol-e/grpcapp/blog/blogpb"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type Blog struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
}

type server struct {
}

func (s *server) CreateBlog(ctx context.Context, request *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	panic("implement me")
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err != nil {
		log.Fatalf("Cant connect to the mongo with an err=%v", err)
	}
	err = client.Connect(ctx)

	blog := &Blog{
		Id:   1,
		Text: "Blog text Blog text Blog text Blog text ",
	}
	collection := client.Database("mg_db_rpc").Collection("blog")
	_, err = collection.InsertOne(ctx, blog)
	if err != nil {
		log.Fatalf("Can't insert row %v", err)
	}

	result, err := collection.Find(ctx, bson.D{{"id", 1}})
	if err != nil {
		log.Fatalf("Can't find data %v", err)
	}
	result.Decode(&blog)
	log.Infof("coll %v", *blog)

	// SSL
	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("Failing loading certificates : %v", sslErr)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	blogpb.RegisterBlogServiceServer(s, &server{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Info("Listening...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
