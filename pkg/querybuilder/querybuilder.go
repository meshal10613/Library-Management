package querybuilder

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Builder struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Builder {
	return &Builder{DB: db}
}

func (b *Builder) Search(keyword string, columns ...string) *Builder {
	if keyword == "" {
		return b
	}

	var conditions []string
	var values []interface{}

	for _, col := range columns {
		conditions = append(conditions, col+" ILIKE ?")
		values = append(values, "%"+keyword+"%")
	}

	b.DB = b.DB.Where(strings.Join(conditions, " OR "), values...)

	return b
}

func (b *Builder) Sort(sortBy, order string) *Builder {
	if sortBy == "" {
		sortBy = "created_at"
	}

	if order != "asc" {
		order = "desc"
	}

	b.DB = b.DB.Order(sortBy + " " + order)

	return b
}

func (b *Builder) Paginate(page, limit int) *Builder {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	b.DB = b.DB.Offset(offset).Limit(limit)

	return b
}

func (b *Builder) Filter(c fiber.Ctx, exclude ...string) *Builder {
	query := c.Queries()

	// fields to ignore (pagination, search, etc.)
	skip := map[string]bool{
		"page":  true,
		"limit": true,
		"search": true,
		"sortBy": true,
		"order": true,
	}

	for _, e := range exclude {
		skip[e] = true
	}

	for key, value := range query {
		if skip[key] || value == "" {
			continue
		}

		// LIKE: name=john
		if strings.Contains(value, ":") {
			parts := strings.Split(value, ":")
			b.DB = b.DB.Where(fmt.Sprintf("%s ILIKE ?", parts[0]), "%"+parts[1]+"%")
			continue
		}

		// GTE: age>=18
		if strings.Contains(value, ">=") {
			parts := strings.Split(value, ">=")
			b.DB = b.DB.Where(fmt.Sprintf("%s >= ?", parts[0]), parts[1])
			continue
		}

		// LTE: age<=50
		if strings.Contains(value, "<=") {
			parts := strings.Split(value, "<=")
			b.DB = b.DB.Where(fmt.Sprintf("%s <= ?", parts[0]), parts[1])
			continue
		}

		// GT
		if strings.Contains(value, ">") {
			parts := strings.Split(value, ">")
			b.DB = b.DB.Where(fmt.Sprintf("%s > ?", parts[0]), parts[1])
			continue
		}

		// LT
		if strings.Contains(value, "<") {
			parts := strings.Split(value, "<")
			b.DB = b.DB.Where(fmt.Sprintf("%s < ?", parts[0]), parts[1])
			continue
		}

		// IN: role=admin,lib
		if strings.Contains(value, ",") {
			list := strings.Split(value, ",")
			b.DB = b.DB.Where(fmt.Sprintf("%s IN ?", key), list)
			continue
		}

		// default equals
		b.DB = b.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}

	return b
}