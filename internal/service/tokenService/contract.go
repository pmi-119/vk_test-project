package token_service

import (
	"VK_test_proect/internal/model"

	"github.com/google/uuid"
)

type Token interface {
	Generate(user model.User) (string, error)
	Validate(token string) (uuid.UUID, error)
}
