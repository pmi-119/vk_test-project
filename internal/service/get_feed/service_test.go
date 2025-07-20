package get_feed

import (
	"context"
	"testing"

	"VK_test_proect/internal/model"
	"VK_test_proect/internal/repository/product_info"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetFeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockproductRepo(ctrl)

	// Подготовка входных и ожидаемых данных
	ctx := context.Background()
	expectedProducts := []model.ProductInfo{
		{
			Title:         "Test Product 1",
			Description:   "Description 1",
			ImageUrl:      "http://image1.com",
			Price:         100.0,
			UserLogin:     "user1",
			IsCurrentUser: true,
		},
		{
			Title:         "Test Product 2",
			Description:   "Description 2",
			ImageUrl:      "http://image2.com",
			Price:         200.0,
			UserLogin:     "user2",
			IsCurrentUser: false,
		},
	}
	// Настройка ожиданий
	mockRepo.
		EXPECT().
		Select(ctx, nil, product_info.Sorting{}, product_info.Paging{}).
		Return(expectedProducts, nil)

	service := New(mockRepo)

	// Вызов
	out, err := service.GetFeed(ctx, In{})

	// Проверки
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, out.Items)
}
