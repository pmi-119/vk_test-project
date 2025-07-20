package api

import (
	"time"

	"github.com/google/uuid"
)

type ProductDTOIn struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	Price       float64   `json:"price"`
	UserID      uuid.UUID `json:"user_id"`
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

type (
	FilterDTO struct {
		PriceFilter struct {
			Min *int `json:"min"`
			Max *int `json:"max"`
		}
		Sorting struct {
			SortingByColumn string `json:"sortingbycolumn"`
			SortingOrder    string `json:"sortingorder"`
		}
		Paging struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
		}
	}
)

//------------------

type ProductDTOInfo struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	Price       float64 `json:"price"`
	UserLogin   string  `json:"userlogin"`
}
