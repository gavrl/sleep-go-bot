package dto

import "time"

type SaveSleepRateDto struct {
	UserName string
	Rate     int
	Time     time.Time
}
