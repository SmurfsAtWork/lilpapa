package evy

import (
	"encoding/json"

	"github.com/SmurfsAtWork/lilpapa/evy/events"
)

type EventPayload struct {
	Id uint `gorm:"primaryKey;autoIncrement" json:"-"`

	Topic string `json:"topic"`
	Body  string `json:"body"`
}

func (EventPayload) TableName() string {
	return "event_payloads"
}

type Evy struct {
	repo Repository
}

func New(repo Repository) *Evy {
	return &Evy{
		repo: repo,
	}
}

func (e *Evy) Publish(event events.Event) error {
	eventBody, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fullEvent := EventPayload{
		Topic: event.Topic(),
		Body:  string(eventBody),
	}

	return e.repo.CreateEvent(fullEvent)
}
