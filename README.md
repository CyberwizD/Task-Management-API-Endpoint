# Task Management API

A simple RESTful API for task management built with Go and Gin framework. This API provides endpoints for creating, retrieving, and managing tasks with proper validation, error handling, and CORS support.

## Features

- ✅ RESTful API design with proper HTTP methods and status codes
- ✅ Robust input validation and error handling
- ✅ Thread-safe in-memory data storage
- ✅ UUID-based unique ID generation
- ✅ CORS support for cross-origin requests
- ✅ Comprehensive API documentation
- ✅ Clean project structure and code organization

## Technology Stack

- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Data Storage**: In-memory (thread-safe with mutex)
- **Dependencies**: 
  - `github.com/gin-gonic/gin` - Web framework
  - `github.com/gin-contrib/cors` - CORS middleware
  - `github.com/google/uuid` - UUID generation

## Project Structure

```
task-management-api/
├── main.go                              # Application entry point
├── go.mod                               # Go module dependencies
├── go.sum                               # Go module checksums
├── .gitignore                           # Git ignore file
├── test_examples.sh                     # API test script
├── README.md                            # This documentation
└── internal/                            # Internal application code
    ├── config/
    │   └── config.go                    # Configuration management
    ├── models/
    │   ├── task.go                      # Task model and DTOs
    │   └── errors.go                    # Error definitions
    ├── repository/
    │   └── task_repository.go           # Data access layer
    ├── services/
    │   └── task_service.go              # Business logic layer
    ├── handlers/
    │   ├── task_handler.go              # HTTP handlers for tasks
    │   └── health_handler.go            # Health check handler
    └── middleware/
        ├── cors.go                      # CORS middleware
        ├── logger.go                    # Request logging middleware
        └── error_handler.go             # Error handling middleware
```

## Installation and Setup

### Prerequisites

- Go 1.21 or higher installed on your system
- Git (for cloning the repository)

### Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/CyberwizD/Task-Management-API-Endpoint.git
   cd Task-Management-API-Endpoint
   ```

2. **Initialize Go modules and install dependencies**
   ```bash
   go mod init Task-Management-API-Endpoint
   go mod tidy
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

4. **Verify the server is running**
   The server will start on port 8080. You should see output similar to:
   ```
   [GIN] Listening and serving HTTP on :8080
   ```

5. **Test the health endpoint**
   ```bash
   curl http://localhost:8080/health
   ```

## API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Health Check
```
GET /health
```
Returns server health status.

### 1. Create Task

**Endpoint**: `POST /api/v1/tasks`

**Request Body**:
```json
{
  "title": "Complete project documentation",
  "description": "Write comprehensive README and API docs"
}
```

**Validation Rules**:
- `title` (required): Must be a non-empty string
- `description` (optional): Must be a string if provided

**Success Response** (201 Created):
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Complete project documentation",
  "description": "Write comprehensive README and API docs",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

**Error Response** (400 Bad Request):
```json
{
  "error": "Validation failed",
  "message": "Title must be a non-empty string"
}
```

### 2. Get All Tasks

**Endpoint**: `GET /api/v1/tasks`

**Success Response** (200 OK):
```json
{
  "tasks": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "title": "Complete project documentation",
      "description": "Write comprehensive README and API docs",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ],
  "count": 1
}
```

### 3. Get Task by ID

**Endpoint**: `GET /api/v1/tasks/{id}`

**Success Response** (200 OK):
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "Complete project documentation",
  "description": "Write comprehensive README and API docs",
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

**Error Response** (404 Not Found):
```json
{
  "error": "Task not found",
  "message": "Task with ID 550e8400-e29b-41d4-a716-446655440000 does not exist"
}
```

## Example Usage

### Using cURL

1. **Create a new task**:
```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go programming",
    "description": "Complete Go tutorial and build a REST API"
  }'
```

2. **Get all tasks**:
```bash
curl -X GET http://localhost:8080/api/v1/tasks
```

3. **Get a specific task** (replace {id} with actual task ID):
```bash
curl -X GET http://localhost:8080/api/v1/tasks/{id}
```

4. **Create a task with only title**:
```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buy groceries"
  }'
```

### Using Postman/Insomnia

1. **POST /api/v1/tasks**
   - Method: POST
   - URL: `http://localhost:8080/api/v1/tasks`
   - Headers: `Content-Type: application/json`
   - Body (JSON):
     ```json
     {
       "title": "Sample Task",
       "description": "This is a sample task"
     }
     ```

2. **GET /api/v1/tasks**
   - Method: GET
   - URL: `http://localhost:8080/api/v1/tasks`

