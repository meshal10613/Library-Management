package book

import (
	"library-management/internel/domain/book/dto"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title           string    `json:"title" gorm:"unique" validate:"required,min=2,max=100"`
	Author          string    `json:"author" validate:"required"`
	TotalCopies     int       `json:"total_copies" validate:"required"`
	AvailableCopies int       `json:"available_copies"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

func (e *Book) ToResponse() *dto.BookResponse {
	return &dto.BookResponse{
		ID:              e.ID,
		Title:           e.Title,
		Author:          e.Author,
		TotalCopies:     e.TotalCopies,
		AvailableCopies: e.AvailableCopies,
		CreatedAt:       e.CreatedAt.String(),
		UpdatedAt:       e.UpdatedAt.String(),
	}
}
