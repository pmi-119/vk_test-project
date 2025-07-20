package product_info

import (
	"context"
	"fmt"
	"log"
	"time"

	"VK_test_proect/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type productRow struct {
	Id          string  `db:"id"`
	UserID      string  `db:"user_id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	ImageUrl    string  `db:"image_url"`
	Price       float64 `db:"price"`

	CreatedAt time.Time `db:"created_at"`
}

func toRow(product model.Product) productRow {
	return productRow{
		Id:          product.Id.String(),
		UserID:      product.UserID.String(),
		Title:       product.Title,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
	}
}

type productInfoRow struct {
	Title       string  `db:"title"`
	Description string  `db:"description"`
	ImageUrl    string  `db:"image_url"`
	Price       float64 `db:"price"`
	Email       string  `db:"email"`
}

func toProductInfoModel(row productInfoRow) model.ProductInfo {
	return model.ProductInfo{
		Title:       row.Title,
		Description: row.Description,
		ImageUrl:    row.ImageUrl,
		Price:       row.Price,
		UserLogin:   row.Email,
	}
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Select(ctx context.Context, priceFilter *PriceFilter, sort Sorting, paging Paging) ([]model.ProductInfo, error) {
	// // Установите значения по умолчанию, если они пустые
	// if sort.SortingByColumn == "" {
	// 	sort.SortingByColumn = "created_at" // или "price", "title"
	// }
	// if sort.SortingOrder == "" {
	// 	sort.SortingOrder = "DESC" // или "ASC"
	// }
	// if paging.Limit <= 0 {
	// 	paging.Limit = 10
	// }
	// if paging.Offset < 0 {
	// 	paging.Offset = 0
	// }

	query := sq.Select("title, description, image_url, price, email").
		From("products").
		Join("users ON users.id = products.user_id")

	if priceFilter != nil {
		if priceFilter.Min != nil {
			query = query.Where(sq.GtOrEq{"price": *priceFilter.Min})
		}

		if priceFilter.Max != nil {
			query = query.Where(sq.LtOrEq{"price": *priceFilter.Max})
		}
	}

	query = query.OrderBy(sort.SortingByColumn + " " + sort.SortingOrder)
	query = query.Offset(uint64(paging.Offset))
	query = query.Limit(uint64(paging.Limit))

	sqlString, args, err := query.ToSql()
	if err != nil {
		fmt.Printf("SQL query: %s\nArgs: %v\n", sqlString, args)
		return nil, fmt.Errorf("error in building sql query: %w", err)
	}

	var rows []productInfoRow

	err = r.db.SelectContext(ctx, &rows, sqlString, args...)
	if err != nil {
		fmt.Printf("SQL query: %s\nArgs: %v\n", sqlString, args)
		fmt.Printf("Selected rows: %+v\n", rows)
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	res := make([]model.ProductInfo, 0, len(rows))
	for _, row := range rows {
		res = append(res, toProductInfoModel(row))
	}

	return res, nil
}

func (r *Repository) Save(ctx context.Context, product model.Product) error {
	row := toRow(product)
	query := `
		INSERT INTO products (
			id,
			user_id,
			title,
			description,
			image_url,
			price,
			created_at
		) VALUES (
			:id,
			:user_id,
			:title,
			:description,
			:image_url,
			:price,
			:created_at
		)`
	_, err := r.db.NamedExecContext(ctx, query, row)
	if err != nil {
		log.Printf("Error saving product: %v", err)
		return err
	}

	return nil
}
