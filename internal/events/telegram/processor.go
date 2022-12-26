package telegram

import (
	"errors"
	"github.com/gavrl/sleep-go-bot/internal/clients/telegram"
	"github.com/gavrl/sleep-go-bot/internal/events"
	"github.com/gavrl/sleep-go-bot/internal/service"
	"github.com/gavrl/sleep-go-bot/pkg/e"
)

type Processor struct {
	tg      *telegram.Client
	service service.Service
	offset  int
}

type Meta struct {
	ChatID   int
	Username string
}

var ErrUnknownEventType = errors.New("unknown event type")
var ErrUnknownMetaType = errors.New("unknown meta type")

func NewProcessor(tg *telegram.Client, service service.Service) *Processor {
	return &Processor{tg: tg, service: service}
}

func (p Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))
	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

	if err := p.doCmd(&command{
		text:     event.Text,
		chatID:   meta.ChatID,
		username: meta.Username,
	}); err != nil {
		return e.Wrap("can't process message", err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMetaType)
	}
	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}
