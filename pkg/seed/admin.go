package seed

import (
	"errors"
	"library-management/internel/domain/auth"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminSeed struct {
	Name     string
	Email    string
	Password string
}

func SeedAdmin(db *gorm.DB, admin AdminSeed) {
	var existing auth.User

	err := db.Where("email = ?", admin.Email).First(&existing).Error

	// Admin already exists — skip seeding
	if err == nil {
		color.Yellow("⚠️  Admin already exists — skipping seed  [%s]", existing.Email)
		return 
	}

	// Unexpected DB error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		color.Red("❌ Failed to check admin existence: %v", err)
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		color.Red("❌ Failed to hash admin password: %v", err)
		return
	}

	user := auth.User{
		ID:       uuid.New(),
		Name:     admin.Name,
		Email:    admin.Email,
		Password: string(hash),
		Role:     auth.RoleAdmin,
	}

	if err := db.Create(&user).Error; err != nil {
		color.Red("❌ Failed to seed admin user: %v", err)
		return
	}

	color.Green("✅ Admin seeded successfully  [%s]", user.Email)
}
