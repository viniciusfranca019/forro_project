package Location

import (
	"forro_project/packages/Util"
	"github.com/google/uuid"
)

type State struct {
	id           uuid.UUID
	country      *Country
	name         string
	abbreviation string
}

func NewState(country *Country, name string, abbreviation string) *State {
	return &State{
		id:           Util.GenarateUUid(),
		country:      country,
		name:         name,
		abbreviation: abbreviation,
	}
}

func (s *State) ID() uuid.UUID {
	return s.id
}

func (s *State) Name() string {
	return s.name
}

func (s *State) Country() *Country {
	return s.country
}

func (s *State) Abbreviation() string {
	return s.abbreviation
}
