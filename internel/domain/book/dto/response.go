package dto

import (
	"library-management/pkg/httpresponse"

	"github.com/google/uuid"
)

type BookResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	TotalCopies     int       `json:"total_copies"`
	AvailableCopies int       `json:"available_copies"`
	CreatedAt       string    `json:"created_at,omitempty"`
	UpdatedAt       string    `json:"updated_at,omitempty"`
}

type BookPaginationResponse struct {
	Data []BookResponse
	Meta httpresponse.Meta
}
