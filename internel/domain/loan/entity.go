package loan

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	UserID     uuid.UUID
	BookID     uuid.UUID
	BorrowDate time.Time
	ReturnDate *time.Time
	Status     string
}
