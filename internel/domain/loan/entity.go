package loan

import (
	"library-management/internel/domain/auth"
	"library-management/internel/domain/book"
	"time"

	"github.com/google/uuid"
)

type LoanStatus string

const (
	StatusBorrowed = "borrowed"
	StatusReturned = "returned"
)

type Loan struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	BookID     uuid.UUID  `gorm:"type:uuid" json:"book_id"`
	BorrowedAt time.Time  `json:"borrow_date"`
	ReturnedAt *time.Time `json:"return_date,omitempty"`
	Status     LoanStatus `json:"status" gorm:"type:varchar(20);default:borrowed"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`

	User auth.User `gorm:"constraint:OnDelete:CASCADE;"`
	Book book.Book `gorm:"constraint:OnDelete:CASCADE;"`
}
