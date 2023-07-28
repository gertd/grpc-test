package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bufbuild/protovalidate-go"
	pb "github.com/gertd/grpc-test/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	if err = v.Validate(in); err != nil {
		return nil, err
	}

	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHallo(ctx context.Context, in *pb.HalloRequest) (*pb.HalloReply, error) {
	v, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	if err = v.Validate(in); err != nil {
		return nil, err
	}

	log.Printf("Received: %v", in.GetName())
	return &pb.HalloReply{Message: "Hallo " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
