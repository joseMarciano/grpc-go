package main

import (
	"context"
	"fmt"
	"github.com/joseMarciano/grpc-go/pb/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "José",
		Email: "jose@jose.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)

}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "José",
		Email: "jose@jose.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive gRPC stream request: %v", err)
		}

		fmt.Println("Status: ", stream.Status)

	}

}
