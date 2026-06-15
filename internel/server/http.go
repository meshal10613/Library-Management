package server

import (
	"fmt"
	"library-management/internel/config"
	"library-management/internel/domain/auth"
	"library-management/internel/domain/book"
	"library-management/internel/domain/loan"
	"library-management/internel/routes"
	"library-management/pkg/httpresponse"
	"library-management/pkg/seed"
	"library-management/pkg/utils"
	"library-management/pkg/validation"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB, cfg *config.Config) {
	app := fiber.New(fiber.Config{
		// Return structured JSON errors instead of plain-text HTML error pages
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := "Internal Server Error"

			if ok := false; !ok {
				if fe, ok2 := err.(*fiber.Error); ok2 {
					code = fe.Code
					msg = fe.Message
					_ = fe
				}
			}

			return c.Status(code).JSON(httpresponse.Error{
				Success: false,
				Message: msg,
				Details: err.Error(),
			})
		},
	})

	//? Migrations
	color.Cyan("⏳ Running database migrations...")
	if err := db.AutoMigrate(&auth.User{}, &book.Book{}, &loan.Loan{}); err != nil {
		color.Red("❌ Migration failed: %v", err)
		return
	}
	color.Green("✅ Migrations completed")

	// ── Seed admin ────────────────────────────────────────────────────────────
	color.Cyan("⏳ Checking admin seed...")
	seed.SeedAdmin(db, seed.AdminSeed{
		Name:     cfg.AdminName,
		Email:    cfg.AdminEmail,
		Password: cfg.AdminPassword,
	})

	//? Validator & JWT
	v := validator.New()
	customValidator := validation.NewCustomValidator(v)
	jwtService := utils.NewJwtService("")

	//? Global middleware
	app.Use(func(c fiber.Ctx) error {
		c.Locals("validator", customValidator)
		return c.Next()
	})

	//? Routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(httpresponse.Success{
			Success: true,
			Message: "Hello, World 👋!",
		})
	})

	api := app.Group("/api/v1")
	routes.RegisterRoutes(api, db, customValidator, jwtService)

	//? Start
	port := fmt.Sprintf(":%s", cfg.Port)
	color.Green("✅ Database connected successfully")
	color.Cyan("🚀 Server running on http://localhost%s", port)

	if err := app.Listen(port); err != nil {
		color.Red("❌ Failed to start server: %v", err)
	}
}
