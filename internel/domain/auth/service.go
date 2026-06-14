package auth

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

func (s *service) RegisterUser(req *dto.RegisterUserRequest) (*dto.UserTokenResponse, error) {
	var err error
	user := User{
		Name:  req.Name,
		Email: req.Email,
	}

	//? Hash password before saving to database
	err = user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.Register(&user)
	if err != nil {
		return nil, err
	}

	// //? generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Name, user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	return user.ToResponse(token), nil
}

func (s *service) LoginUser(req *dto.LoginUserRequest) (*dto.UserTokenResponse, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = user.checkPassword(req.Password)
	if user == nil || err != nil {
		return nil, ErrorInvalidCredentials
	}

	//? generate token
	token, err := s.jwtService.GenerateToken(user.ID, user.Name, user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	return user.ToResponse(token), nil
}

func (s *service) GetMe(Email string) (*dto.UserResponse, error) {
	user, err := s.repo.GetByEmail(Email)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
