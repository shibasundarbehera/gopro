package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/goAPI/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUsers() ([]models.UserResponse, error) {
	args := m.Called()
	return args.Get(0).([]models.UserResponse), args.Error(1)
}

func (m *MockUserService) GetUserByID(id string) (*models.UserResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UserResponse), args.Error(1)
}

func TestGetUsers(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	handler := NewUserHandler(mockService)

	router := gin.New()
	router.GET("/users", handler.GetUsers)

	// Test data
	expectedUsers := []models.UserResponse{
		{Name: "Alice", Email: "alice@example.com", Phone: "1234567890", Status: "active"},
		{Name: "Bob", Email: "bob@example.com", Phone: "2345678901", Status: "inactive"},
	}

	// Setup mock expectations
	mockService.On("GetAllUsers").Return(expectedUsers, nil)

	// Create request
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	// Execute request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.True(t, response["success"].(bool))
	assert.Equal(t, float64(2), response["count"].(float64))

	mockService.AssertExpectations(t)
}
