package user

import (
	"library-management/internel/domain/auth"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers() ([]auth.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllUsers() ([]auth.User, error) {
	var users []auth.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
