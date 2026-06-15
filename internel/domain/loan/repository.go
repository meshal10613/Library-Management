package loan

import (
	"errors"
	// "library-management/pkg/querybuilder"

	// "github.com/gofiber/fiber/v3"
	// "github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrLoanNotFound = errors.New("Loan not found")
var ErrorAlreadyExist = errors.New("Loan already exist")

type Repository interface {
	CreateLoan(*Loan) error
	// GetLoanByID(uuid.UUID) (*Loan, error)
	// GetAllLoans(fiber.Ctx, querybuilder.QueryParams) ([]*Loan, int64, error)
	// UpdateLoan(*Loan) error
	// DeleteLoan(uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateLoan(loan *Loan) error {
	result := r.db.Create(loan)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}

		return result.Error
	}

	return nil
}
