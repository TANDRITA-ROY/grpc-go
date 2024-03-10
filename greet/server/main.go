package main

import (
	"context"
	pb "grpc-go/greet/proto"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello  " + in.FirstName + "!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error while listening %v\n", err)
	}
	log.Printf("Listing on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error while listening %v\n", err)
	}
}
