# Aws-SES-Backend-task

Directory Structure 
/mock-ses
├── main.go
├── handlers
│   ├── email.go
│   └── stats.go
├── models
│   ├── email.go
│   └── stats.go
├── middleware
│   └── rate_limit.go
├── utils
│   └── validation.go
└── README.md

# Mock SES API

This is a mock implementation of the AWS Simple Email Service (SES) API using Go and Gin.

## Endpoints

- `POST /send-email`: Simulates sending an email.
- `GET /statistics`: Returns statistics of API usage.

## Setup

1. Clone the repository.
2. Run `go mod init mock-ses` to initialize the Go module.
3. Install Gin: `go get -u github.com/gin-gonic/gin`.
4. Run the application: `go run main.go`.

## Usage

Use a tool like Postman or curl to test the API endpoints.
