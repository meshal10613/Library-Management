package auth

import (
	"library-management/pkg/middlewares"
	"library-management/pkg/utils"
	"library-management/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func AuthRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	v.RegisterValidation("password", validation.PasswordValidation)

	repository := NewRepository(db)
	service := NewService(repository, jwt)
	handler := NewHandler(service)

	router := api.Group("/auth")

	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Post("/register", handler.RegisterUser)
	router.Post("/login", handler.LoginUser)

	router.Get("/me",
		middlewares.Authentication(jwt),
		middlewares.Authorization(string(RoleAdmin)),
		handler.GetMe,
	)
}
