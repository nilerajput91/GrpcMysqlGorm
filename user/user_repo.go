package user

import "github.com/nilerajput91/RPc/grpc-server-1/model"

// UserRepo interface
type UserRepo interface {
	AddUser(user *model.User) (*model.User, error)
	FindUserById(id model.UserId) (*model.User, error)
	FindUsers() (*[]model.User, error)
	UpdateUser(user *model.UserUpdate) (*model.User, error)
	DeleteUser(id *model.UserId) error
}
