package get_feed

import (
	"VK_test_proect/internal/model"

	"github.com/google/uuid"
)

type In struct {
	UserID      *uuid.UUID
	PriceFilter *PriceFilter
	Sorting     *Sorting
	Paging      *Paging
}

type Out struct {
	Items []model.ProductInfo
}

type PriceFilter struct {
	Min *int
	Max *int
}
type Sorting struct {
	Column string
	Order  string
}
type Paging struct {
	Page  int
	Limit int
}
