package postgres

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type Migrator struct {
	db *sqlx.DB
}

func NewMigrator(db *sqlx.DB) *Migrator {
	return &Migrator{
		db: db,
	}
}

func (m *Migrator) Init() error {
	bytes, err := os.ReadFile("0001_create_table.sql")
	if err != nil {
		return err
	}

	_, err = m.db.Exec(string(bytes))

	return err
}
