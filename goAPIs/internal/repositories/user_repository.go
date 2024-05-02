package repositories

import (
	"fmt"
	"time"

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

func (r *UserRepository) Create(user *models.User) error {
	users, err := r.GetAll()
	if err != nil {
		return err
	}

	// Generate new ID (simple increment)
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1

	// Set timestamps
	now := time.Now().Format("2006-01-02 15:04:05")
	user.CreatedAt = now
	user.UpdatedAt = now

	users = append(users, *user)
	return utils.SaveToFile(r.filePath, users)
}

func (r *UserRepository) Update(id int, user *models.User) error {
	users, err := r.GetAll()
	if err != nil {
		return err
	}

	for i, u := range users {
		if u.ID == id {
			user.ID = id
			user.CreatedAt = u.CreatedAt
			user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			users[i] = *user
			return utils.SaveToFile(r.filePath, users)
		}
	}

	return fmt.Errorf("user with ID %d not found", id)
}

func (r *UserRepository) Delete(id int) error {
	users, err := r.GetAll()
	if err != nil {
		return err
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return utils.SaveToFile(r.filePath, users)
		}
	}

	return fmt.Errorf("user with ID %d not found", id)
}
