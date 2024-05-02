# Swagger Documentation

This project includes comprehensive API documentation generated using Swagger/OpenAPI.

## Accessing the Documentation

Once the server is running, you can access the Swagger UI at:

```
http://localhost:8080
```

## API Endpoints

The following endpoints are documented:

### Health Check
- `GET /health` - Health check endpoint

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/{id}` - Get user by ID
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/{id}` - Update an existing user
- `DELETE /api/v1/users/{id}` - Delete a user

## Regenerating Documentation

If you make changes to the API endpoints or models, you need to regenerate the Swagger documentation:

```bash
# Make sure swag is installed
go install github.com/swaggo/swag/cmd/swag@latest

# Add swag to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Generate documentation
swag init -g cmd/server/main.go -o docs
```

## Swagger Annotations

The project uses Swagger annotations to document the API. Key annotations include:

### Package Level (main.go)
- `@title` - API title
- `@version` - API version
- `@description` - API description
- `@host` - API host
- `@BasePath` - Base path for all endpoints

### Function Level (handlers)
- `@Summary` - Brief description of the endpoint
- `@Description` - Detailed description
- `@Tags` - Group endpoints by tags
- `@Accept` - Expected content type
- `@Produce` - Response content type
- `@Param` - Parameter documentation
- `@Success` - Success response documentation
- `@Failure` - Error response documentation
- `@Router` - Route path and method

### Model Level (models)
- `example` - Example values for fields
- `binding` - Validation rules
- Comments - Field descriptions

## Files Generated

The `swag` tool generates the following files in the `docs/` directory:

- `docs.go` - Go code containing the Swagger specification
- `swagger.json` - OpenAPI specification in JSON format
- `swagger.yaml` - OpenAPI specification in YAML format

## Testing the API

You can use the Swagger UI to:

1. **Explore endpoints** - See all available endpoints and their documentation
2. **Test requests** - Execute API calls directly from the browser
3. **View schemas** - See the structure of request and response models
4. **Try examples** - Use pre-filled example data for testing

## Example Usage

### Creating a User
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

### Updating a User
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

## Dependencies

The following packages are used for Swagger documentation:

- `github.com/swaggo/swag` - Swagger generation tool
- `github.com/swaggo/gin-swagger` - Gin integration for Swagger
- `github.com/swaggo/files` - Static file serving for Swagger UI 