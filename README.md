# Go Auth

A Sample Project using Gin Framework that implements an auth-based system using JWT Token with OTP authentication and user management features.

## Features

- OTP-based registration and login
- JWT token authentication
- Rate limiting (3 OTP requests per phone number within 10 minutes)
- User management with pagination and search
- MongoDB database integration
- Swagger API documentation

## API Endpoints

### Authentication
- `POST /auth/signup` - Register user with phone number (generates OTP)
- `POST /auth/signup/confirm-otp` - Confirm OTP and get JWT token

### User Management (Requires Authentication)
- `GET /user/` - Get current user details
- `GET /user/{userId}` - Get user by ID
- `GET /user/all` - Get all users with pagination and phone search
  - Query parameters: `page`, `pageSize`, `phone`

## Commands

### Run the application
```sh
go run main.go
```

### Build the application
```sh
go build
```

### Swagger generate
```sh
swag init
```

## API Documentation

Once the application is running, access the Swagger documentation at:
```
http://localhost:3000/swagger/index.html
```

For detailed API examples and usage, see [API_EXAMPLES.md](./API_EXAMPLES.md)
