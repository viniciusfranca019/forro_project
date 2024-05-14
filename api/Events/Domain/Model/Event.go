package Model

import (
	"forro_project/packages/Util"
	"time"
)

type Event struct {
	id          string
	date        time.Time
	description string
	location    string
	link        string
	eventType   string
}

func NewEvent(date time.Time, description, location, link, eventType string) *Event {

	return &Event{
		id:          Util.GenarateUUid(),
		date:        date,
		description: description,
		location:    location,
		link:        link,
		eventType:   eventType,
	}
}

func (e *Event) ID() string {
	return e.id
}

func (e *Event) Date() time.Time {
	return e.date
}

func (e *Event) Description() string {
	return e.description
}

func (e *Event) Location() string {
	return e.location
}

func (e *Event) Link() string {
	return e.link
}
