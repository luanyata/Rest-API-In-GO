# REST API in Go

A simple RESTful API built with Go, featuring user management endpoints and error handling.

## Features

- CRUD operations for user management
- CRUD operations for company management
- Custom error handling
- JSON response formatting
- Request middleware
- Hot reload support with Air

## Prerequisites

- Go 1.22.2 or higher
- [Air](https://github.com/cosmtrek/air) for hot reloading (optional)

## Installation

1. Clone the repository
2. Install dependencies:
```sh
go mod download
```

## Running the Application

### Standard Mode
```sh
go run main.go
```

### Development Mode (with hot reload)
```sh
air
```

The server will start at `http://localhost:4545`

## API Endpoints

### Users

| Method | Endpoint    | Description         |
| ------ | ----------- | ------------------- |
| POST   | /users      | Create a new user   |
| GET    | /users      | List all users      |
| GET    | /users/{id} | Get a specific user |
| PUT    | /users/{id} | Update a user       |
| DELETE | /users/{id} | Delete a user       |

### Companies

| Method | Endpoint        | Description            |
| ------ | --------------- | ---------------------- |
| POST   | /companies      | Create a new company   |
| GET    | /companies      | List all companies     |
| GET    | /companies/{id} | Get a specific company |
| PUT    | /companies/{id} | Update a company       |
| DELETE | /companies/{id} | Delete a company       |

### Health Check
- `GET /ping` - Returns "pong!!!" to verify the server is running

## Request/Response Examples

### Create User
```http
POST /users
Content-Type: application/json

{
    "name": "John Doe",
    "email": "john@example.com"
}
```

### Create Company
```http
POST /companies
Content-Type: application/json

{
    "name": "Acme Inc",
    "cnpj": "12345678000199"
}
```

### Error Response Format
```json
{
    "timestamp": "2024-03-21T10:00:00Z",
    "message": "User not found",
    "type": "user_not_found",
    "statusCode": 404,
    "path": "/users/123"
}
```

## Project Structure

```
.
├── internal/
│   ├── company/       # Company handlers
│   ├── errs/          # Error handling
│   ├── httpx/         # HTTP utilities
│   ├── middleware/    # HTTP middleware
│   └── user/          # User handlers
├── main.go            # Application entry point
├── go.mod            
└── go.sum
```

## Dependencies

- [Chi Router](https://github.com/go-chi/chi) - Lightweight HTTP router