package middlewares

import (
	"library-management/pkg/httpresponse"

	"github.com/gofiber/fiber/v3"
)

func Authorization(roles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(httpresponse.Error{
				Success: false,
				Message: "Unauthorized access.",
			})
		}

		for _, allowedRole := range roles {
			if role == allowedRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(httpresponse.Error{
			Success: false,
			Message: "Forbidden. You don't have sufficient permissions to access this resource.",
		})
	}
}
