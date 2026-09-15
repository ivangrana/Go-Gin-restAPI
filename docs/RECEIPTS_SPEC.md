# Receipts CRUD — Implementation Spec

> **Status:** Draft / Specification
> **Branch:** `feat/receipts-crud-spec`
> **Target Project:** Phinance — `Go-Gin-restAPI`
> **Author:** Generated spec for new contributor implementation

---

## 1. Overview

This document specifies the implementation of a full **Receipts CRUD** resource
for the Phinance personal-finance REST API. The feature follows the project's
existing layered architecture (`models` → `dto` → `services` → `controllers` →
`routes`) and mirrors the patterns already used by `Budget`, `Goal`, and
`Transaction`.

A **Receipt** represents proof-of-purchase data attached to a `Transaction`. It
is the next-level detail record for a single expense — typically produced when
a user records a transaction (or shortly after) and wants to keep merchant,
tax, line-item, and image evidence for later audit / reimbursement.

### 1.1 Goals

- Allow a user to attach one or more Receipts to any of their Transactions.
- Provide standard CRUD operations: list, get-by-id, create, update, delete.
- Keep parity with the existing project conventions (no new dependencies, no
  architectural drift).
- Stay consistent with existing endpoint shapes (`/users/:id/<resource>`) and
  JWT-based auth middleware.

### 1.2 Non-Goals (out of scope)

- File/image storage backend (the spec stores only metadata + a URL string;
  binary uploads are deferred to a follow-up spec).
- OCR / line-item parsing.
- Receipt sharing between users.
- Soft-delete / archive semantics (the existing codebase uses hard deletes).
- Pagination (matches existing endpoints which return full slices).

---

## 2. Domain Model

### 2.1 Relationships

```
User ──┬── Transactions ─── Receipt (1..n)
       │                       │
       └── Receipt (direct ownership, optional)
```

A `Receipt` **must** belong to a `Transaction`. The Transaction must belong to
the same `User` who owns the Receipt. For consistency with how `Budget` and
`Goal` are scoped, the URL will hang Receipt routes under `users/:id/receipts`
so the user-scoping pattern is preserved.

> **Design decision:** Receipts are intentionally not nested under
> `/transactions/:transaction_id/receipts` because the project's existing
> user-centric routing pattern (`/users/:id/<resource>`) is the established
> convention. The `TransactionID` foreign key on the Receipt still enforces the
> relationship at the data layer.

### 2.2 Entity: `Receipt`

| Field            | Type        | Constraints                                | Notes                              |
| ---------------- | ----------- | ------------------------------------------ | ---------------------------------- |
| `ID`             | `uint`      | `gorm:"primary_key"`                       | Auto-incremented.                  |
| `UserID`         | `uint`      | `gorm:"not null"`                          | FK → `users.id`.                   |
| `TransactionID`  | `uint`      | `gorm:"not null"`                          | FK → `transactions.id`.            |
| `Merchant`       | `string`    | `gorm:"size:255; not null"`                | Store / vendor name.               |
| `TotalAmount`    | `float64`   | `gorm:"not null"`                          | Snapshot of receipt total.         |
| `TaxAmount`      | `float64`   | (nullable)                                 | Defaults to `0`.                   |
| `Currency`       | `string`    | `gorm:"size:3; not null; default:'USD'"`   | ISO-4217 code.                     |
| `PurchasedAt`    | `time.Time` | `gorm:"not null"`                          | Date/time on the receipt.          |
| `Notes`          | `string`    | `gorm:"size:1024"`                         | Free-form user notes.              |
| `ImageURL`       | `string`    | `gorm:"size:512"`                          | Optional CDN / storage URL.        |
| `CreatedAt`      | `time.Time` | GORM-managed                               | Auto-set on insert.                |
| `UpdatedAt`      | `time.Time` | GORM-managed                               | Auto-set on insert/update.         |

### 2.3 Inverse relations to add on `User` and `Transactions`

For consistency with how `User` already lists `Budget`, `Transac`, and `Goals`
relationships, the following slice fields should be appended (mirroring existing
style — note the existing `Transac` typo is **not** repeated; use `Receipts`):

