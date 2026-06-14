package user

import (
	"library-management/internel/domain/auth/dto"
	"library-management/pkg/utils"
)

type service struct {
	repo       Repository
	jwtService utils.JwtService
}

func NewService(repo Repository, jwtService utils.JwtService) *service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *service) GetAllUsers() (*[]dto.UserResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		responses = append(responses, *user.ToResponse())
	}

	return &responses, nil
}
