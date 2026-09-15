# Phinance - Architecture Documentation

## Table of Contents

- [1. Project Overview](#1-project-overview)
- [2. Technology Stack](#2-technology-stack)
- [3. Architecture Overview](#3-architecture-overview)
- [4. Directory Structure](#4-directory-structure)
- [5. Component Details](#5-component-details)
  - [5.1 Controllers](#51-controllers)
  - [5.2 Services](#52-services)
  - [5.3 Models](#53-models)
  - [5.4 DTOs (Data Transfer Objects)](#54-dtos-data-transfer-objects)
  - [5.5 Handlers](#55-handlers)
  - [5.6 Middleware](#56-middleware)
  - [5.7 Database](#57-database)
  - [5.8 Routes](#58-routes)
- [6. Database Schema & Entity Relationships](#6-database-schema--entity-relationships)
- [7. API Endpoints](#7-api-endpoints)
- [8. Authentication & Authorization](#8-authentication--authorization)
- [9. Docker Configuration](#9-docker-configuration)
- [10. Environment Variables](#10-environment-variables)
- [11. Request Flow](#11-request-flow)
- [12. Dependencies](#12-dependencies)

---

## 1. Project Overview

**Phinance** is a personal finance REST API built with Go and the Gin web framework. The application provides endpoints for managing users, budgets, goals, transactions, categories, and market products. It follows a layered architecture pattern with clear separation of concerns between controllers, services, and data access layers.

---

## 2. Technology Stack

| Technology | Version | Purpose |
|------------|---------|---------|
| Go | 1.22+ | Programming Language |
| Gin | v1.10.0 | HTTP Web Framework |
| GORM | v1.25.12 | ORM (Object-Relational Mapping) |
| PostgreSQL | 11.2+ | Database |
| JWT (golang-jwt) | v4.5.1 | Authentication |
| godotenv | v1.5.1 | Environment Variable Management |
| Docker | - | Containerization |
| Docker Compose | v3 | Multi-container Orchestration |

---

## 3. Architecture Overview

The project follows a **Layered Architecture** pattern with the following layers:

```
┌─────────────────────────────────────────────────────────────┐
│                      Client (HTTP Request)                   │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        Routes Layer                         │
│  (Defines API endpoints and maps them to handlers)          │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Middleware Layer                        │
│  (JWT Authentication, Request Validation)                   │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                     Controllers Layer                        │
│  (Handles HTTP request/response, input binding)             │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Services Layer                          │
│  (Business logic, data transformation)                      │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                       Models Layer                           │
│  (GORM models, database schema definition)                  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        Database                              │
│  (PostgreSQL - Data persistence)                            │
└─────────────────────────────────────────────────────────────┘
```

### Key Design Principles:

- **Separation of Concerns**: Each layer has a distinct responsibility
- **DTO Pattern**: Data Transfer Objects are used to decouple API contracts from database models
- **Dependency Injection**: Database connection is shared via a package-level variable
- **Middleware-based Auth**: JWT authentication is applied at the route group level

---

## 4. Directory Structure

```
Go-Gin-restAPI/
├── main.go                          # Application entry point
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── Dockerfile                       # Docker build configuration
├── docker-compose.yml               # Docker Compose configuration
├── .env.example                     # Environment variable template
├── LICENSE                          # License file
├── README.md                        # Project README
│
├── controllers/                     # HTTP request handlers
│   ├── budgets.go                   # Budget controller functions
│   ├── categories.go                # Category controller functions
│   ├── goals.go                     # Goal controller functions
│   ├── marketProduct.go             # Market product controller functions
│   ├── transactions.go              # Transaction controller functions
│   └── user.go                      # User controller functions
│
├── database/                        # Database connection
│   └── connection.go                # DB initialization and migration
│
├── dto/                             # Data Transfer Objects
│   ├── budgetDTO.go                 # Budget DTOs
│   ├── categoryDTO.go               # Category DTOs
│   ├── goalDTO.go                   # Goal DTOs
│   ├── productDTO.go                # Market product DTOs
│   ├── transactionDTO.go            # Transaction DTOs
│   └── userDTO.go                   # User DTOs
│
├── handlers/                        # Specialized handlers
│   └── auth.go                      # Authentication handler (Login)
│
├── middleware/                      # HTTP middleware
│   └── auth.go                      # JWT authentication middleware
│
├── models/                          # GORM database models
│   ├── Budget.go                    # Budget model
│   ├── Category.go                  # Category model
│   ├── Goals.go                     # Goals model
│   ├── MarketProduct.go             # Market product model
│   ├── Transactions.go              # Transactions model
│   └── User.go                      # User model
│
├── routes/                          # Route definitions
│   └── routes.go                    # All API route registrations
│
└── services/                        # Business logic layer
    ├── budget_service.go            # Budget business logic
    ├── category_service.go          # Category business logic
    └── goal_service.go              # Goal business logic
```

---

## 5. Component Details

### 5.1 Controllers

Controllers handle HTTP requests, bind input data, and return responses. They act as the entry point for API requests.

**Location:** `controllers/`

| File | Functions | Description |
|------|-----------|-------------|
| `budgets.go` | `GetAllBudgets`, `GetBudgetById`, `CreateBudget`, `UpdateBudget`, `DeleteBudget` | Handles budget-related HTTP requests |
| `categories.go` | `GetAllCategories`, `CreateCategory`, `GetCategoryById`, `UpdateCategory`, `DeleteCategory` | Handles category-related HTTP requests |
| `goals.go` | `GetAllGoals`, `GetGoalById`, `CreateGoal`, `UpdateGoal`, `DeleteGoal` | Handles goal-related HTTP requests |
| `marketProduct.go` | `GetAllProducts`, `CreateMarketProduct`, `GetMarketProductById`, `UpdateMarketProduct`, `DeleteMarketProduct` | Handles market product HTTP requests |
| `transactions.go` | `GetAllTransactions`, `CreateTransaction`, `GetTransactionById`, `DeleteTransaction` | Handles transaction HTTP requests |
| `user.go` | `GetAllUsers`, `CreateUser`, `GetUserByID`, `UpdateUser`, `DeleteUser` | Handles user HTTP requests |

**Note:** Some controllers (user, transactions, marketProduct) directly access the database without using the service layer, while others (budgets, categories, goals) use the service layer for business logic.

---

### 5.2 Services

The service layer contains business logic and acts as an intermediary between controllers and the database.

**Location:** `services/`

| File | Functions | Description |
|------|-----------|-------------|
| `budget_service.go` | `GetAllBudgets`, `GetBudgetById`, `CreateBudget`, `UpdateBudget`, `DeleteBudget` | Budget business logic with DTO conversion |
| `category_service.go` | `GetAllCategories`, `GetCategoryById`, `CreateCategory`, `UpdateCategory`, `DeleteCategory` | Category business logic with DTO conversion |
| `goal_service.go` | `GetAllGoals`, `GetGoalById`, `CreateGoal`, `UpdateGoal`, `DeleteGoal` | Goal business logic with DTO conversion |

**Responsibilities:**
- Data transformation between models and DTOs
- Business rule enforcement
- Database query execution
- Error handling

---

### 5.3 Models

Models define the database schema using GORM tags and represent the data structure.

**Location:** `models/`

| File | Struct | Fields | Description |
|------|--------|--------|-------------|
| `User.go` | `User` | `ID`, `Name`, `Password`, `Budget`, `Transac`, `Goals` | User entity with relationships |
| `Budget.go` | `Budget` | `ID`, `UserID`, `LimitValue`, `InitialDate`, `FinalDate`, `UpdatedAt` | Budget entity linked to User |
| `Goals.go` | `Goals` | `ID`, `UserID`, `Amount`, `BudgetID`, `Budget` | Goal entity linked to User and Budget |
| `Category.go` | `Categories` | `ID`, `Transactions`, `Name` | Category entity with transactions |
| `Transactions.go` | `Transactions` | `ID`, `UserID`, `CategoryID`, `Value`, `Description`, `Date` | Transaction entity linked to User and Category |
| `MarketProduct.go` | `MarketProduct` | `ID`, `ProductName`, `AveragePrice`, `Priority` | Market product entity |

---

### 5.4 DTOs (Data Transfer Objects)

DTOs define the API request/response contracts, decoupling the external API from internal database models.

**Location:** `dto/`

| File | Types | Description |
|------|-------|-------------|
| `userDTO.go` | `UserDTO`, `UserCreateDTO`, `UserUpdateDTO`, `UserLoginDTO` | User data transfer objects |
| `budgetDTO.go` | `BudgetDTO`, `BudgetCreateDTO`, `BudgetUpdateDTO` | Budget data transfer objects |
| `goalDTO.go` | `GoalDTO`, `GoalCreateDTO`, `GoalUpdateDTO` | Goal data transfer objects |
| `categoryDTO.go` | `CategoryDTO`, `CategoryCreateDTO`, `CategoryUpdateDTO` | Category data transfer objects |
| `transactionDTO.go` | `TransactionDTO`, `TransactionCreateDTO` | Transaction data transfer objects |
| `productDTO.go` | `MarketProductDTO`, `MarketProductCreateDTO` | Market product data transfer objects |

**DTO Pattern Benefits:**
- Prevents over-posting attacks
- Allows different views of the same model
- Decouples API contract from database schema
- Enables input validation

---

### 5.5 Handlers

Handlers contain specialized business logic that doesn't fit into the standard controller pattern.

**Location:** `handlers/`

| File | Functions | Description |
|------|-----------|-------------|
| `auth.go` | `Login`, `createToken` | Authentication handler for user login and JWT token generation |

**Note:** The `Login` handler is placed in handlers (not controllers) because it involves authentication logic that differs from standard CRUD operations.

---

### 5.6 Middleware

Middleware functions process requests before they reach the controllers.

**Location:** `middleware/`

| File | Functions | Description |
|------|-----------|-------------|
| `auth.go` | `AuthMiddleware` | JWT authentication middleware |

**AuthMiddleware Flow:**
1. Extracts `Authorization` header
2. Validates `Bearer` token format
3. Parses and validates JWT token
4. Verifies signing method (HS256)
5. Sets claims in context for downstream handlers
6. Aborts request on authentication failure

---

### 5.7 Database

**Location:** `database/`

| File | Functions | Description |
|------|-----------|-------------|
| `connection.go` | `Init` | Database connection initialization and auto-migration |

**Database Configuration:**
- **Driver:** PostgreSQL via `gorm.io/driver/postgres`
- **Connection:** Constructed from environment variables
- **Auto-Migration:** All models are automatically migrated on startup
- **SSL Mode:** Disabled (`sslmode=disable`)

**Migrated Models:**
- `User`
- `Budget`
- `Goals`
- `Categories`
- `Transactions`
- `MarketProduct`

---

### 5.8 Routes

**Location:** `routes/`

| File | Functions | Description |
|------|-----------|-------------|
| `routes.go` | `RegisterRoutes` | Registers all API routes |

**Route Groups:**
- `/auth` - Authentication endpoints (no middleware)
- `/users` - User endpoints (with auth middleware)
- `/categories` - Category endpoints (no middleware)
- `/market-products` - Market product endpoints (no middleware)

---

## 6. Database Schema & Entity Relationships

### Entity Relationship Diagram

```
┌─────────────────┐       ┌─────────────────┐
│      User       │       │    Category     │
├─────────────────┤       ├─────────────────┤
│ ID (PK)         │       │ ID (PK)         │
│ Name            │       │ Name            │
│ Password        │       └────────┬────────┘
└────────┬────────┘                │
         │                         │
         │ 1:N                     │ 1:N
         ▼                         ▼
┌─────────────────┐       ┌─────────────────┐
│     Budget      │       │  Transactions   │
├─────────────────┤       ├─────────────────┤
│ ID (PK)         │       │ ID (PK)         │
│ UserID (FK)     │       │ UserID (FK)     │
│ LimitValue      │       │ CategoryID (FK) │
│ InitialDate     │       │ Value           │
│ FinalDate       │       │ Description     │
│ UpdatedAt       │       │ Date            │
└────────┬────────┘       └─────────────────┘
         │
         │ 1:N
         ▼
┌─────────────────┐
│      Goals      │
├─────────────────┤
│ ID (PK)         │
│ UserID (FK)     │
│ Amount          │
│ BudgetID (FK)   │
└─────────────────┘

┌─────────────────┐
│  MarketProduct  │
├─────────────────┤
│ ID (PK)         │
│ ProductName     │
│ AveragePrice    │
│ Priority        │
└─────────────────┘
```

### Relationships

| Parent | Child | Foreign Key | Type |
|--------|-------|-------------|------|
| User | Budget | `UserID` | One-to-Many |
| User | Transactions | `UserID` | One-to-Many |
| User | Goals | `UserID` | One-to-Many |
| Budget | Goals | `BudgetID` | One-to-Many |
| Category | Transactions | `CategoryID` | One-to-Many |

---

## 7. API Endpoints

### Authentication

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| POST | `/auth/login` | No | User login, returns JWT token |

### Users

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | `/users` | Yes | Get all users |
| POST | `/users` | Yes | Create a new user |
| GET | `/users/:id` | Yes | Get user by ID |
| PUT | `/users/:id` | Yes | Update user |
| DELETE | `/users/:id` | Yes | Delete user |

### Goals (Nested under Users)

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | `/users/:id/goals` | Yes | Get all goals for a user |
| POST | `/users/:id/goals` | Yes | Create a new goal |
| GET | `/users/:id/goals/:goal_id` | Yes | Get goal by ID |
| PUT | `/users/:id/goals/:goal_id` | Yes | Update goal |
| DELETE | `/users/:id/goals/:goal_id` | Yes | Delete goal |

### Budgets (Nested under Users)

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | `/users/:id/budgets` | Yes | Get all budgets for a user |
| POST | `/users/:id/budgets` | Yes | Create a new budget |
| GET | `/users/:id/budgets/:budget_id` | Yes | Get budget by ID |
| PUT | `/users/:id/budgets/:budget_id` | Yes | Update budget |
| DELETE | `/users/:id/budgets/:budget_id` | Yes | Delete budget |

### Transactions (Nested under Users)

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | `/users/:id/transactions` | Yes | Get all transactions for a user |
| POST | `/users/:id/transactions` | Yes | Create a new transaction |
| GET | `/users/:id/transactions/:transaction_id` | Yes | Get transaction by ID |
| DELETE | `/users/:id/transactions/:transaction_id` | Yes | Delete transaction |

### Categories

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | `/categories` | No | Get all categories |
| POST | `/categories` | No | Create a new category |
| GET | `/categories/:category_id` | No | Get category by ID |
| PUT | `/categories/:category_id` | No | Update category |
| DELETE | `/categories/:category_id` | No | Delete category |

### Market Products

| Method | Endpoint | Auth Required | Description |
|--------|----------|---------------|-------------|
| GET | `/market-products` | No | Get all market products |
| POST | `/market-products` | No | Create a new market product |
| GET | `/market-products/:product_id` | No | Get market product by ID |
| PUT | `/market-products/:product_id` | No | Update market product |
| DELETE | `/market-products/:product_id` | No | Delete market product |

---

## 8. Authentication & Authorization

### JWT Authentication Flow

```
┌─────────┐                    ┌─────────┐                    ┌─────────┐
│  Client │                    │  Server │                    │Database │
└────┬────┘                    └────┬────┘                    └────┬────┘
     │                              │                              │
     │  POST /auth/login            │                              │
     │  {name, password}            │                              │
     │─────────────────────────────>│                              │
     │                              │                              │
     │                              │  Query user by name          │
     │                              │─────────────────────────────>│
     │                              │                              │
     │                              │  Return user record          │
     │                              │<─────────────────────────────│
     │                              │                              │
     │                              │  Validate password           │
     │                              │  Generate JWT token          │
     │                              │                              │
     │  {token: "eyJhbG..."}       │                              │
     │<─────────────────────────────│                              │
     │                              │                              │
     │  GET /users                  │                              │
     │  Authorization: Bearer <token>                              │
     │─────────────────────────────>│                              │
     │                              │                              │
     │                              │  Validate token              │
     │                              │  Extract claims              │
     │                              │  Process request             │
     │                              │                              │
     │  {users: [...]}              │                              │
     │<─────────────────────────────│                              │
     │                              │                              │
```

### JWT Token Details

| Property | Value |
|----------|-------|
| Signing Method | HS256 |
| Secret Source | `JWT_SECRET` environment variable |
| Token Expiration | 3600 seconds (1 hour) |
| Claims | `name` (username), `exp` (expiration) |
| Header Format | `Authorization: Bearer <token>` |

---

## 9. Docker Configuration

### Dockerfile (Multi-stage Build)

```dockerfile
# Build Stage
FROM golang:1.12-alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Runtime Stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/app .
EXPOSE 8080
CMD ["./app"]
```

### Docker Compose

| Service | Image | Ports | Environment |
|---------|-------|-------|-------------|
| `api` | Custom (from Dockerfile) | 8080:8080 | `PORT`, `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` |
| `db` | postgres:11.2-alpine | 5432:5432 | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` |

**Volume Mounts:**
- API: `./` → `/go/src/app` (development hot-reload)
- DB: `./db/init.sql` → `/docker-entrypoint-initdb.d/init.sql` (initialization script)

---

## 10. Environment Variables

| Variable | Description | Example Value |
|----------|-------------|---------------|
| `DB_USER` | PostgreSQL username | `root` |
| `DB_PASSWORD` | PostgreSQL password | `root` |
| `DB_NAME` | Database name | `phinance` |
| `DB_HOST` | Database host | `localhost` or `db` (Docker) |
| `DB_PORT` | Database port | `5432` |
| `JWT_SECRET` | Secret key for JWT signing | `your-super-secret-jwt-key` |
| `PORT` | Server port (optional) | `8080` |

---

## 11. Request Flow

### Authenticated Request Flow

```
Client Request
      │
      ▼
┌─────────────────┐
│  Gin Router     │
│  (routes.go)    │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Auth Middleware│
│  (middleware/   │
│   auth.go)      │
│                 │
│  • Extract token│
│  • Validate JWT │
│  • Set claims   │
└────────┬────────┘
         │ (if valid)
         ▼
┌─────────────────┐
│   Controller    │
│ (controllers/)  │
│                 │
│  • Bind JSON    │
│  • Validate ID  │
│  • Call service │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Service      │
│  (services/)    │
│                 │
│  • Business logic│
│  • DTO conversion│
│  • DB operations │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Model       │
│   (models/)     │
│                 │
│  • GORM queries │
│  • Relationships│
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   PostgreSQL    │
│   Database      │
└─────────────────┘
```

---

## 12. Dependencies

### Direct Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| `github.com/gin-gonic/gin` | v1.10.0 | HTTP web framework |
| `github.com/golang-jwt/jwt/v4` | v4.5.1 | JWT authentication |
| `gorm.io/driver/postgres` | v1.5.11 | PostgreSQL driver for GORM |
| `gorm.io/gorm` | v1.25.12 | ORM library |
| `github.com/golang-jwt/jwt` | v3.2.2+incompatible | JWT (legacy, used in handlers) |
| `github.com/joho/godotenv` | v1.5.1 | Environment file loading |

### Key Indirect Dependencies

| Package | Purpose |
|---------|---------|
| `github.com/jackc/pgx/v5` | PostgreSQL driver |
| `github.com/go-playground/validator/v10` | Request validation |
| `github.com/bytedance/sonic` | JSON serialization |
| `github.com/goccy/go-json` | JSON processing |
| `golang.org/x/crypto` | Cryptographic functions |

---

## Appendix: Module Information

- **Module Name:** `Phinance`
- **Go Version:** 1.22
- **Toolchain:** go1.23.5
- **Server Port:** 8080 (default)
