package service

import (
	"example.com/internal/api/model"
	"example.com/internal/api/repository"
)

type UserService interface {
	GetUsers() ([]model.User, error)
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
