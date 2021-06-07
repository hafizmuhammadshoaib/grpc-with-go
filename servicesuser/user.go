package servicesuser

import (
	"example.com/grpc_with_go/models"
	repositories "example.com/grpc_with_go/repositories"
	userProto "example.com/grpc_with_go/userproto"
	"golang.org/x/net/context"
)

type Server struct {
	userProto.UnimplementedUserProtoServer
}

func (s *Server) CreateUser(ctx context.Context, in *userProto.UserCreateRequest) (*userProto.UserCreateReply, error) {
	repositories.Mgr.Create(&models.User{Name: in.GetName(), Email: in.GetEmail()})
	return &userProto.UserCreateReply{Message: "user created"}, nil
}

func (s *Server) GetUserById(ctx context.Context, in *userProto.UserGetRequest) (*userProto.UserGetReply, error) {
	user, err := repositories.Mgr.GetByID(&models.User{Id: in.GetId()})
	return &userProto.UserGetReply{Name: user.Name, Email: user.Email}, err
}
