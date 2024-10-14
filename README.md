# Go Authentication System 🚀

This project contains a **robust authentication system** built with Go. It supports user registration, login, and token-based authentication. The project uses **Swaggo** for API documentation and **Air** for hot-reloading during development to ensure seamless coding experience.

## Features 🛠️
- **User Registration & Login**: Secure authentication with password hashing.
- **JWT Authentication**: Token-based access control for protected routes.
- **Swagger Documentation**: Interactive API documentation using Swaggo.
- **Air Hot Reloading**: Automatic reloading during code changes for faster development.

---

## Getting Started 💻

### Prerequisites
- [Go](https://go.dev/dl/) (version >= 1.18)
- [Air](https://github.com/cosmtrek/air) for hot reloading
- Swag CLI for generating Swagger docs: 
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest

#### start with project

### Inline Code Example
Run the following command:
`go mod tidy`

```bash
# Clone the repository
git clone 
# Navigate into the project directory


# Install dependencies
go mod tidy

# Generate Swagger docs
swag init

# Start the server with Air
air
