package service

import (
	"github.com/gavrl/sleep-go-bot/internal"
	"github.com/gavrl/sleep-go-bot/internal/dto"
	"github.com/gavrl/sleep-go-bot/internal/repository"
	validator "github.com/go-playground/validator/v10"
)

type SleepRate interface {
	Get(dto dto.GetSleepRateDto) (internal.SleepRate, error)
	Save(dto dto.SaveSleepRateDto) (int, error)
}

type Service struct {
	SleepRate
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SleepRate: NewSleepRateService(repo, validator.New()),
	}
}
