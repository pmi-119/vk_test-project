package add_product

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"VK_test_proect/internal/model"
	"VK_test_proect/internal/repository"

	"github.com/google/uuid"
)

type Service struct {
	product productRepo
	user    userRepo
}

func New(product productRepo, user userRepo) *Service {
	return &Service{
		product: product,
		user:    user,
	}
}

func (s *Service) AddingProduct(ctx context.Context, in In) (Out, error) {
	err := s.user.Exists(ctx, in.UserID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return Out{}, errors.New("user not found")
		}
		return Out{}, err
	}

	if len(in.Title) < 5 || len(in.Title) > 100 {
		return Out{}, errors.New("title must be between 5 and 100 characters")
	}
	if len(in.Description) < 10 || len(in.Description) > 1000 {
		return Out{}, errors.New("description must be between 10 and 1000 characters")
	}
	if in.Price <= 0 || in.Price > 1_000_000 {
		return Out{}, errors.New("price must be a positive number less than 1,000,000")
	}
	if !strings.HasPrefix(in.ImageURL, "http") ||
		!regexp.MustCompile(`(?i)\.(jpg|jpeg|png|webp|gif)$`).MatchString(in.ImageURL) {
		return Out{}, errors.New("invalid image URL format")
	}

	product := model.Product{
		Id:          uuid.New(),
		UserID:      uuid.New(),
		Title:       in.Title,
		Description: in.Description,
		ImageUrl:    in.ImageURL,
		Price:       in.Price,
		CreatedAt:   time.Now().UTC(),
	}

	if err := s.product.Save(ctx, product); err != nil {
		return Out{}, err
	}
	return Out{Product: product}, nil
}
