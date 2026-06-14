package user

import (
	"library-management/internel/domain/auth/dto"
	"library-management/pkg/httpresponse"
	"library-management/pkg/querybuilder"
	"library-management/pkg/utils"

	"github.com/gofiber/fiber/v3"
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

func (s *service) GetAllUsers(ctx fiber.Ctx, req querybuilder.QueryParams) (*dto.PaginationResponse, error) {
	users, total, err := s.repo.GetAllUsers(ctx, querybuilder.QueryParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
		SortBy: req.SortBy,
		Order:  req.Order,
	})
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		responses = append(responses, *user.ToResponse())
	}

	totalPages := int64((int(total) + req.Limit - 1) / req.Limit)

	return &dto.PaginationResponse{
		Data: responses,
		Meta: httpresponse.Meta{
			Page:       req.Page,
			Limit:      req.Limit,
			Total:      total,
			TotalPage: totalPages,
		},
	}, nil
}
