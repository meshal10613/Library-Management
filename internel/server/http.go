package server

import (
	"fmt"
	"library-management/internel/config"
	"library-management/internel/domain/user"
	"library-management/pkg/httpresponse"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB, cfg *config.Config) {
	app := fiber.New()

	// Run migrations
	if err := db.AutoMigrate(&user.User{}); err != nil {
		color.Red("❌ Failed to migrate database: %v", err)
		return
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(httpresponse.Success{
			Success: true,
			Message: "Hello, World 👋!",
		})
	})

	port := fmt.Sprintf(":%s", cfg.Port)

	color.Green("✅ Database connected successfully")
	color.Cyan("🚀 Server running on http://localhost%s", port)

	if err := app.Listen(port); err != nil {
		color.Red("❌ Failed to start server: %v", err)
	}
}
