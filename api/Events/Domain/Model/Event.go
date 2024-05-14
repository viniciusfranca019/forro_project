package Event

import (
	Location "forro_project/api/Location/Domain/Model"
	"forro_project/packages/Util"
	"time"
)

type Event struct {
	id          string
	title       string
	date        time.Time
	description string
	link        string
	city        *Location.City
}

func NewEvent(city Location.City, date time.Time, description, title, link string) *Event {
	return &Event{
		id:          Util.GenarateUUid(),
		title:       title,
		date:        date,
		description: description,
		link:        link,
		city:        &city,
	}
}

func (e *Event) Title() string {
	return e.title
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

func (e *Event) City() *Location.City {
	return e.city
}
