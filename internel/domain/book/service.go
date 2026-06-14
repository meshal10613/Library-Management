package book

import (
	"library-management/internel/domain/book/dto"
	"library-management/pkg/httpresponse"
	"library-management/pkg/querybuilder"
	"library-management/pkg/utils"

	"github.com/gofiber/fiber/v3"
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

func (s *service) GetBookByID(id uuid.UUID) (*dto.BookResponse, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	return book.ToResponse(), nil
}

func (s *service) GetAllBooks(ctx fiber.Ctx, req querybuilder.QueryParams) (*dto.BookPaginationResponse, error) {
	books, total, err := s.repo.GetAllBooks(ctx, querybuilder.QueryParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
		SortBy: req.SortBy,
		Order:  req.Order,
	})
	if err != nil {
		return nil, err
	}

	response := make([]dto.BookResponse, 0, len(books))

	for _, book := range books {
		response = append(response, *book.ToResponse())
	}

	totalPages := int64((int(total) + req.Limit - 1) / req.Limit)

	return &dto.BookPaginationResponse{
		Data: response,
		Meta: httpresponse.Meta{
			Page:      req.Page,
			Limit:     req.Limit,
			Total:     total,
			TotalPage: totalPages,
		},
	}, nil
}

func (s *service) UpdateBook(id uuid.UUID, req *dto.UpdateBookRequest) (*dto.BookResponse, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		book.Title = *req.Title
	}

	if req.Author != nil {
		book.Author = *req.Author
	}

	if req.TotalCopies != nil {
		borrowed := book.TotalCopies - book.AvailableCopies

		book.TotalCopies = *req.TotalCopies
		book.AvailableCopies = *req.TotalCopies - borrowed

		if book.AvailableCopies < 0 {
			book.AvailableCopies = 0
		}
	}

	err = s.repo.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	return book.ToResponse(), nil
}

func (s *service) DeleteBook(id uuid.UUID) error {
	return s.repo.DeleteBook(id)
}
