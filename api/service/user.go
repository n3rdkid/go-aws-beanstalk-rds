package service

import (
	"errors"
	"go-aws-beanstalk-rds/api/repository"
	"go-aws-beanstalk-rds/models"
)

// UserService ->
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService -> returns new user service
func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		userRepository: repo,
	}
}

// Save ->
func (s *UserService) Save(user models.User) (models.User, error) {
	_, err := s.userRepository.GetByID(user.ID)
	if err == nil {
		return user, errors.New("user already exists")
	}
	return s.userRepository.Save(user)
}

// GetAll ->
func (s *UserService) GetAll() (users []models.User, err error) {
	return s.userRepository.GetAll()
}

// GetByID ->
func (s *UserService) GetByID(id uint) (models.User, error) {
	return s.userRepository.GetByID(id)
}
