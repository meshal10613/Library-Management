package user

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorAlreadyExist = errors.New("User with this email already exists")
var ErrorInvalidCredentials = errors.New("Invalid email or password")

type Repository interface {
	Register(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Register(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}

		return result.Error
	}

	return nil
}
