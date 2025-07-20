package register

import (
	"context"
	"errors"
	"testing"

	"VK_test_proect/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)

	svc := New(mockRepo)

	ctx := context.Background()
	email := "test@example.com"
	password := "password123"

	in := In{
		Email:    email,
		Password: password,
	}

	// ожидаем, что Save будет вызван с правильными аргументами
	mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).DoAndReturn(
		func(ctx context.Context, user model.User, pass string) error {
			assert.Equal(t, email, user.Email)
			return nil
		},
	)

	out, err := svc.RegisterUser(ctx, in)

	assert.NoError(t, err)
	assert.Equal(t, email, out.Email)
}

func TestRegisterUser_InvalidEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)

	svc := New(mockRepo)

	in := In{
		Email:    "invalid-email",
		Password: "password123",
	}

	_, err := svc.RegisterUser(context.Background(), in)
	assert.Equal(t, ErrInvalidEmail, err)
}

func TestRegisterUser_WeakPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)

	svc := New(mockRepo)

	in := In{
		Email:    "test@example.com",
		Password: "123", // слабый пароль
	}

	_, err := svc.RegisterUser(context.Background(), in)
	assert.Equal(t, ErrWeakPassword, err)
}

func TestRegisterUser_RepoFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)

	svc := New(mockRepo)

	in := In{
		Email:    "test@example.com",
		Password: "password123",
	}

	expectedErr := errors.New("db error")

	mockRepo.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(expectedErr)

	_, err := svc.RegisterUser(context.Background(), in)
	assert.Equal(t, expectedErr, err)
}
