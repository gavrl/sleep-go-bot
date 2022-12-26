package telegram

import (
	"github.com/gavrl/sleep-go-bot/internal/dto"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type command struct {
	text     string
	chatID   int
	username string
}

const (
	HelpCmd  = "/help"
	HelloCmd = "/start"
)

func (p *Processor) doCmd(cmd *command) error {
	text := strings.TrimSpace(cmd.text)

	logrus.Printf("got new command '%s' from '%s'", text, cmd.username)

	if isAddRateCmd(text) {
		return p.addRate(cmd)
	}

	switch text {
	case HelpCmd:
		return p.sendHelp(cmd.chatID)
	case HelloCmd:
		return p.sendHello(cmd.chatID)
	default:
		return p.tg.SendMessage(cmd.chatID, msgUnknownCommand)
	}

}

func (p Processor) addRate(cmd *command) error {
	rate, _ := strconv.Atoi(cmd.text)
	_, err := p.service.Save(&dto.SaveSleepRateDto{
		UserName: cmd.username,
		Rate:     rate,
		Time:     time.Now(),
	})
	if err != nil {
		return p.tg.SendMessage(cmd.chatID, msgSaveRateErr)
	}

	return p.tg.SendMessage(cmd.chatID, msgSavedSleepRate)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func isAddRateCmd(text string) bool {
	num, err := strconv.Atoi(text)
	if err != nil {
		return false
	}
	if num < 1 || num > 100 {
		return false
	}

	return true
}
