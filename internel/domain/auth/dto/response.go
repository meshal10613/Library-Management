package dto

import (
	"library-management/pkg/httpresponse"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt string    `json:"created_at,omitempty"`
	UpdatedAt string    `json:"updated_at,omitempty"`
}

type UserTokenResponse struct {
	Token string       `json:"token,omitempty"`
	User  UserResponse `json:"user"`
}

type UserPaginationResponse struct {
	Data []UserResponse
	Meta httpresponse.Meta
}
