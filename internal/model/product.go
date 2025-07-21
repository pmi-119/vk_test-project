package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Description string
	ImageUrl    string
	Price       float64

	CreatedAt time.Time
}

type ProductInfo struct {
	Title         string
	Description   string
	ImageUrl      string
	Price         float64
	UserLogin     string
	IsCurrentUser bool
}

type ProductWithUser struct {
	Title       string
	Description string
	ImageUrl    string
	Price       float64
	UserLogin   string
	UserID      uuid.UUID
}
