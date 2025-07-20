package add_product

import (
	"context"

	contract "VK_test_proect/internal/service/add_product"

	"github.com/google/uuid"
)

type service interface {
	AddingProduct(ctx context.Context, in contract.In) (contract.Out, error)
}
type tokenService interface {
	Validate(token string) (uuid.UUID, error)
}
