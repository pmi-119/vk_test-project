//go:generate mockgen -destination=mock_deps_test.go -package=${GOPACKAGE} -source=deps.go
package authorize

import (
	"context"

	"VK_test_proect/internal/model"
)

type tokenService interface {
	Generate(user model.User) (string, error)
}

type userRepo interface {
	GetUserByEmailAndPassword(ctx context.Context, email string, password string) (model.User, error)
}
