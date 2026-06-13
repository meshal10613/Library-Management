package user

import (
	"library-management/internel/domain/user/dto"

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
	Name     string `validate:"required,min=2,max=100"`
	Email    string `gorm:"unique" validate:"required,email"`
	Password string `validate:"required,password"`
	Role     Role   `gorm:"type:varchar(20);default:public"`
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
