package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"app/internal/config"
)

type Store struct {
	DB *sqlx.DB
}

func New(cfg config.Datebase) (*Store, error) {
	dsn := fmt.Sprint(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Database,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Open")
	}
	return &Store{
		DB: db,
	}, nil
}
