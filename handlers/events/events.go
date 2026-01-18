package events

import (
	"github.com/SmurfsAtWork/lilpapa/actions"
)

type EventHandlers struct {
	usecases *actions.Actions
}

func New(usecases *actions.Actions) *EventHandlers {
	return &EventHandlers{
		usecases: usecases,
	}
}
