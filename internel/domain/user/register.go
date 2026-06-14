package user

import (
	"library-management/internel/domain/auth"
	"library-management/pkg/middlewares"
	"library-management/pkg/utils"
	"library-management/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository, jwt)
	handler := NewHandler(service)

	router := api.Group("/user")

	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Get("",
		middlewares.Authentication(jwt),
		middlewares.Authorization(string(auth.RoleAdmin)),
		handler.GetAllUsers,
	)
}
