package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/nilerajput91/RPc/grpc-server-1/model"
	"github.com/nilerajput91/RPc/user"
)

// UserRepoImpl for gorm
type UserRepoImpl struct {
	db *gorm.DB
}

// CreateUserRepoImpl for user
func CreateUserRepoImpl(db *gorm.DB) user.UserRepo {
	return &UserRepoImpl{db}
}

// AddUser for add the user in db
func (e *UserRepoImpl) AddUser(user *model.User) (*model.User, error) {

	var userDB = model.UserDB{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Password,
	}
	err := e.db.Table("user").Save(&userDB).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserById find the user by its id
func (e *UserRepoImpl) FindUserById(id model.UserId) (*model.User, error) {
	var user model.User
	err := e.db.Table("user").Where("id=?", id.Id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUsers
func (e *UserRepoImpl) FindUsers() (*[]model.User, error) {
	var users []model.User
	err := e.db.Table("user").Find(&users).Error

	if err != nil {
		return nil, err
	}
	return &users, err
}

// UpdateUser
func (e *UserRepoImpl) UpdateUser(user *model.UserUpdate) (*model.User, error) {
	var us model.User
	err := e.db.Table("user").Where("id=?", user.Id).First(&us).Update(&user.User).Error
	if err != nil {
		return nil, err
	}
	return &us, nil
}

// DeleteUser
func (e *UserRepoImpl) DeleteUser(id *model.UserId) error {
	var user model.User
	err := e.db.Table("user").Find(&user).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
