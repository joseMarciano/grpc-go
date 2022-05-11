package services

import (
	"context"
	"fmt"
	"github.com/joseMarciano/grpc-go/pb/pb"
)

//type UserServiceServer interface {
//	AddUser(context.Context, *User) (*User, error)
//	mustEmbedUnimplementedUserServiceServer()
//}
//

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
