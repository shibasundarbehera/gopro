package models

// User represents a user in the system
type User struct {
	ID        int    `json:"id" example:"1"`                           // Unique identifier for the user
	Name      string `json:"name" example:"John Doe"`                  // Full name of the user
	Email     string `json:"email" example:"john@example.com"`         // Email address of the user
	Phone     string `json:"phone" example:"+1234567890"`              // Phone number of the user
	Status    string `json:"status" example:"active"`                  // Current status of the user
	Address   string `json:"address" example:"123 Main St"`            // Address of the user
	Role      string `json:"role" example:"user"`                      // Role of the user in the system
	CreatedAt string `json:"created_at" example:"2024-01-01 12:00:00"` // Timestamp when user was created
	UpdatedAt string `json:"updated_at" example:"2024-01-01 12:00:00"` // Timestamp when user was last updated
}

// UserResponse represents the user data returned in API responses
type UserResponse struct {
	ID     int    `json:"id" example:"1"`                   // Unique identifier for the user
	Name   string `json:"name" example:"John Doe"`          // Full name of the user
	Email  string `json:"email" example:"john@example.com"` // Email address of the user
	Phone  string `json:"phone" example:"+1234567890"`      // Phone number of the user
	Status string `json:"status" example:"active"`          // Current status of the user
}

// CreateUserRequest represents the request body for creating a new user
type CreateUserRequest struct {
	Name    string `json:"name" binding:"required" example:"John Doe"`                // Full name of the user (required)
	Email   string `json:"email" binding:"required,email" example:"john@example.com"` // Email address of the user (required)
	Phone   string `json:"phone" example:"+1234567890"`                               // Phone number of the user
	Status  string `json:"status" example:"active"`                                   // Current status of the user
	Address string `json:"address" example:"123 Main St"`                             // Address of the user
	Role    string `json:"role" example:"user"`                                       // Role of the user in the system
}

// UpdateUserRequest represents the request body for updating an existing user
type UpdateUserRequest struct {
	Name    string `json:"name" example:"John Doe"`                          // Full name of the user
	Email   string `json:"email" binding:"email" example:"john@example.com"` // Email address of the user
	Phone   string `json:"phone" example:"+1234567890"`                      // Phone number of the user
	Status  string `json:"status" example:"active"`                          // Current status of the user
	Address string `json:"address" example:"123 Main St"`                    // Address of the user
	Role    string `json:"role" example:"user"`                              // Role of the user in the system
}
