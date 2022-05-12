package services

import (
	"context"
	"fmt"
	"github.com/joseMarciano/grpc-go/pb/pb"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUseService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert - Database
	fmt.Println("Inserting.... " + req.GetName())

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	err := stream.Send(
		&pb.UserResultStream{
			Status: "Init",
			User:   &pb.User{},
		},
	)

	if err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	err = stream.Send(
		&pb.UserResultStream{
			Status: "Inserting on database",
			User:   &pb.User{},
		},
	)

	if err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	err = stream.Send(
		&pb.UserResultStream{
			Status: "User has been inserted",
			User: &pb.User{
				Id:    "123",
				Name:  req.GetName(),
				Email: req.GetEmail(),
			},
		},
	)

	if err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	err = stream.Send(
		&pb.UserResultStream{
			Status: "Process finished",
			User: &pb.User{
				Id:    "123",
				Name:  req.GetName(),
				Email: req.GetEmail(),
			},
		},
	)

	if err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	return nil
}
