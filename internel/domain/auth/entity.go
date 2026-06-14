package auth

import (
	"library-management/internel/domain/auth/dto"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleLibrarian Role = "librarian"
	RolePublic    Role = "public"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Role     Role   `json:"role" gorm:"type:varchar(20);default:public"`
}

func (u *User) hashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) checkPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (e *User) ToResponse(token string) *dto.UserTokenResponse {
	return &dto.UserTokenResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        e.ID,
			Name:      e.Name,
			Email:     e.Email,
			Role:      string(e.Role),
			CreatedAt: e.CreatedAt.String(),
			UpdatedAt: e.UpdatedAt.String(),
		},
	}
}
