package book

import (
	"library-management/internel/domain/auth"
	"library-management/pkg/middlewares"
	"library-management/pkg/utils"
	"library-management/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func BookRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	repository := NewRepository(db)
	service := NewService(repository, jwt)
	handler := NewHandler(service)

	router := api.Group("/book")

	router.Use(func(c fiber.Ctx) error {
		c.Locals("validator", v)
		return c.Next()
	})

	router.Post("",
		middlewares.Authentication(jwt),
		middlewares.Authorization(string(auth.RoleAdmin), string(auth.RoleLibrarian)),
		handler.CreateBook,
	)

	router.Get("",
		middlewares.Authentication(jwt),
		handler.GetAllBooks,
	)

	router.Get("/:id",
		middlewares.Authentication(jwt),
		handler.GetBookByID,
	)

	router.Patch("/:id",
		middlewares.Authentication(jwt),
		middlewares.Authorization(string(auth.RoleAdmin), string(auth.RoleLibrarian)),
		handler.UpdateBook,
	)

	router.Delete("/:id",
		middlewares.Authentication(jwt),
		middlewares.Authorization(string(auth.RoleAdmin), string(auth.RoleLibrarian)),
		handler.DeleteBook,
	)
}
