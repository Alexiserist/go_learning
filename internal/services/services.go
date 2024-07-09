package services

import (
	"go_learning/internal/models"
	"go_learning/internal/repository"
)

type TestingService struct {
	repo *repository.DatabaseRepository
}

func NewTestingService(repo *repository.DatabaseRepository) *TestingService {
	return &TestingService{repo: repo}
}

func (s *TestingService) GetAll() ([]*models.Testing, error) {
	return s.repo.GetAll()
}