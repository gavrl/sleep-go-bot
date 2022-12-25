package internal

import "time"

const SleepRateDateFormat = "2000-01-01"

type SleepRate struct {
	Id        int           `db:"id"`
	UserName  string        `db:"username"`
	Rate      int           `db:"rate"`
	Calories  int           `db:"calories"`
	SleepTime time.Duration `db:"username"`
	date      time.Time     `db:"username"`
}
