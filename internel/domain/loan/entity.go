package loan

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"  json:"id"`
	UserID     uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	BookID     uuid.UUID  `gorm:"type:uuid" json:"book_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date,omitempty"`
	Status     string
}
