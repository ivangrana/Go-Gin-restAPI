# Phinance

A RESTful API built with Go and the Gin framework for personal finance management — track budgets, transactions, categories, financial goals and market products.

## Getting Started

### Prerequisites

- Go 1.16+
- PostgreSQL
- Docker

### Installation

1. Clone the repository
2. Create a .env file in the root directory and add the following variables:

```
DB_USER=postgresql
DB_PASSWORD=password
DB_NAME=postgres
DB_HOST=localhost
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

3. Run the following command to start the database:

```
docker-compose up -d
```

4. Run the following command to install the dependencies:

```
go mod download
```

5. Run the following command to run the application:

```
go run main.go
```

## Architecture

The application is split into well-defined layers: HTTP routing, authentication middleware, controllers, services, DTOs and GORM models. The diagrams below illustrate how an incoming HTTP request flows through the system, how the components are wired together, and how the database entities relate to each other.

### Request Workflow

```mermaid
sequenceDiagram
    participant Client
    participant Gin as Gin Router
    participant MW as Auth Middleware
    participant Ctrl as Controller
    participant Svc as Service
    participant DB as PostgreSQL (GORM)

    Client->>Gin: HTTP Request (JWT in Authorization header)
    Gin->>MW: Route to handler
    alt Authenticated route
        MW->>MW: Parse & validate JWT
        alt Invalid token
            MW-->>Client: 401 Unauthorized
        end
    end
    MW->>Ctrl: Forward request
    Ctrl->>Ctrl: Bind JSON payload into DTO
    Ctrl->>Svc: Call business logic (optional)
    Svc->>DB: Query / Create / Update / Delete
    DB-->>Svc: Result
    Svc-->>Ctrl: Domain object / DTO
    Ctrl-->>Client: JSON Response (200 / 400 / 404)
```

### Components Diagram

```mermaid
flowchart LR
    subgraph Client
        A[HTTP Client]
    end

    subgraph Server[Go-Gin Server]
        R[main.go<br/>gin.Default]
        RT[routes/routes.go<br/>RegisterRoutes]
        MW[middleware/auth.go<br/>AuthMiddleware]
        H[handlers/auth.go<br/>Login / createToken]
        C[controllers/*.go]
        S[services/*.go]
        DTO[dto/*.go]
        M[models/*.go]
    end

    DB[(PostgreSQL<br/>via GORM)]

    A --> R
    R --> RT
    RT --> MW
    MW -- valid JWT --> C
    RT -. unprotected .-> C
    RT --> H
    H --> M
    C --> S
    S --> M
    C --> M
    C --> DTO
    M <--> DB
```

### Database Schema

```mermaid
erDiagram
    User ||--o{ Budget       : "has"
    User ||--o{ Transactions : "has"
    User ||--o{ Goals        : "has"
    Budget ||--o{ Goals      : "has"
    Categories ||--o{ Transactions : "categorizes"

    User {
        uint   ID "PK"
        string Name
        string Password
    }

    Budget {
        uint      ID "PK"
        uint      UserID "FK"
        float64   LimitValue
        datetime  InitialDate
        datetime  FinalDate
        datetime  UpdatedAt
    }

    Goals {
        uint    ID "PK"
        uint    UserID "FK"
        float64 Amount
        uint    BudgetID "FK"
    }

    Categories {
        uint   ID "PK"
        string Name
    }

    Transactions {
        uint     ID "PK"
        uint     UserID "FK"
        uint     CategoryID "FK"
        float64  Value
        string   Description
        datetime Date
    }

    MarketProduct {
        uint    ID "PK"
        string  ProductName
        float32 AveragePrice
        string  Priority
    }
```

### Module / Package Layout

```mermaid
flowchart TB
    main[main.go] --> routes
    main --> database

    routes[routes/] --> controllers
    routes[routes/] --> handlers
    routes[routes/] --> middleware

    controllers[controllers/] --> dto
    controllers[controllers/] --> models
    controllers[controllers/] --> database

    services[services/] --> dto
    services[services/] --> models
    services[services/] --> database

    handlers[handlers/] --> dto
    handlers[handlers/] --> models
    handlers[handlers/] --> database

    middleware[middleware/] --> handlers

    database[database/] --> models
    models[models/]

    subgraph Entities
        User
        Budget
        Goals
        Categories
        Transactions
        MarketProduct
    end
```

## API endpoints

| Method | Endpoint                | Description                  |
| ------ | ----------------------- | ---------------------------- |
| POST   | /auth/login             | Authenticate and receive JWT |
| GET    | /users                  | Get all users                |
| POST   | /users                  | Create a new user            |
| GET    | /users/:id              | Get user by ID               |
| PUT    | /users/:id              | Update user                  |
| DELETE | /users/:id              | Delete user                  |
| GET    | /users/:id/goals        | List goals for a user        |
| POST   | /users/:id/goals        | Create a goal for a user     |
| GET    | /users/:id/goals/:goal_id | Get goal by ID             |
| PUT    | /users/:id/goals/:goal_id | Update goal                |
| DELETE | /users/:id/goals/:goal_id | Delete goal                |
| GET    | /users/:id/budgets      | List budgets for a user      |
| POST   | /users/:id/budgets      | Create a budget for a user   |
| GET    | /users/:id/budgets/:budget_id | Get budget by ID        |
| PUT    | /users/:id/budgets/:budget_id | Update budget           |
| DELETE | /users/:id/budgets/:budget_id | Delete budget           |
| GET    | /users/:id/transactions | List transactions for a user |
| POST   | /users/:id/transactions | Create a transaction         |
| GET    | /users/:id/transactions/:transaction_id | Get transaction by ID |
| DELETE | /users/:id/transactions/:transaction_id | Delete transaction     |
| GET    | /categories             | Get all categories           |
| POST   | /categories             | Create a new category        |
| GET    | /categories/:category_id | Get category by ID           |
| PUT    | /categories/:category_id | Update category              |
| DELETE | /categories/:category_id | Delete category              |
| GET    | /market-products        | Get all market products      |
| POST   | /market-products        | Create a new market product  |
| GET    | /market-products/:product_id | Get market product by ID     |
| PUT    | /market-products/:product_id | Update market product        |
| DELETE | /market-products/:product_id | Delete market product        |
