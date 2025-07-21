package product_info

type PriceFilter struct {
	Min *int
	Max *int
}

type Sorting struct {
	Column string // "created_at"
	Order  string // ASC, DESC
}

type Paging struct {
	Offset int
	Limit  int
}
