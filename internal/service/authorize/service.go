package authorize

import (
	"context"
	"errors"

	"VK_test_proect/internal/repository"
)

var ErrUserNotFound = errors.New("user not found")

type Service struct {
	tokenService tokenService
	userRepo     userRepo
}

func New(tokenService tokenService, userRepo userRepo) *Service {
	return &Service{
		tokenService: tokenService,
		userRepo:     userRepo,
	}
}

func (s *Service) Authorize(ctx context.Context, in In) (Out, error) {
	user, err := s.userRepo.GetUserByEmailAndPassword(ctx, in.Email, in.Password)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return Out{}, ErrUserNotFound
		}

		return Out{}, err
	}

	tokenString, err := s.tokenService.Generate(user)
	if err != nil {
		return Out{}, err
	}

	return Out{Token: tokenString}, nil
}
