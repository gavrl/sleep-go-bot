package main

import (
	tgClient "github.com/gavrl/sleep-go-bot/internal/clients/telegram"
	event_consumer "github.com/gavrl/sleep-go-bot/internal/consumer/event-consumer"
	"github.com/gavrl/sleep-go-bot/internal/events/telegram"
	"github.com/gavrl/sleep-go-bot/internal/repository"
	"github.com/gavrl/sleep-go-bot/internal/service"
	"github.com/gavrl/sleep-go-bot/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// todo перенести в env
const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func init() {
	util.InitConfig()
	util.InitLogger()
}

func main() {
	db := initDatabase()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	eventsProcessor := telegram.NewProcessor(
		tgClient.NewClient(tgBotHost, viper.GetString("TELEGRAM_BOT_TOKEN")),
		service,
	)

	logrus.Info("bot started")

	consumer := event_consumer.NewConsumer(
		eventsProcessor,
		eventsProcessor,
		batchSize,
	)

	if err := consumer.Start(); err != nil {
		logrus.Fatal("service is stopped", err)
	}
}

func initDatabase() *sqlx.DB {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "55432",
		Username: viper.GetString("POSTGRES_USER"),
		DBName:   viper.GetString("POSTGRES_DB"),
		SSLMode:  viper.GetString("POSTGRES_SSL_MODE"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	return db
}
