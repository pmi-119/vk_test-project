package user_info

import (
	"context"
	"database/sql"
	"errors"

	"VK_test_proect/internal/model"
	"VK_test_proect/internal/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRow struct {
	Id       uuid.UUID `db:"id"`
	Login    string    `db:"login"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Save(ctx context.Context, user model.User) error {
	args := map[string]interface{}{
		"id":       user.Id.String(),
		"login":    user.Login,
		"email":    user.Email,
		"password": user.Password,
	}

	query := `INSERT INTO users (
		id, login, email, password
	) VALUES (
		:id, :login, :email, :password
	)`

	_, err := sqlx.NamedExecContext(ctx, r.db, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (model.User, error) {
	var row userRow

	query := "SELECT * FROM users WHERE email = $1 AND password = $2"

	err := r.db.GetContext(ctx, &row, query, email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, repository.ErrNotFound
		}

		return model.User{}, err
	}

	return toModel(row), nil
}

func (r *Repository) Exists(ctx context.Context, id uuid.UUID) error {
	query := `SELECT true FROM users WHERE id = $1`

	var dest bool

	err := r.db.GetContext(ctx, &dest, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}

		return err
	}

	return nil
}

func toModel(in userRow) model.User {
	return model.User{
		Id:       in.Id,
		Login:    in.Login,
		Email:    in.Email,
		Password: in.Password,
	}
}
