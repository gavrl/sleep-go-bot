package event_consumer

import (
	"github.com/gavrl/sleep-go-bot/internal/events"
	"github.com/sirupsen/logrus"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func NewConsumer(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{fetcher: fetcher, processor: processor, batchSize: batchSize}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			logrus.Errorf("[ERR] consumer: %s", err)
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			logrus.Error(err)

			continue
		}
	}
}

// todo многопоточная обработка
func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		logrus.Infof("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			logrus.Errorf("can't handle event: %s", err)

			continue
		}
	}

	return nil
}
