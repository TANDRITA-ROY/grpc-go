package main

import (
	"context"
	"log"

	pb "grpc-go/greet/proto"

	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func DoGreet(c pb.GreetServiceClient) {

	log.Println("doGreet is running")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Rimiiiii",
	})
	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect%v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	DoGreet(c)

}
