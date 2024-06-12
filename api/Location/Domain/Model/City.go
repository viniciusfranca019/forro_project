package Location

import (
	"forro_project/packages/Util"
	"github.com/google/uuid"
)

type City struct {
	id           uuid.UUID
	state        *State
	name         string
	cityIbgeCode string
}

func NewCity(state *State, name string, cityIbgeCode string) *City {
	return &City{
		id:           Util.GenarateUUid(),
		name:         name,
		cityIbgeCode: cityIbgeCode,
		state:        state,
	}
}

func (c *City) ID() uuid.UUID {
	return c.id
}

func (c *City) Name() string {
	return c.name
}

func (c *City) State() *State {
	return c.state
}

func (c *City) CityIbgeCode() string {
	return c.cityIbgeCode
}
