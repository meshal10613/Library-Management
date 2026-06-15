package book

import (
	"errors"
	"library-management/pkg/querybuilder"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrorAlreadyExist = errors.New("Book with this title already exists")
var ErrBookNotFound = errors.New("Book not found")

type Repository interface {
	CreateBook(book *Book) error
	GetBookByID(id uuid.UUID) (*Book, error)
	GetAllBooks(ctx fiber.Ctx, opts querybuilder.QueryParams) ([]*Book, int64, error)
	UpdateBook(book *Book) error
	DeleteBook(id uuid.UUID) error
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

func (r *repository) GetBookByID(id uuid.UUID) (*Book, error) {
	var book Book

	result := r.db.First(&book, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrBookNotFound
		}
		return nil, result.Error
	}

	return &book, nil
}

func (r *repository) GetAllBooks(ctx fiber.Ctx, opts querybuilder.QueryParams) ([]*Book, int64, error) {
	var books []*Book
	var total int64

	qb := querybuilder.New(r.db.Model(&Book{}))

	qb.Search(opts.Search, "title", "author")
	qb.Sort(opts.SortBy, opts.Order)
	qb.Filter(ctx)

	qb.DB.Count(&total)

	qb.Paginate(opts.Page, opts.Limit)

	err := qb.DB.Find(&books).Error

	return books, total, err
}

func (r *repository) UpdateBook(book *Book) error {
	result := r.db.Save(book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}
		return result.Error
	}

	return nil
}

func (r *repository) DeleteBook(id uuid.UUID) error {
	var book Book

	if err := r.db.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrBookNotFound
		}
		return err
	}

	return r.db.Delete(&book).Error
}
