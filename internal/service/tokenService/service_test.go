package token_service

import (
	"testing"

	"VK_test_proect/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	t.Parallel()

	service := Service{
		secret: "123secret",
	}

	testID := uuid.New()

	testUser := model.User{
		Id: testID,
	}

	got, err := service.Generate(testUser)
	assert.NoError(t, err)

	userID, err := service.Validate(got)
	assert.NoError(t, err)

	assert.Equal(t, testID, userID)
}
