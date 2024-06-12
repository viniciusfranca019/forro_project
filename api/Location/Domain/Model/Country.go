package Location

import (
	"forro_project/packages/Util"
	"github.com/google/uuid"
)

type Country struct {
	id           uuid.UUID
	name         string
	abbreviation string
}

func NewCountry(name string, abbreviation string) *Country {
	return &Country{
		id:           Util.GenarateUUid(),
		name:         name,
		abbreviation: abbreviation,
	}
}

func (c *Country) ID() uuid.UUID {
	return c.id
}

func (c *Country) Name() string {
	return c.name
}

func (c *Country) Abbreviation() string {
	return c.abbreviation
}