```go
// models/User.go  (add)
Receipts []Receipt `gorm:"foreignKey:UserID"`

// models/Transactions.go  (add)
Receipts []Receipt `gorm:"foreignKey:TransactionID"`
```

---

## 3. File-Level Plan

| New / Modified | Path                              | Purpose                                              |
| -------------- | --------------------------------- | ---------------------------------------------------- |
| **New**        | `models/Receipt.go`               | GORM model.                                          |
| **New**        | `dto/receiptDTO.go`               | `ReceiptDTO`, `ReceiptCreateDTO`, `ReceiptUpdateDTO`.|
| **New**        | `services/receipt_service.go`     | DB-layer business logic.                             |
| **New**        | `controllers/receipts.go`         | Gin HTTP handlers.                                   |
| **Modified**   | `routes/routes.go`                | Register `/users/:id/receipts` group.                |
| **Modified**   | `database/connection.go`          | Add `AutoMigrate(&models.Receipt{})`.                |
| **Modified**   | `models/User.go`                  | Add `Receipts []Receipt` relation.                   |
| **Modified**   | `models/Transactions.go`          | Add `Receipts []Receipt` relation.                   |
| **Modified**   | `README.md`                       | Add receipts rows to the endpoints table.            |
| **New**        | `docs/RECEIPTS_SPEC.md`           | This document.                                       |

No new Go module dependencies are required.

---

## 4. API Design

All routes are nested under an authenticated user group, matching the existing
convention in `routes/routes.go`.

### 4.1 Routes

| Method | Endpoint                                | Handler                | Auth | Description                |
| ------ | --------------------------------------- | ---------------------- | ---- | -------------------------- |
| GET    | `/users/:id/receipts`                   | `GetAllReceipts`       | ✅   | List the user's receipts.  |
| GET    | `/users/:id/receipts/:receipt_id`       | `GetReceiptById`       | ✅   | Get a single receipt.      |
| POST   | `/users/:id/receipts`                   | `CreateReceipt`        | ✅   | Create a new receipt.      |
| PUT    | `/users/:id/receipts/:receipt_id`       | `UpdateReceipt`        | ✅   | Update an existing receipt.|
| DELETE | `/users/:id/receipts/:receipt_id`       | `DeleteReceipt`        | ✅   | Delete a receipt.          |

The existing `middleware.AuthMiddleware()` must wrap the receipts sub-group.

### 4.2 Request / Response payloads

#### `GET /users/:id/receipts` — `200 OK`

```json
[
  {
    "id": 12,
    "user_id": 1,
    "transaction_id": 42,
    "merchant": "Whole Foods Market",
    "total_amount": 38.47,
    "tax_amount": 2.91,
    "currency": "USD",
    "purchased_at": "2025-03-14T18:25:43Z",
    "notes": "Weekly groceries",
    "image_url": "https://cdn.example.com/r/abc123.jpg",
    "created_at": "2025-03-14T18:31:02Z",
    "updated_at": "2025-03-14T18:31:02Z"
  }
]
```

#### `POST /users/:id/receipts` — body

```json
{
  "transaction_id": 42,
  "merchant": "Whole Foods Market",
  "total_amount": 38.47,
  "tax_amount": 2.91,
  "currency": "USD",
  "purchased_at": "2025-03-14T18:25:43Z",
  "notes": "Weekly groceries",
  "image_url": "https://cdn.example.com/r/abc123.jpg"
}
```

Response: `200 OK` with `{"message": "receipt created"}`.

#### `PUT /users/:id/receipts/:receipt_id` — body

All fields optional except those the caller wishes to modify. Mirrors the
`BudgetUpdateDTO` / `GoalUpdateDTO` style:

```json
{
  "merchant": "Whole Foods Market #218",
  "total_amount": 39.10,
  "tax_amount": 2.99,
  "notes": "Corrected total",
  "image_url": "https://cdn.example.com/r/abc123-v2.jpg"
}
```

Response: `200 OK` with `{"message": "receipt updated"}`.

#### `DELETE /users/:id/receipts/:receipt_id`

Response: `200 OK` with `{"success": "receipt deleted"}`.

