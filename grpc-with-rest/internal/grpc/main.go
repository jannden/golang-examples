package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/jannden/golang-examples/grpc-with-rest/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (*server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := request.Name
	response := &pb.HelloResponse{
		Message: "Hello " + name,
	}
	return response, nil
}



func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}