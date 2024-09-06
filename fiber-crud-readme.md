# Simple CRUD Application with Fiber and GORM

This project is a simple CRUD (Create, Read, Update, Delete) application built with Go, using the Fiber web framework and GORM ORM library with MySQL database.

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (version 1.16 or later)
- MySQL
- Air (optional, for hot reloading during development)

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Set up your MySQL database:
   ```sql
   CREATE DATABASE simple_fiber;
   ```

4. Update the database connection string in `database/connection.go` if necessary.

5. Run the migrations (if you have any):
   ```
   go run migrations/migrate.go
   ```

## Running the Application

### For Development (with hot reloading):

1. Install Air if you haven't already:
   ```
   go install github.com/cosmtrek/air@latest
   ```

2. Run the application:
   ```
   air
   ```

### For Production:

1. Build the application:
   ```
   go build -o main .
   ```

2. Run the built binary:
   ```
   ./main
   ```

The application will start and listen on `http://localhost:3000` by default.

## API Endpoints

- **Create a new post**
  - POST `/api/v1/post`
  - Body: `{ "title": "Your Title", "body": "Your Content" }`

- **Get all posts**
  - GET `/api/v1/posts`

- **Get a specific post**
  - GET `/api/v1/post/{id}`

- **Update a post**
  - PUT `/api/v1/post/{id}`
  - Body: `{ "title": "Updated Title", "body": "Updated Content" }`

- **Delete a post**
  - DELETE `/api/v1/post/{id}`

## Project Structure

```
.
├── main.go
├── go.mod
├── go.sum
├── .air.toml
├── README.md
├── database
│   └── connection.go
├── models
│   └── post.go
├── controllers
│   └── post_controller.go
└── routes
    └── routes.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
