package product_info

type PriceFilter struct {
	Min *int
	Max *int
}

type Sorting struct {
	SortingByColumn string // "created_at"
	SortingOrder    string // ASC, DESC
}

type Paging struct {
	Offset int
	Limit  int
}
