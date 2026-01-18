package actions

import "github.com/SmurfsAtWork/lilpapa/evy/events"

// EventHub handles events publishing.
type EventHub interface {
	// Publish publishes the given event to the given eventhub implementation.
	Publish(event events.Event) error
}
