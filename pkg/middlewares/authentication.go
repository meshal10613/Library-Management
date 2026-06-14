package middlewares

import (
	"library-management/pkg/httpresponse"
	"library-management/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func Authentication(jwtService utils.JwtService) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get token from cookie
		tokenString := c.Cookies("token")

		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
				Success: false,
				Message: "Authentication required",
			})
		}

		// Validate token
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(httpresponse.Error{
				Success: false,
				Message: "Authentication failed. Invalid or expired token",
			})
		}

		// Store user information in context
		c.Locals("user_id", claims.UserID)
		c.Locals("name", claims.Name)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}
