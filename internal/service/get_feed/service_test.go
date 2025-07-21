package get_feed

import (
	"context"
	"testing"

	"VK_test_proect/internal/model"
	"VK_test_proect/internal/repository/product_info"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService_GetFeed_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockproductRepo(ctrl)
	service := New(mockRepo)

	// входные данные
	userID := uuid.New()
	minPrice := int(100)
	maxPrice := int(500)
	in := In{
		UserID: &userID,
		PriceFilter: &PriceFilter{
			Min: &minPrice,
			Max: &maxPrice,
		},
		Sorting: &Sorting{
			Column: "price",
			Order:  "DESC",
		},
		Paging: &Paging{
			Limit: 10,
			Page:  1,
		},
	}

	expectedFilter := &product_info.PriceFilter{
		Min: &minPrice,
		Max: &maxPrice,
	}
	expectedSorting := product_info.Sorting{
		Column: "price",
		Order:  "DESC",
	}
	expectedPaging := product_info.Paging{
		Limit:  10,
		Offset: 10, // Page * Limit = 1 * 10
	}

	mockData := []model.ProductWithUser{
		{
			Title:       "Test Product",
			Description: "Test Desc",
			ImageUrl:    "http://image.com",
			Price:       123.45,
			UserLogin:   "test_user",
			UserID:      userID,
		},
	}

	mockRepo.EXPECT().
		Select(gomock.Any(), expectedFilter, expectedSorting, expectedPaging).
		Return(mockData, nil)

	out, err := service.GetFeed(context.Background(), in)

	assert.NoError(t, err)
	assert.Len(t, out.Items, 1)
	assert.Equal(t, "Test Product", out.Items[0].Title)
	assert.True(t, out.Items[0].IsCurrentUser)
}
