package main

import (
	"context"
	"fmt"
	"github.com/joseMarciano/grpc-go/pb/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	//AddUserVerbose(client)
	AddUsers(client)
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

func AddUsers(client pb.UserServiceClient) {

	usersToSend := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "User 1",
			Email: "Email 1",
		},
		&pb.User{
			Id:    "2",
			Name:  "User 2",
			Email: "Email 2",
		},
		&pb.User{
			Id:    "3",
			Name:  "User 3",
			Email: "Email 3",
		},
		&pb.User{
			Id:    "4",
			Name:  "User 4",
			Email: "Email 4",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, user := range usersToSend {
		err := stream.Send(user)
		if err != nil {
			log.Fatalf("Error to send request: %v", err)
		}

		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error on receiving response: %v", err)
	}

	fmt.Println(res)

}
