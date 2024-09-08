# Healthcare App - Backend

This is the backend for the Healthcare Fullstack App. It provides the APIs for user signup, login, and a protected resource accessible only to authenticated users.

## Features

- **Signup API**: Users can create an account.
- **Login API**: Users can log in with their email and password to receive a JWT token.
- **Protected Resource**: A protected endpoint that requires a valid JWT token for access.

## Technologies Used

- **Go (Golang)**: Backend programming language
- **Gin**: HTTP web framework for building APIs
- **GORM**: ORM for interacting with the database
- **SQLite**: Lightweight database for data storage
- **JWT (JSON Web Tokens)**: For authentication and session management
- Hosted on [Render](https://render.com)

## Project Structure

```plaintext
backend/
├── config/
│   └── config.go            # Database configuration (SQLite)
├── controllers/
│   └── authController.go    # Handles user signup, login, and JWT generation
├── middleware/
│   └── jwtMiddleware.go     # Middleware for JWT authentication
├── models/
│   └── user.go              # User model (GORM)
├── routes/
│   └── auth.go              # Routes for signup, login, and protected resource
├── healthcare_app.db        # SQLite database file
├── main.go                  # Main entry point of the application
