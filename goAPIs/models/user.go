package models

type User struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Status    string
	Address   string
	Role      string
	CreatedAt string
	UpdatedAt string
}

type UserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Status string `json:"status"`
}