### 4.3 Error responses (consistent with existing handlers)

| Status | Body                                   | Trigger                                |
| ------ | -------------------------------------- | -------------------------------------- |
| `400`  | `{"error": "<gin binding error>"}`     | Malformed JSON.                        |
| `400`  | `{"error": "user id is required"}`     | Empty `:id` param.                     |
| `400`  | `{"error": "receipt_id is required"}`  | Empty `:receipt_id` param.             |
| `400`  | `{"error": "<service error>"}`         | DB / conversion errors.                |
| `404`  | `{"error": "Receipt not found"}`       | `gorm.ErrRecordNotFound` (read paths). |
| `500`  | `{"error": "<service error>"}`         | Delete failures.                       |

---

## 5. Detailed Implementation Outline

### 5.1 `models/Receipt.go`

```go
package models

import "time"

type Receipt struct {
    ID            uint      `gorm:"primary_key"`
    UserID        uint      `gorm:"not null"`
    TransactionID uint      `gorm:"not null"`
    Merchant      string    `gorm:"size:255; not null"`
    TotalAmount   float64   `gorm:"not null"`
    TaxAmount     float64
    Currency      string    `gorm:"size:3; not null; default:'USD'"`
    PurchasedAt   time.Time `gorm:"not null"`
    Notes         string    `gorm:"size:1024"`
    ImageURL      string    `gorm:"size:512"`
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
```

> **Note on naming:** the project is inconsistent (`Budget` is singular,
> `Transactions`/`Categories`/`Goals` are plural). Use the singular form
> `Receipt` for the new struct — it reads better and matches `Budget`.

### 5.2 `dto/receiptDTO.go`

```go
package dto

import "time"

type ReceiptDTO struct {
    ID            uint      `json:"id"`
    UserID        uint      `json:"user_id"`
    TransactionID uint      `json:"transaction_id"`
    Merchant      string    `json:"merchant"`
    TotalAmount   float64   `json:"total_amount"`
    TaxAmount     float64   `json:"tax_amount"`
    Currency      string    `json:"currency"`
    PurchasedAt   time.Time `json:"purchased_at"`
    Notes         string    `json:"notes"`
    ImageURL      string    `json:"image_url"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}

type ReceiptCreateDTO struct {
    TransactionID uint      `json:"transaction_id" binding:"required"`
    Merchant      string    `json:"merchant"        binding:"required"`
    TotalAmount   float64   `json:"total_amount"    binding:"required"`
    TaxAmount     float64   `json:"tax_amount"`
    Currency      string    `json:"currency"`
    PurchasedAt   time.Time `json:"purchased_at"    binding:"required"`
    Notes         string    `json:"notes"`
    ImageURL      string    `json:"image_url"`
}

