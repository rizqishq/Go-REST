# 🚀 Go REST User API

A RESTful API project built with **Go**, created as part of my learning journey into **Layered Architecture**.  
It implements user management features using a modular structure (Controller → Service → Repository), in-memory storage, and includes Swagger documentation for API exploration.


---

## ✨ Features

- ✅ Full **User CRUD** operations (Create, Read, Update, Delete)
- 🧠 In-memory data repository (no external database required)
- 🔐 **Password hashing** using SHA-256 (for demonstration purposes)
- 🧩 Middleware for **logging** and **panic recovery**
- ❤️ `/health` endpoint for monitoring server status
- 📚 Interactive API documentation with **Swagger UI**
- 🧪 Simple, extensible structure for adding tests and new features

---

## 📌 API Overview

> **Base Path:** `/api/v1`

### 🔄 Health Check
- `GET /api/v1/health` → Returns API status and uptime

### 👤 User Endpoints
- `GET /users` → List all users  
- `POST /users` → Create a new user  
- `GET /users/{id}` → Get user by ID  
- `PUT /users/{id}` → Update user  
- `DELETE /users/{id}` → Delete user  

---

## 📦 Request / Response Schemas

| Type                 | Fields                                                                 |
|----------------------|------------------------------------------------------------------------|
| `CreateUserRequest`  | `username`, `email`, `password`, `first_name`, `last_name`             |
| `UpdateUserRequest`  | `username`, `email`, `first_name`, `last_name`                         |
| `UserResponse`       | `id`, `username`, `email`, `first_name`, `last_name`, `created_at`, `updated_at` |

### 📝 Example: Create User

**Request**
```json
POST /api/v1/users
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "secret123",
  "first_name": "John",
  "last_name": "Doe"
}
```

**Response**
```json
{
  "id": "1",
  "username": "johndoe",
  "email": "john@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "created_at": "2025-07-11T10:00:00Z",
  "updated_at": "2025-07-11T10:00:00Z"
}
```

---

## ⚙️ Getting Started

### 📋 Prerequisites
- Go **v1.18+**

### 🛠 Installation

```bash
git clone https://github.com/rizqishq/Go-REST.git
cd Go-REST
go mod tidy
```

### 🚀 Run the Server

```bash
go run main.go
```

By default, the server runs at:  
👉 `http://localhost:8080`

---

## 📖 API Documentation

Swagger UI is available at:  
👉 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Use Swagger UI to interactively test endpoints, view request/response schemas, and explore the API.

---

## 🗂 Project Structure

```
.
├── main.go             # Entry point / Server setup
├── config/             # App configuration
├── controllers/        # HTTP handlers (API endpoints)
├── services/           # Business logic
├── repositories/       # In-memory data storage
├── models/             # Data models and request/response structs
├── middleware/         # Logging & recovery middleware
├── utils/              # Utility functions (e.g., password hashing)
└── docs/               # Swagger/OpenAPI docs
```

---

## 🏗️ Architecture

This project follows a **layered architecture** for clarity and maintainability:

- **Controllers:** Handle HTTP requests and responses.
- **Services:** Contain business logic and validation.
- **Repositories:** Manage data storage (in-memory for this project).
- **Middleware:** Add cross-cutting concerns like logging and error recovery.
- **Utils:** Provide helper functions (e.g., password hashing).

---

## ⚙️ Environment Variables

| Variable                  | Default   | Description                   |
|---------------------------|-----------|-------------------------------|
| `SERVER_PORT`             | `8080`    | Port for server               |
| `SERVER_READ_TIMEOUT`     | `15s`     | Max time to read request      |
| `SERVER_WRITE_TIMEOUT`    | `15s`     | Max time to write response    |
| `SERVER_IDLE_TIMEOUT`     | `60s`     | Max keep-alive timeout        |
| `SERVER_SHUTDOWN_TIMEOUT` | `15s`     | Graceful shutdown timeout     |
| `DB_MAX_CONNECTIONS`      | `10`      | Max pseudo-connections in mem |

You can override these by setting environment variables before running the server.

---

## 📝 Notes

- This project uses **in-memory** storage for simplicity and learning.  
- Passwords are hashed using **SHA-256**, which is **not secure for production use** (no salt, no bcrypt).
- The codebase is designed for easy extension—swap out the repository layer for a real database, or add JWT authentication as needed.

---

## 🤝 Contributing

Contributions are welcome!  
Feel free to open issues or submit pull requests to improve features, fix bugs, or enhance documentation.

---

## 📄 License

Released under the [MIT License](LICENSE).