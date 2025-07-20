//go:generate mockgen -destination=mock_deps_test.go -package=${GOPACKAGE} -source=deps.go
package register

import (
	"context"

	"VK_test_proect/internal/model"
)

type userRepo interface {
	Save(ctx context.Context, user model.User) error
}
