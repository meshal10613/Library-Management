package book

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorAlreadyExist = errors.New("Book with this title already exists")

type Repository interface {
	CreateBook(book *Book) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateBook(book *Book) error {
	book.AvailableCopies = book.TotalCopies

	result := r.db.Create(book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}

		return result.Error
	}

	return nil
}
