package api

import (
	"time"

	"github.com/google/uuid"
)

type ProductDTOIn struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	Price       float64 `json:"price"`
}

type ProductDTOOut struct {
	Id          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	Price       float64   `json:"price"`

	CreatedAt time.Time `json:"created_at"`
}

//-----------------

type UserDTOIn struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserDTOOut struct {
	Id       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
type Token = string

//------------------

type FilterDTO struct {
	PriceFilter *PriceFilter `json:"price_filter"`
	Sorting     *Sorting     `json:"sorting"`
	Paging      *Paging      `json:"paging"`
}

type PriceFilter struct {
	Min *int `json:"min"`
	Max *int `json:"max"`
}

type Sorting struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type Paging struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

//------------------

type ProductDTOInfo struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	ImageUrl      string  `json:"imageUrl"`
	Price         float64 `json:"price"`
	UserLogin     string  `json:"user_login"`
	IsCurrentUser bool    `json:"is_current_user"`
}
