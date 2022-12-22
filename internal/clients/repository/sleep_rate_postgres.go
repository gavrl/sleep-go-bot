package repository

import "github.com/jmoiron/sqlx"

type SleepRatePostgres struct {
	db *sqlx.DB
}

func NewSleepRatePostgres(db *sqlx.DB) *SleepRatePostgres {
	return &SleepRatePostgres{db: db}
}
