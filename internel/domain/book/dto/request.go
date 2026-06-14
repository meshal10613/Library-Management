package dto

import "errors"

type CreateBookRequest struct {
	Title       string `json:"title" gorm:"unique" validate:"required,min=2,max=100"`
	Author      string `json:"author" validate:"required"`
	TotalCopies int    `json:"total_copies" validate:"required"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title" validate:"omitempty,min=2,max=100"`
	Author      *string `json:"author" validate:"omitempty"`
	TotalCopies *int    `json:"total_copies" validate:"omitempty,gt=0"`
}

func (r *UpdateBookRequest) Validate() error {
	if r.Title == nil && r.Author == nil && r.TotalCopies == nil {
		return errors.New("at least one field must be provided")
	}
	return nil
}
