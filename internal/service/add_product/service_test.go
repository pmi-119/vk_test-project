package add_product

import (
	"context"
	"testing"
	"time"

	"VK_test_proect/internal/repository"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestService_AddingProduct_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := NewMockproductRepo(ctrl)
	mockUserRepo := NewMockuserRepo(ctrl)

	svc := New(mockProductRepo, mockUserRepo)

	userID := uuid.New()
	ctx := context.Background()

	// входные данные
	in := In{
		Title:       "Тестовое объявление",
		Description: "Описание объявления минимум из 10 символов",
		ImageURL:    "https://example.com/image.jpg",
		Price:       999.99,
		UserID:      userID,
	}

	// Ожидания
	mockUserRepo.EXPECT().
		Exists(ctx, userID).
		Return(nil)

	mockProductRepo.EXPECT().
		Save(ctx, gomock.Any()).
		Return(nil)

	// Вызов
	out, err := svc.AddingProduct(ctx, in)

	require.NoError(t, err)
	require.Equal(t, in.Title, out.Product.Title)
	require.Equal(t, in.Description, out.Product.Description)
	require.Equal(t, in.ImageURL, out.Product.ImageUrl)
	require.Equal(t, in.Price, out.Product.Price)
	require.WithinDuration(t, time.Now(), out.Product.CreatedAt, time.Second)
}

func TestService_AddingProduct_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := NewMockproductRepo(ctrl)
	mockUserRepo := NewMockuserRepo(ctrl)

	svc := New(mockProductRepo, mockUserRepo)

	userID := uuid.New()
	ctx := context.Background()

	in := In{
		Title:       "Some title",
		Description: "Valid description here",
		ImageURL:    "https://example.com/image.jpg",
		Price:       50.0,
		UserID:      userID,
	}

	mockUserRepo.EXPECT().
		Exists(ctx, userID).
		Return(repository.ErrNotFound)

	_, err := svc.AddingProduct(ctx, in)
	require.ErrorContains(t, err, "user not found")
}
