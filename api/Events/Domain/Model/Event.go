package main

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	id          string    // ID of the event (UUID)
	date        time.Time // Date and time of the event
	description string    // Description of the event
	location    string    // Geographic coordinates of the event
	link        string    // Link related to the event
	eventType   string    // Type of the event
}

// NewEvent creates a new event with provided values
func NewEvent(date time.Time, description, location, link, eventType string) *Event {
	return &Event{
		id:          uuid.New().String(), // Generate UUID for ID
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
