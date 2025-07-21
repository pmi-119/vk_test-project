package get_feed

import (
	"context"
	"errors"
	"slices"

	"VK_test_proect/internal/model"
	"VK_test_proect/internal/repository/product_info"

	"github.com/google/uuid"
)

const (
	sortingOrderASC  = "ASC"
	sortingOrderDESC = "DESC"
)

var (
	defaultSorting = product_info.Sorting{
		Column: "created_at",
		Order:  "ASC",
	}

	defaultPaging = product_info.Paging{
		Limit:  10,
		Offset: 0,
	}

	availableSortingColumns = []string{
		"created_at",
		"price",
		"title",
	}
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
	priceFilter, err := s.getPriceFilter(in.PriceFilter)
	if err != nil {
		return Out{}, err
	}

	sorting, err := s.getSorting(in.Sorting)
	if err != nil {
		return Out{}, err
	}

	paging, err := s.getPaging(in.Paging)
	if err != nil {
		return Out{}, err
	}

	products, err := s.repo.Select(ctx, priceFilter, sorting, paging)
	if err != nil {
		return Out{}, nil
	}

	return Out{Items: s.convertToProductWithUser(in.UserID, products)}, nil
}

func (s *Service) getPriceFilter(filter *PriceFilter) (*product_info.PriceFilter, error) {
	if filter == nil {
		return nil, nil
	}

	if filter.Max == nil && filter.Min == nil {
		return nil, errors.New("one of min or max must be provided")
	}

	if filter.Min != nil && *filter.Min < 0 {
		return nil, errors.New("min must be grater then zero")
	}

	if filter.Max != nil && *filter.Max < 0 {
		return nil, errors.New("max must be grater then zero")
	}

	if filter.Min != nil && filter.Max != nil && *filter.Max <= *filter.Min {
		return nil, errors.New("max must be greater then min")
	}

	return &product_info.PriceFilter{
		Min: filter.Min,
		Max: filter.Max,
	}, nil
}

func (s *Service) getSorting(sorting *Sorting) (product_info.Sorting, error) {
	if sorting == nil {
		return defaultSorting, nil
	}

	if sorting.Order != sortingOrderASC && sorting.Order != sortingOrderDESC {
		return product_info.Sorting{}, errors.New("invalid sorting order")
	}

	if !slices.Contains(availableSortingColumns, sorting.Column) {
		return product_info.Sorting{}, errors.New("invalid sorting column")
	}

	return product_info.Sorting{
		Column: sorting.Column,
		Order:  sorting.Order,
	}, nil
}

func (s *Service) getPaging(paging *Paging) (product_info.Paging, error) {
	if paging == nil {
		return defaultPaging, nil
	}

	if paging.Limit <= 0 {
		return product_info.Paging{}, errors.New("invalid limit")
	}

	if paging.Page < 0 {
		return product_info.Paging{}, errors.New("invalid page")
	}

	return product_info.Paging{
		Limit:  paging.Limit,
		Offset: paging.Page * paging.Limit,
	}, nil
}

func (s *Service) convertToProductWithUser(userID *uuid.UUID, products []model.ProductWithUser) []model.ProductInfo {
	res := make([]model.ProductInfo, 0, len(products))

	notNullableUserID := getValueOrDefault(userID)

	for _, product := range products {
		res = append(res, model.ProductInfo{
			Title:         product.Title,
			Description:   product.Description,
			ImageUrl:      product.ImageUrl,
			Price:         product.Price,
			UserLogin:     product.UserLogin,
			IsCurrentUser: product.UserID == notNullableUserID,
		})
	}

	return res
}

func getValueOrDefault(in *uuid.UUID) uuid.UUID {
	if in == nil {
		return uuid.UUID{}
	}

	return *in
}
