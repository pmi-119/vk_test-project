package register

import (
	"context"

	"VK_test_proect/internal/model"
	contract "VK_test_proect/internal/service/register"
)

type service interface {
	RegisterUser(ctx context.Context, in contract.In) (model.User, error)
}
