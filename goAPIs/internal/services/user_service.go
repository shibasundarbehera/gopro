package services

import (
	"fmt"
	"strconv"

	"example.com/goAPI/internal/repositories"
	"example.com/goAPI/models"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Phone:  user.Phone,
			Status: user.Status,
		})
	}

	return userResponses, nil
}

func (s *UserService) GetUserByID(id string) (*models.UserResponse, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %s", id)
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &models.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Status: user.Status,
	}, nil
}

func (s *UserService) CreateUser(user *models.User) (*models.UserResponse, error) {
	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &models.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Status: user.Status,
	}, nil
}

func (s *UserService) UpdateUser(id string, user *models.User) (*models.UserResponse, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %s", id)
	}

	if err := s.userRepo.Update(userID, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &models.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Status: user.Status,
	}, nil
}

func (s *UserService) DeleteUser(id string) error {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid user ID: %s", id)
	}

	if err := s.userRepo.Delete(userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
