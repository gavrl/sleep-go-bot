package repository

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/gavrl/sleep-go-bot/internal"
	"github.com/gavrl/sleep-go-bot/internal/dto"
	"github.com/jmoiron/sqlx"
)

type SleepRatePostgres struct {
	db *sqlx.DB
}

func NewSleepRatePostgres(db *sqlx.DB) *SleepRatePostgres {
	return &SleepRatePostgres{db: db}
}

func (r *SleepRatePostgres) Save(dto *dto.SaveSleepRateDto) (int, error) {
	var id int
	saveSleepRateQuery := fmt.Sprintf(
		"INSERT INTO %s (username, rate, date) VALUES ($1, $2, $3) RETURNING id",
		sleepRateTable,
	)
	logrus.Debug(dto.Time.Format(internal.SleepRateDateFormat))
	row := r.db.QueryRow(saveSleepRateQuery, dto.UserName, dto.Rate, dto.Time.Format(internal.SleepRateDateFormat))
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SleepRatePostgres) Get(dto *dto.GetSleepRateDto) (internal.SleepRate, error) {
	var sleepRate internal.SleepRate
	query := fmt.Sprintf(
		"SELECT id, username, rate, calories, sleep_time, date FROM %s where username = $1 and date = $2",
		sleepRateTable,
	)
	err := r.db.Get(&sleepRate, query, dto.Username, dto.Time.Format(internal.SleepRateDateFormat))
	if err != nil {
		return sleepRate, err
	}
	return internal.SleepRate{}, nil
}
