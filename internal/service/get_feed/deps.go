//go:generate mockgen -destination=mock_deps_test.go -package=${GOPACKAGE} -source=deps.go
package get_feed

import (
	"context"

	"VK_test_proect/internal/model"
	repo "VK_test_proect/internal/repository/product_info"
)

type productRepo interface {
	Select(ctx context.Context, priceFilter *repo.PriceFilter, sort repo.Sorting, paging repo.Paging) ([]model.ProductInfo, error)
}
