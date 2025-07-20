package get_feed

import (
	"context"

	"VK_test_proect/internal/repository/product_info"
)

type Service struct {
	repo productRepo
}

func New(repo productRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetFeed(ctx context.Context, in In) (Out, error) {
	products, err := s.repo.Select(ctx, nil, product_info.Sorting{}, product_info.Paging{})
	if err != nil {
		return Out{}, nil
	}
	return Out{Items: products}, nil
}