type ReceiptUpdateDTO struct {
    Merchant    string  `json:"merchant"`
    TotalAmount float64 `json:"total_amount"`
    TaxAmount   float64 `json:"tax_amount"`
    Notes       string  `json:"notes"`
    ImageURL    string  `json:"image_url"`
}
```

> **Style alignment:** `binding:"required"` is **not** currently used anywhere
> in `dto/*.go`, but it is a reasonable, low-risk improvement. If preserving
> exact parity with existing DTOs is preferred, drop the tags and rely on
> service-layer checks instead — either is acceptable.

### 5.3 `services/receipt_service.go`

Mirrors `services/budget_service.go` exactly.

```go
package services

import (
    "Phinance/database"
    "Phinance/dto"
    "Phinance/models"
    "strconv"

    "gorm.io/gorm"
)

func GetAllReceipts(userID string) ([]dto.ReceiptDTO, error) {
    var receipts []models.Receipt
    var receiptDTOs []dto.ReceiptDTO

    resp := database.DB.Find(&receipts, "user_id = ?", userID)
    if resp.Error != nil {
        return nil, resp.Error
    }

    for _, r := range receipts {
        receiptDTOs = append(receiptDTOs, dto.ReceiptDTO{
            ID:            r.ID,
            UserID:        r.UserID,
            TransactionID: r.TransactionID,
            Merchant:      r.Merchant,
            TotalAmount:   r.TotalAmount,
            TaxAmount:     r.TaxAmount,
            Currency:      r.Currency,
            PurchasedAt:   r.PurchasedAt,
            Notes:         r.Notes,
            ImageURL:      r.ImageURL,
            CreatedAt:     r.CreatedAt,
            UpdatedAt:     r.UpdatedAt,
        })
    }

    return receiptDTOs, nil
}

func GetReceiptById(receiptID string) (*dto.ReceiptDTO, error) {
    var r models.Receipt

    resp := database.DB.First(&r, receiptID)
    if resp.Error != nil {
        if resp.Error == gorm.ErrRecordNotFound {
            return nil, resp.Error
        }
        return nil, resp.Error
    }

    return &dto.ReceiptDTO{
        ID:            r.ID,
        UserID:        r.UserID,
        TransactionID: r.TransactionID,
        Merchant:      r.Merchant,
        TotalAmount:   r.TotalAmount,
        TaxAmount:     r.TaxAmount,
        Currency:      r.Currency,
        PurchasedAt:   r.PurchasedAt,
        Notes:         r.Notes,
        ImageURL:      r.ImageURL,
        CreatedAt:     r.CreatedAt,
        UpdatedAt:     r.UpdatedAt,
    }, nil
}

func CreateReceipt(userID string, receiptDTO dto.ReceiptCreateDTO) error {
    id, err := strconv.Atoi(userID)
    if err != nil {
        return err
    }

    receipt := models.Receipt{
        UserID:        uint(id),
        TransactionID: receiptDTO.TransactionID,
        Merchant:      receiptDTO.Merchant,
        TotalAmount:   receiptDTO.TotalAmount,
        TaxAmount:     receiptDTO.TaxAmount,
        Currency:      receiptDTO.Currency,
        PurchasedAt:   receiptDTO.PurchasedAt,
        Notes:         receiptDTO.Notes,
        ImageURL:      receiptDTO.ImageURL,
    }

    if receipt.Currency == "" {
        receipt.Currency = "USD"
    }

    return database.DB.Create(&receipt).Error
}

func UpdateReceipt(receiptID string, receiptDTO dto.ReceiptUpdateDTO) error {
    id, err := strconv.Atoi(receiptID)
    if err != nil {
        return err
    }

    editReceipt := models.Receipt{
        ID:          uint(id),
        Merchant:    receiptDTO.Merchant,
        TotalAmount: receiptDTO.TotalAmount,
        TaxAmount:   receiptDTO.TaxAmount,
        Notes:       receiptDTO.Notes,
        ImageURL:    receiptDTO.ImageURL,
    }

    return database.DB.Updates(&editReceipt).Error
}

func DeleteReceipt(receiptID string) error {
    return database.DB.Delete(&models.Receipt{}, receiptID).Error
}
```

### 5.4 `controllers/receipts.go`

Mirrors `controllers/budgets.go`.

```go
package controllers

import (
    dto "Phinance/dto"
    "Phinance/services"
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetAllReceipts(c *gin.Context) {
    receipts, err := services.GetAllReceipts(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, receipts)
}

func GetReceiptById(c *gin.Context) {
    if c.Param("receipt_id") == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "receipt_id is required"})
        return
    }

    receipt, err := services.GetReceiptById(c.Param("receipt_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, receipt)
}

func CreateReceipt(c *gin.Context) {
    var receiptDTO dto.ReceiptCreateDTO
    if err := c.ShouldBindJSON(&receiptDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if c.Param("id") == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
        return
    }

    if err := services.CreateReceipt(c.Param("id"), receiptDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "receipt created"})
}

func UpdateReceipt(c *gin.Context) {
    var receiptDTO dto.ReceiptUpdateDTO
    if err := c.ShouldBindJSON(&receiptDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if c.Param("receipt_id") == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "receipt_id is required"})
        return
    }

    if err := services.UpdateReceipt(c.Param("receipt_id"), receiptDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "receipt updated"})
}

func DeleteReceipt(c *gin.Context) {
    if err := services.DeleteReceipt(c.Param("receipt_id")); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"success": "receipt deleted"})
}
```

### 5.5 `routes/routes.go` — additions

Inside the existing `user := userGroup.Group("/:id") { ... }` block, add
(after the `transactions` group, before its closing brace):

```go
receipts := user.Group("/receipts")
receipts.Use(middleware.AuthMiddleware())
{
    receipts.GET("/",        controllers.GetAllReceipts)
    receipts.GET("/:receipt_id", controllers.GetReceiptById)
    receipts.POST("/",       controllers.CreateReceipt)
    receipts.PUT("/:receipt_id",  controllers.UpdateReceipt)
    receipts.DELETE("/:receipt_id", controllers.DeleteReceipt)
}
```

### 5.6 `database/connection.go` — additions

Add the line below alongside the other `AutoMigrate` calls inside `Init()`:

```go
db.AutoMigrate(&models.Receipt{})
```

### 5.7 `models/User.go` — additions

Add a relation slice alongside the existing `Budget`, `Transac`, `Goals`:

```go
Receipts []Receipt `gorm:"foreignKey:UserID"`
```

### 5.8 `models/Transactions.go` — additions

```go
Receipts []Receipt `gorm:"foreignKey:TransactionID"`
```

### 5.9 `README.md` — endpoints table

Append a new block to the existing endpoints table:

```markdown
| GET    | /users/:id/receipts              | Get all receipts                |
| POST   | /users/:id/receipts              | Create a new receipt            |
| GET    | /users/:id/receipts/:receipt_id  | Get receipt by ID               |
| PUT    | /users/:id/receipts/:receipt_id  | Update receipt                  |
| DELETE | /users/:id/receipts/:receipt_id  | Delete receipt                  |
```

---

## 6. Acceptance Criteria

The implementation is considered complete when **all** of the following hold:

1. `models/Receipt.go` compiles and the `Receipt` struct matches §2.2.
2. `database/connection.go` registers `models.Receipt` for auto-migration.
3. `models/User.go` and `models/Transactions.go` declare the inverse relation
   slices described in §5.7 / §5.8.
4. `services/receipt_service.go` exposes the five public functions listed in
   §5.3, each with a Go signature matching the existing budget/goal services.
5. `controllers/receipts.go` exposes the five handlers listed in §5.4, each
   returning the status codes listed in §4.3.
6. `routes/routes.go` registers the receipts sub-group under `user/:id`, with
   `middleware.AuthMiddleware()` applied.
7. `go build ./...` succeeds with no new module dependencies.
8. Manual smoke test (with `Authorization: Bearer <jwt>`):
   - `POST /users/1/receipts` with the body from §4.2 returns
     `{"message":"receipt created"}`.
   - `GET  /users/1/receipts` returns a JSON array containing the created
     receipt.
   - `GET  /users/1/receipts/:receipt_id` returns the single receipt DTO.
   - `PUT  /users/1/receipts/:receipt_id` updates the receipt.
   - `DELETE /users/1/receipts/:receipt_id` returns `{"success":"receipt deleted"}`.
9. The README endpoints table is updated.

---

## 7. Open Questions / Follow-ups

- **Image storage:** `ImageURL` is metadata-only. A follow-up spec should
  cover upload (multipart) and a storage backend (S3-compatible).
- **Tax breakdown:** currently a single `TaxAmount`. If jurisdictions require
  per-line tax, add a follow-up model `ReceiptLineItem`.
- **Receipt ↔ Transaction validation:** the current spec does **not** verify
  that `TransactionID` belongs to the same `UserID` before insert. Consider
  adding a service-layer guard in a follow-up.
- **Authorization at the row level:** the middleware only verifies the JWT
  itself; it does not check that the `:id` URL param matches the JWT subject.
  This is a pre-existing project-wide gap and out of scope here, but worth
  tracking.
- **Tests:** the project has no test suite today. Adding one is out of scope
  for this spec but would be a high-value follow-up.

---

## 8. Changelog

| Date       | Author         | Change                                       |
| ---------- | -------------- | -------------------------------------------- |
| (init)     | Spec generator | Initial draft on branch `feat/receipts-crud-spec`. |
