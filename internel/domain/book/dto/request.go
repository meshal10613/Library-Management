package dto

type CreateBookRequest struct {
	Title           string `json:"title" gorm:"unique" validate:"required,min=2,max=100"`
	Author          string `json:"author" validate:"required"`
	TotalCopies     int    `json:"total_copies" validate:"required"`
}
