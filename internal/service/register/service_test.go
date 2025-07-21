package register

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"VK_test_proect/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService_RegisterUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)
	service := New(mockRepo)

	input := In{
		Login:    "testuser",
		Email:    "user@example.com",
		Password: "Pass1234",
	}

	// Здесь важно использовать gomock.Matching по полям, которые известны
	mockRepo.EXPECT().
		Save(gomock.Any(), gomock.AssignableToTypeOf(model.User{})).
		Return(nil)

	user, err := service.RegisterUser(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, input.Login, user.Login)
	assert.Equal(t, input.Email, user.Email)
	assert.Equal(t, input.Password, user.Password)
	assert.True(t, isUUIDv4(user.Id.String()))
}

func TestService_RegisterUser_InvalidEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)
	service := New(mockRepo)

	input := In{
		Login:    "testuser",
		Email:    "invalid-email",
		Password: "Pass1234",
	}

	user, err := service.RegisterUser(context.Background(), input)

	assert.ErrorIs(t, err, ErrInvalidEmail)
	assert.Equal(t, uuid.Nil, user.Id)
}

func TestService_RegisterUser_WeakPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)
	service := New(mockRepo)

	input := In{
		Login:    "testuser",
		Email:    "user@example.com",
		Password: "weak", // no digits or too short
	}

	user, err := service.RegisterUser(context.Background(), input)

	assert.ErrorIs(t, err, ErrWeakPassword)
	assert.Equal(t, uuid.Nil, user.Id)
}

func TestService_RegisterUser_SaveFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockuserRepo(ctrl)
	service := New(mockRepo)

	input := In{
		Login:    "testuser",
		Email:    "user@example.com",
		Password: "Pass1234",
	}

	mockRepo.EXPECT().
		Save(gomock.Any(), gomock.AssignableToTypeOf(model.User{})).
		Return(errors.New("db error"))

	user, err := service.RegisterUser(context.Background(), input)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "db error")
	assert.Equal(t, uuid.Nil, user.Id)
}

// Вспомогательная проверка UUID
func isUUIDv4(id string) bool {
	r := regexp.MustCompile(`^[a-f\d]{8}-([a-f\d]{4}-){3}[a-f\d]{12}$`)
	return r.MatchString(id)
}
