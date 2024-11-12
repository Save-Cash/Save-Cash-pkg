package models

import (
	"github.com/google/uuid"
    "time"
)

type User struct {
    ID    uuid.UUID `json:"id,omitempty" db:"id"`
    Name  string    `json:"name" db:"name"`
    Email string    `json:"email" db:"email"`
}

type TransactionCategory struct {
    ID     uuid.UUID `json:"id,omitempty" db:"id"`
    Name   string    `json:"name" db:"name"`
    Type   string    `json:"type" db:"type"` 
}

type FinancialTransaction struct {
    ID          uuid.UUID `json:"id,omitempty" db:"id"`
    UserID      uuid.UUID `json:"user_id" db:"user_id"`  
    CategoryID  uuid.UUID `json:"category_id" db:"category_id"`
    Amount      float64   `json:"amount" db:"amount"`    
    Description string    `json:"description" db:"description"`
    Date        time.Time `json:"date" db:"date"`     
}
