package loan

import (
	"errors"
	"library-management/internel/domain/auth"
	"library-management/internel/domain/book"
	"library-management/internel/domain/loan/dto"
	"library-management/pkg/utils"

	"github.com/google/uuid"
)


var ErrNoAvailableCopies = errors.New("No available copies for this book")

type service struct {
	repo       Repository
	authRepo   auth.Repository
	bookRepo   book.Repository
	jwtService utils.JwtService
}

func NewService(
	repo Repository,
	authRepo auth.Repository,
	bookRepo book.Repository,
	jwtService utils.JwtService,
) *service {
	return &service{
		repo:       repo,
		authRepo:   authRepo,
		bookRepo:   bookRepo,
		jwtService: jwtService,
	}
}

func (s *service) CreateLoan(req *dto.CreateLoanRequest) (*dto.LoanResponse, error) {
	loan := Loan{
		ID:     uuid.New(),
		UserID: req.UserID,
		BookID: req.BookID,
	}

	_, err := s.authRepo.GetById(loan.UserID)
	if err != nil {
		return nil, err
	}

	bookData, err := s.bookRepo.GetBookByID(loan.BookID)
	if err != nil {
		return nil, err
	}

	if bookData.AvailableCopies <= 0 {
		return nil, ErrNoAvailableCopies
	}

	if err := s.repo.CreateLoan(&loan); err != nil {
		return nil, err
	}

	bookData.AvailableCopies--

	if err := s.bookRepo.UpdateBook(bookData); err != nil {
		return nil, err
	}

	return &dto.LoanResponse{
		ID:         loan.ID,
		UserID:     loan.UserID,
		BookID:     loan.BookID,
		BorrowedAt: loan.BorrowedAt,
		ReturnedAt: loan.ReturnedAt,
	}, nil
}
