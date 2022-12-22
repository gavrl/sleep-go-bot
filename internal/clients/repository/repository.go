package repository

import (
	"github.com/jmoiron/sqlx"
)

const sleepRateTable = "sleep_rate"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type SleepRateRepository interface {
}

type Repository struct {
	SleepRateRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SleepRateRepository: NewSleepRatePostgres(db),
	}
}
