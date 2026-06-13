package auth

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func SetAuthCookie(c fiber.Ctx, token string) {
	tokenDuration := 24 * time.Hour // default 24h

	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true, // set false in local dev (http)
		SameSite: "Lax",
		MaxAge:   int(tokenDuration.Seconds()),
	}

	c.Cookie(cookie)
}

func ClearAuthCookie(c fiber.Ctx) {
	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		MaxAge:   -1,
	}

	c.Cookie(cookie)
}
