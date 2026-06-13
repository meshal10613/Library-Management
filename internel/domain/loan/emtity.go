package loan

import "time"

type Loan struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	BookID     uint
	BorrowDate time.Time
	ReturnDate *time.Time
	Status     string // "ACTIVE" or "RETURNED"
}
