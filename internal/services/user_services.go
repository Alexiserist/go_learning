package services

import (
	"go_learning/internal/models"
	"go_learning/internal/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User,error);
	CreateUser(user models.User) (models.User, error);
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService{
	return &userService{
		userRepository: repo,
	}
}

func (s *userService) GetAllUsers() ([]models.User,error){
	return s.userRepository.FindAll()
}

func (s *userService) CreateUser(user models.User) (models.User, error){
	return s.userRepository.CreateUser(user)
}