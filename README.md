# Phinance

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

## Request workflow example:

```mermaid
sequenceDiagram
    participant Client
    participant Gin as Gin Router
    participant MW as Auth Middleware
    participant Ctrl as Controller
    participant Svc as Service Layer
    participant GORM as GORM
    participant DB as PostgreSQL

    Client->>Gin: HTTP Request (Authorization: Bearer <JWT>)
    Gin->>MW: Route matching
    alt Protected route
        MW->>MW: Parse & validate JWT
        alt Token invalid
            MW-->>Client: 401 Unauthorized
        else Token valid
            MW->>Ctrl: Forward to controller
        end
    else Public route (e.g. /auth/login)
        Gin->>Ctrl: Forward directly
    end
    Ctrl->>Ctrl: Bind/validate JSON payload
    Ctrl->>Svc: Call business logic function
    Svc->>GORM: Query / Create / Update / Delete
    GORM->>DB: SQL statement
    DB-->>GORM: Result set
    GORM-->>Svc: Model / error
    Svc-->>Ctrl: DTO / error
    Ctrl-->>Client: JSON response (200 / 400 / 404)
```

## Components Diagram:

```mermaid
flowchart LR
    subgraph Client
        U[HTTP Client]
    end

    subgraph Server["Go-Gin REST API :8080"]
        direction TB
        M[main.go<br/>Bootstrap]
        R[routes/routes.go<br/>RegisterRoutes]
        MW[middleware/auth.go<br/>AuthMiddleware JWT]
        H[handlers/auth.go<br/>Login / Token]
        C[controllers/<br/>users, budgets, goals,<br/>categories, transactions,<br/>market products]
        S[services/<br/>budget_service,<br/>goal_service,<br/>category_service]
        DTO[dto/<br/>Request/Response DTOs]
        MOD[models/<br/>User, Budget, Goals,<br/>Categories, Transactions,<br/>MarketProduct]
        DB[database/connection.go<br/>GORM Init + AutoMigrate]
    end

    PG[(PostgreSQL)]

    U -->|HTTP| M
    M --> R
    R --> MW
    R --> H
    MW --> C
    C --> S
    C --> DTO
    S --> MOD
    C --> MOD
    H --> MOD
    H --> DB
    S --> DB
    C --> DB
    DB -->|GORM| PG
```

## Database Diagram:

```mermaid
erDiagram
    USER ||--o{ BUDGET       : "has"
    USER ||--o{ TRANSACTION  : "has"
    USER ||--o{ GOALS        : "has"
    BUDGET ||--o{ GOALS      : "owns"
    CATEGORIES ||--o{ TRANSACTION : "classifies"

    USER {
        uint   id PK
        string name
        string password
    }

    BUDGET {
        uint     id PK
        uint     user_id FK
        float    limit_value
        datetime initial_date
        datetime final_date
        datetime updated_at
    }

    TRANSACTION {
        uint     id PK
        uint     user_id FK
        uint     category_id FK
        float    value
        string   description
        datetime date
    }

    GOALS {
        uint   id PK
        uint   user_id FK
        float  amount
        uint   budget_id FK
    }

    CATEGORIES {
        uint   id PK
        string name
    }

    MARKET_PRODUCT {
        uint    id PK
        string  product_name
        float   average_price
        string  priority
    }
```


## API endpoints

| Method | Endpoint                | Description                  |
| ------ | ----------------------- | ---------------------------- |
| GET    | /users                  | Get all users                |
| POST   | /users                  | Create a new user            |
| GET    | /users/:id              | Get user by ID               |
| PUT    | /users/:id              | Update user                  |
| DELETE | /users/:id              | Delete user                  |
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
| GET    | /budgets                | Get all budgets              |
| POST   | /budgets                | Create a new budget          |
| GET    | /budgets/:budget_id     | Get budget by ID             |
| PUT    | /budgets/:budget_id     | Update budget                |
| DELETE | /budgets/:budget_id     | Delete budget                |
| GET    | /goals                  | Get all goals                |
| POST   | /goals                  | Create a new goal            |
| GET    | /goals/:goal_id         | Get goal by ID               |
| PUT    | /goals/:goal_id         | Update goal                  |
| DELETE | /goals/:goal_id         | Delete goal                  |
| GET    | /transactions           | Get all transactions         |
| POST   | /transactions           | Create a new transaction     |
| GET    | /transactions/:transaction_id | Get transaction by ID       |
| PUT    | /transactions/:transaction_id | Update transaction          |
| DELETE | /transactions/:transaction_id | Delete transaction          |
