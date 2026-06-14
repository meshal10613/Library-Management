package routes

import (
	"library-management/internel/domain/auth"
	"library-management/pkg/utils"
	"library-management/pkg/validation"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterRoutes(
	api fiber.Router,
	db *gorm.DB,
	v *validation.CustomValidator,
	jwt utils.JwtService,
) {
	auth.AuthRoutes(api, db, v, jwt)
}