3. **GET /api/v1/tasks/{id}**
   - Method: GET
   - URL: `http://localhost:8080/api/v1/tasks/{task-id}`

## Design Choices and Implementation Details

### Error Handling
- **Consistent Error Format**: All errors return a structured JSON response with `error` and optional `message` fields
- **Appropriate HTTP Status Codes**: 
  - 200 OK for successful retrieval
  - 201 Created for successful task creation
  - 400 Bad Request for validation errors
  - 404 Not Found for non-existent resources
- **Detailed Error Messages**: Provide clear, actionable error descriptions

### Input Validation
- **Server-side Validation**: All input validation happens on the server
- **Required Field Validation**: Title field is mandatory and cannot be empty
- **Type Validation**: Ensures proper data types for all fields
- **String Trimming**: Automatic whitespace trimming for title and description

### Data Storage
- **In-Memory Storage**: Uses a thread-safe map with mutex locks
- **Concurrent Access**: Read-write mutex ensures data consistency across goroutines
- **UUID Generation**: Server-generated UUIDs ensure unique task identifiers

### Security Considerations

#### CORS Configuration
```go
config := cors.DefaultConfig()
config.AllowAllOrigins = true // DEVELOPMENT ONLY
```

**Important**: The current CORS configuration allows requests from any origin (`*`). This is **ONLY appropriate for development environments**. 

**For Production**: 
- Restrict `AllowOrigins` to specific trusted domains
- Consider implementing additional security headers
- Add authentication and authorization mechanisms

#### Additional Security Measures Implemented
- **Input Sanitization**: All string inputs are trimmed and validated
- **Error Information Disclosure**: Error messages are informative but don't expose internal system details
- **Thread Safety**: Concurrent access protection prevents race conditions

## Architecture and Design

### Clean Architecture Principles

This project follows clean architecture principles with clear separation of concerns:

- **Models**: Domain entities and data transfer objects
- **Repository**: Data access layer with interface abstraction
- **Services**: Business logic layer
- **Handlers**: HTTP request/response handling
- **Middleware**: Cross-cutting concerns (CORS, logging, error handling)
- **Config**: Centralized configuration management

### Dependency Injection

The application uses dependency injection to maintain loose coupling:
- Repository → Service → Handler
- Each layer depends on interfaces, not concrete implementations
- Makes testing and future changes easier

### Error Handling Strategy

- **Domain Errors**: Defined in `models/errors.go`
- **Error Propagation**: Errors bubble up through layers
- **HTTP Error Mapping**: Handlers map domain errors to appropriate HTTP status codes
- **Centralized Error Handling**: Middleware handles panics and unhandled errors

## Testing the API

### Manual Testing Checklist

1. ✅ Create task with valid data
2. ✅ Create task with missing title (should fail)
3. ✅ Create task with empty title (should fail)
4. ✅ Create task with only title (should succeed)
5. ✅ Retrieve all tasks
6. ✅ Retrieve specific task by valid ID
7. ✅ Retrieve task with invalid ID (should return 404)
8. ✅ Test CORS headers with browser requests

### Sample Test Scenarios

```bash
# Test 1: Valid task creation
curl -X POST http://localhost:8080/api/v1/tasks -H "Content-Type: application/json" -d '{"title":"Test Task","description":"Test Description"}'

# Test 2: Invalid task creation (missing title)
curl -X POST http://localhost:8080/api/v1/tasks -H "Content-Type: application/json" -d '{"description":"Test Description"}'

# Test 3: Invalid task creation (empty title)
curl -X POST http://localhost:8080/api/v1/tasks -H "Content-Type: application/json" -d '{"title":"","description":"Test Description"}'

# Test 4: Get all tasks
curl -X GET http://localhost:8080/api/v1/tasks

# Test 5: Get non-existent task
curl -X GET http://localhost:8080/api/v1/tasks/invalid-id
```

## Future Enhancements

For production use, consider implementing:

- **Database Integration**: Replace in-memory storage with PostgreSQL/MySQL
- **Authentication & Authorization**: JWT-based auth system
- **Rate Limiting**: Prevent API abuse
- **Logging**: Structured logging with levels
- **Metrics**: Prometheus metrics for monitoring
- **Task Updates**: PUT/PATCH endpoints for task modification
- **Task Deletion**: DELETE endpoint for task removal
- **Pagination**: For large task lists
- **Filtering & Sorting**: Query parameters for task filtering
- **Input Sanitization**: XSS protection for text fields
- **API Documentation**: Swagger/OpenAPI integration

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is created for demonstration purposes as part of a technical assessment.
