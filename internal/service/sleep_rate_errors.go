package service

import (
	"fmt"
	"time"

	"github.com/gavrl/sleep-go-bot/internal"
)

type NotExistsSleepRate struct {
	Username string
	Time     time.Time
}

func (e NotExistsSleepRate) Error() string {
	return fmt.Sprintf("не найдено показателей сна для %s %s", e.Username, e.Time.Format(internal.SleepRateDateFormat))
}
