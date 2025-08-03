# Go User API

A RESTful API built with Go and Gin framework for managing user data.

## Features

- RESTful API endpoints for user management
- Clean architecture with separation of concerns
- Configuration management with environment variables
- Request logging and error handling
- JSON file-based data storage
- Health check endpoint

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
└── models/                  # Data models
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

| Variable           | Default             | Description            |
| ------------------ | ------------------- | ---------------------- |
| `SERVER_PORT`    | `8080`            | Server port            |
| `SERVER_HOST`    | `localhost`       | Server host            |
| `DATA_FILE_PATH` | `data/Users.json` | Path to user data file |

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
    "name": "Alice",
    "email": "alice@example.com",
    "phone": "1234567890",
    "status": "active"
  }
}
```

## Architecture

The application follows a clean architecture pattern:

- **Handlers**: Handle HTTP requests and responses
- **Services**: Contain business logic
- **Repositories**: Handle data access
- **Models**: Define data structures
- **Middleware**: Handle cross-cutting concerns

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o bin/server cmd/server/main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.