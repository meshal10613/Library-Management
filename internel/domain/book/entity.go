package book

type Book struct {
	ID              uint `gorm:"primaryKey"`
	Title           string
	Author          string
	TotalCopies     int
	AvailableCopies int
}
