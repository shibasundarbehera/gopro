package repositories

import (
	"fmt"

	"example.com/goAPI/models"
	"example.com/goAPI/pkg/utils"
)

type UserRepository struct {
	filePath string
}

func NewUserRepository(filePath string) *UserRepository {
	return &UserRepository{
		filePath: filePath,
	}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	users, err := utils.LoadFromFile[models.User](r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load users from file: %w", err)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user with ID %d not found", id)
}
