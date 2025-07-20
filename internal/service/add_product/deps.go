//go:generate mockgen -destination=mock_deps_test.go -package=${GOPACKAGE} -source=deps.go
package add_product

import (
	"context"

	"VK_test_proect/internal/model"

	"github.com/google/uuid"
)

type productRepo interface {
	Save(ctx context.Context, product model.Product) error
}

type userRepo interface {
	Exists(ctx context.Context, id uuid.UUID) error
}
