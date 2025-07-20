package authorize

import (
	"context"

	contract "VK_test_proect/internal/service/authorize"
)

type service interface {
	Authorize(ctx context.Context, in contract.In) (contract.Out, error)
}
