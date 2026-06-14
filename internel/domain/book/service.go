package book

import (
	"library-management/internel/domain/book/dto"
	"library-management/pkg/utils"

	"github.com/google/uuid"
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

func (s *service) CreateBook(req *dto.CreateBookRequest) (*dto.BookResponse, error) {
	var err error
	book := Book{
		ID:          uuid.New(),
		Title:       req.Title,
		Author:      req.Author,
		TotalCopies: req.TotalCopies,
	}

	err = s.repo.CreateBook(&book)
	if err != nil {
		return nil, err
	}

	return book.ToResponse(), nil
}
