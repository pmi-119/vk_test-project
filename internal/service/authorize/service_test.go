package authorize

import (
	"context"
	"errors"
	"testing"

	"VK_test_proect/internal/model"
	"VK_test_proect/internal/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := NewMockuserRepo(ctrl)
	mockTokenService := NewMocktokenService(ctrl)
	service := New(mockTokenService, mockUserRepo)

	ctx := context.Background()
	in := In{
		Email:    "test@example.com",
		Password: "password123",
	}

	user := model.User{Email: in.Email}
	expectedToken := "jwt-token"

	mockUserRepo.EXPECT().
		GetUserByEmailAndPassword(ctx, in.Email, in.Password).
		Return(user, nil)

	mockTokenService.EXPECT().
		Generate(user).
		Return(expectedToken, nil)

	out, err := service.Authorize(ctx, in)
	assert.NoError(t, err)
	assert.Equal(t, Out{Token: expectedToken}, out)
}

func TestAuthorize_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := NewMockuserRepo(ctrl)
	mockTokenService := NewMocktokenService(ctrl)
	service := New(mockTokenService, mockUserRepo)

	ctx := context.Background()
	in := In{
		Email:    "missing@example.com",
		Password: "wrongpass",
	}

	mockUserRepo.EXPECT().
		GetUserByEmailAndPassword(ctx, in.Email, in.Password).
		Return(model.User{}, repository.ErrNotFound)

	out, err := service.Authorize(ctx, in)
	assert.ErrorIs(t, err, ErrUserNotFound)
	assert.Empty(t, out.Token)
}

func TestAuthorize_UserRepoOtherError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := NewMockuserRepo(ctrl)
	mockTokenService := NewMocktokenService(ctrl)
	service := New(mockTokenService, mockUserRepo)

	ctx := context.Background()
	in := In{
		Email:    "fail@example.com",
		Password: "pass",
	}

	expectedErr := errors.New("db error")

	mockUserRepo.EXPECT().
		GetUserByEmailAndPassword(ctx, in.Email, in.Password).
		Return(model.User{}, expectedErr)

	out, err := service.Authorize(ctx, in)
	assert.EqualError(t, err, expectedErr.Error())
	assert.Empty(t, out.Token)
}

func TestAuthorize_TokenGenerationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := NewMockuserRepo(ctrl)
	mockTokenService := NewMocktokenService(ctrl)
	service := New(mockTokenService, mockUserRepo)

	ctx := context.Background()
	in := In{
		Email:    "test@example.com",
		Password: "password123",
	}

	user := model.User{Email: in.Email}
	tokenErr := errors.New("token error")

	mockUserRepo.EXPECT().
		GetUserByEmailAndPassword(ctx, in.Email, in.Password).
		Return(user, nil)

	mockTokenService.EXPECT().
		Generate(user).
		Return("", tokenErr)

	out, err := service.Authorize(ctx, in)
	assert.EqualError(t, err, tokenErr.Error())
	assert.Empty(t, out.Token)
}
