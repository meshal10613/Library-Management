package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateLoanRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	BookID uuid.UUID `json:"book_id" validate:"required"`
}

type ReturnLoanRequest struct {
	ReturnedAt *time.Time `json:"return_date"`
}
