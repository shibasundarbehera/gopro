# Go User API

A RESTful API built with Go and Gin framework for managing user data with comprehensive CRUD operations and interactive API documentation.

## Features

- **Complete CRUD Operations** for user management
- **Interactive API Documentation** with Swagger UI
- Clean architecture with separation of concerns
- Configuration management with environment variables
- Request logging and error handling
- JSON file-based data storage
- Health check endpoint
- Input validation and error handling
- Comprehensive API testing interface

## Project Structure

```
goAPIs/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── handlers/            # HTTP request handlers
│   ├── services/            # Business logic layer
│   ├── repositories/        # Data access layer
│   └── middleware/          # HTTP middleware
├── pkg/
│   └── utils/               # Shared utilities
├── configs/                 # Configuration management
├── data/                    # Data files
├── models/                  # Data models
├── docs/                    # Generated Swagger documentation
│   ├── docs.go             # Swagger Go code
│   ├── swagger.json        # OpenAPI JSON spec
│   └── swagger.yaml        # OpenAPI YAML spec
└── SWAGGER_README.md       # Swagger documentation guide
```

## Getting Started

### Prerequisites

- Go 1.22.1 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd goAPIs
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `8080` | Server port |
| `SERVER_HOST` | `localhost` | Server host |
| `DATA_FILE_PATH` | `data/Users.json` | Path to user data file |

## API Documentation

### Interactive Swagger UI

Access the interactive API documentation at:
```
http://localhost:8080
```

The Swagger UI provides:
- **Interactive API testing** - Test endpoints directly from the browser
- **Request/Response examples** - Pre-filled example data
- **Model schemas** - Detailed data structure documentation
- **Error responses** - Complete error documentation
- **Authentication** - Ready for JWT implementation

### Regenerating Documentation

If you modify the API, regenerate the Swagger docs:
```bash
# Install swag tool
go install github.com/swaggo/swag/cmd/swag@latest

# Add to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Generate documentation
swag init -g cmd/server/main.go -o docs
```

## API Endpoints

### Health Check
```
GET /health
```
Returns the health status of the API.

**Response:**
```json
{
  "status": "OK",
  "service": "User API",
  "version": "1.0.0"
}
```

### Get All Users
```
GET /api/v1/users
```
Returns a list of all users.

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name": "Alice",
      "email": "alice@example.com",
      "phone": "1234567890",
      "status": "active"
    }
  ],
  "count": 1
}
```

### Get User by ID
```
GET /api/v1/users/:id
```
Returns a specific user by ID.

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Alice",
    "email": "alice@example.com",
    "phone": "1234567890",
    "status": "active"
  }
}
```

### Create User
```
POST /api/v1/users
```
Creates a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "+1234567890",
  "status": "active",
  "address": "123 Main St",
  "role": "user"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 2,
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "+1234567890",
    "status": "active"
  },
  "message": "User created successfully"
}
```

### Update User
```
PUT /api/v1/users/:id
```
Updates an existing user by ID.

**Request Body:**
```json
{
  "name": "John Smith",
  "email": "john.smith@example.com",
  "phone": "+1234567890",
  "status": "active",
  "address": "456 Oak Ave",
  "role": "admin"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 2,
    "name": "John Smith",
    "email": "john.smith@example.com",
    "phone": "+1234567890",
    "status": "active"
  },
  "message": "User updated successfully"
}
```

### Delete User
```
DELETE /api/v1/users/:id
```
Deletes a user by ID.

**Response:**
```json
{
  "success": true,
  "message": "User deleted successfully"
}
```

## Data Models

### User Model
```go
type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
    Status    string `json:"status"`
    Address   string `json:"address"`
    Role      string `json:"role"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
```

### Request Models
- `CreateUserRequest` - For creating users with validation
- `UpdateUserRequest` - For updating users with validation
- `UserResponse` - For API responses

## Architecture

The application follows a clean architecture pattern:

- **Handlers**: Handle HTTP requests and responses with Swagger documentation
- **Services**: Contain business logic and data transformation
- **Repositories**: Handle data access and persistence
- **Models**: Define data structures with validation rules
- **Middleware**: Handle cross-cutting concerns (logging, error handling)
- **Utils**: Shared utilities for file operations

## Development

### Running Tests
```bash
go test ./...
```

### Building
```bash
go build -o bin/server cmd/server/main.go
```

### Code Quality
The project includes:
- **Input validation** using Gin's binding
- **Error handling** with proper HTTP status codes
- **Logging** for debugging and monitoring
- **Data persistence** with atomic file operations
- **API documentation** with Swagger annotations

## Learning Value

This project demonstrates:

- **Complete CRUD Operations** - Full Create, Read, Update, Delete functionality
- **Swagger Integration** - Professional API documentation
- **Clean Architecture** principles with proper separation of concerns
- **Dependency Injection** patterns
- **Middleware usage** for cross-cutting concerns
- **Error Handling** best practices with proper HTTP responses
- **Configuration Management** with environment variables
- **RESTful API Design** following REST conventions
- **Input Validation** and data sanitization
- **Go project structure** conventions and best practices

The goAPIs project serves as an excellent example for learning Go web development, showing how to build a production-ready API with proper documentation, validation, error handling, and maintainable code structure.

## API Testing

### Using Swagger UI
1. Start the server: `go run cmd/server/main.go`
2. Open: `http://localhost:8080`
3. Test endpoints interactively with pre-filled examples

### Using curl
```bash
# Get all users
curl http://localhost:8080/api/v1/users

# Create a user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"test@example.com","phone":"1234567890","status":"active"}'

# Get user by ID
curl http://localhost:8080/api/v1/users/1

# Update user
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated User","email":"updated@example.com","phone":"1234567890","status":"active"}'

# Delete user
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Update documentation if needed
6. Submit a pull request

## License

This project is licensed under the MIT License. 