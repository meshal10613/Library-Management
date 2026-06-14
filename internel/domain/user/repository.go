package user

import (
	"library-management/internel/domain/auth"
	"library-management/pkg/querybuilder"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers(ctx fiber.Ctx, opts querybuilder.QueryParams) ([]*auth.User, int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllUsers(ctx fiber.Ctx, opts querybuilder.QueryParams) ([]*auth.User, int64, error) {
	var events []*auth.User
	var total int64

	qb := querybuilder.New(r.db.Model(&auth.User{}))

	qb.Search(opts.Search, "name", "email")
	qb.Sort(opts.SortBy, opts.Order)
	qb.Filter(ctx)

	qb.DB.Count(&total)

	qb.Paginate(opts.Page, opts.Limit)

	err := qb.DB.Find(&events).Error

	return events, total, err
}
