package Model

import (
	"forro_project/packages/Util"
	"time"
)

type Event struct {
	id          string
	date        time.Time
	description string
	link        string
}

func NewEvent(date time.Time, description, location, link, eventType string) *Event {
	return &Event{
		id:          Util.GenarateUUid(),
		date:        date,
		description: description,
		link:        link,
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

func (e *Event) Link() string {
	return e.link
}
