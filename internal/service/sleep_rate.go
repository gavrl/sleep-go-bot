package service

import (
	"database/sql"
	"errors"
	"github.com/gavrl/sleep-go-bot/pkg/e"

	"github.com/gavrl/sleep-go-bot/internal"
	"github.com/gavrl/sleep-go-bot/internal/dto"
	"github.com/gavrl/sleep-go-bot/internal/repository"
	validator "github.com/go-playground/validator/v10"
)

type SleepRateService struct {
	repo      repository.SleepRateRepository
	validator *validator.Validate
}

func NewSleepRateService(repo repository.SleepRateRepository, validator *validator.Validate) *SleepRateService {
	return &SleepRateService{repo, validator}
}

func (s *SleepRateService) Get(dto dto.GetSleepRateDto) (internal.SleepRate, error) {
	sleepRate, err := s.repo.Get(dto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return sleepRate, NotExistsSleepRate{Username: dto.Username, Time: dto.Time}
		} else {
			return sleepRate, err
		}
	}
	return sleepRate, nil
}

func (s *SleepRateService) Save(dto dto.SaveSleepRateDto) (int, error) {
	id, err := s.repo.Save(dto)
	if err != nil {
		return 0, e.Wrap("can't save sleep rate", err)
	}
	return id, nil
}
