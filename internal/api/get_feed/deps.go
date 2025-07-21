package get_feed

import (
	"context"

	contract "VK_test_proect/internal/service/get_feed"

	"github.com/google/uuid"
)

type service interface {
	GetFeed(ctx context.Context, in contract.In) (contract.Out, error)
}

type tokenService interface {
	Validate(token string) (uuid.UUID, error)
}
