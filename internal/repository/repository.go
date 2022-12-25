package repository

import (
	"github.com/gavrl/sleep-go-bot/internal"
	"github.com/gavrl/sleep-go-bot/internal/dto"
	"github.com/jmoiron/sqlx"
)

const (
	sleepRateTable = "sleep_rate"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type SleepRateRepository interface {
	Get(dto dto.GetSleepRateDto) (internal.SleepRate, error)
	Save(dto dto.SaveSleepRateDto) (int, error)
}

type Repository struct {
	SleepRateRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SleepRateRepository: NewSleepRatePostgres(db),
	}
}
