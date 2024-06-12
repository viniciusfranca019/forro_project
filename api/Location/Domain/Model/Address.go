package Location

import (
	"forro_project/packages/Util"
	"github.com/google/uuid"
)

type Address struct {
	id        uuid.UUID
	city      *City
	reference string
}

func newAddress(city *City, reference string) *Address {
	return &Address{
		id:        Util.GenarateUUid(),
		reference: reference,
		city:      city,
	}
}

func (address *Address) ID() uuid.UUID {
	return address.id
}

func (address *Address) Reference() string {
	return address.reference
}

func (address *Address) City() *City {
	return address.city
}
