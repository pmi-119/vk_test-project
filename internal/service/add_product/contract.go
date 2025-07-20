package add_product

import (
	"VK_test_proect/internal/model"

	"github.com/google/uuid"
)

type In struct {
	Title       string
	Description string
	ImageURL    string
	Price       float64
	UserID      uuid.UUID
}

type Out struct {
	Product model.Product
}
