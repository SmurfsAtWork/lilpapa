package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/SmurfsAtWork/lilpapa/actions"
	"github.com/SmurfsAtWork/lilpapa/app"
	"github.com/SmurfsAtWork/lilpapa/evy"
	"github.com/SmurfsAtWork/lilpapa/handlers/events"
	"github.com/SmurfsAtWork/lilpapa/log"
	"github.com/SmurfsAtWork/lilpapa/sqlite"
)

const (
	eventsBatchItems     = 25
	fetchWaitTimeSeconds = 5
)

var (
	evyRepo             evy.Repository
	eventsHandlers      *events.EventHandlers
	executingEventsRepo = &executingEvents{
		currentEvents: map[string]struct{}{},
	}
)

func init() {
	sqlite3Repo, err := sqlite.New()
	if err != nil {
		log.Fatalln(err)
	}

	evyRepo = sqlite3Repo
	app := app.New(sqlite3Repo)
	eventhub := evy.New(sqlite3Repo)

	usecases := actions.New(
		app,
		nil,
		eventhub,
		nil,
		nil,
	)

	eventsHandlers = events.New(usecases)
}

func executeEvents(events []evy.EventPayload) error {
	wg := sync.WaitGroup{}
	wg.Add(len(events))

	for _, e := range events {
		log.Warningln("handling event", e.Topic)

		if executingEventsRepo.Exists(e) {
			continue
		}
		executingEventsRepo.Add(e)

		switch e.Topic {
		case "song-played":
			var body any
			err := json.Unmarshal([]byte(e.Body), &body)
			if err != nil {
				log.Errorf("failed unmarshalling event's json: %v\n", err)
				continue
			}

			go func() {
				err := errors.Join(
				// handlers.HandleAddSongToQueue(body),
				)
				if err != nil {
					log.Errorln("song-played", err)
				}

				wg.Done()
			}()
		}
	}

	wg.Wait()

	for _, e := range events {
		executingEventsRepo.Delete(e)
		err := evyRepo.DeleteEvent(e.Id)
		if err != nil {
			log.Errorf("Failed deleting event: %+v, error: %v\n", e, err)
			return err
		}
	}

	return nil
}

func fetchAndExecuteEventsAsync() {
	timer := time.NewTicker(time.Second * fetchWaitTimeSeconds)
	for range timer.C {
		events, err := evyRepo.GetEventsBatch(eventsBatchItems)
		if err != nil {
			continue
		}

		err = executeEvents(events)
		if err != nil {
			log.Errorln("Failed executing events batch", err)
		}
	}
}

func handleEventEmitted(w http.ResponseWriter, r *http.Request) {
	var event evy.EventPayload
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Errorln("Failed marshalling event", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if executingEventsRepo.Exists(event) {
		w.WriteHeader(http.StatusOK)
		return
	}

	err = evyRepo.CreateEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorln(err)
		return
	}
	executingEventsRepo.Add(event)
}
