package service

import (
	"example.com/internal/api/model"
	"example.com/internal/api/repository"
)

type UserService interface {
	GetUsers() ([]model.User, error)
	GetUser(id int) (model.User, error)
	AddUser(user model.User) error
	UpdateUser(id int, user model.User) error
	DeleteUser(id int) error
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *UserServiceImpl) GetUsers() ([]model.User, error) {
	return s.UserRepository.GetUsers()
}

func (s *UserServiceImpl) GetUser(id int) (model.User, error) {
	return s.UserRepository.GetUser(id)
}

func (s *UserServiceImpl) AddUser(user model.User) error {
	return s.UserRepository.AddUser(user)
}

func (s *UserServiceImpl) UpdateUser(id int, user model.User) error {
	return s.UserRepository.UpdateUser(id, user)
}

func (s *UserServiceImpl) DeleteUser(id int) error {
	return s.UserRepository.DeleteUser(id)
}
