package telegram

import "github.com/gavrl/sleep-go-bot/internal/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}

func NewProcessor(client telegram.Client) {

}
