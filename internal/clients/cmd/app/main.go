package main

import (
	"flag"
	"log"

	"github.com/gavrl/sleep-go-bot/internal/clients/repository"
	"github.com/gavrl/sleep-go-bot/internal/clients/telegram"
	"github.com/gavrl/sleep-go-bot/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// todo перенести в env
const (
	tgBotHost = "api.telegram.org"
)

func init() {
	util.InitConfig()
	util.InitLogger()
	initDatabase()
}

func main() {
	tgClient = telegram.New(tgBotHost, viper.GetString("TELEGRAM_BOT_TOKEN"))

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

// todo перенести в env
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

func initDatabase() {
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
}
