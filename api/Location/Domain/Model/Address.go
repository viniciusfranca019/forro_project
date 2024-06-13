package Location

import (
	"forro_project/api/Base/Domain/Model"
)

type Address struct {
	Model.Identity
	Reference string `gorm:"type:varchar(255);not null"`
	CityID    string `gorm:"type:uuid;not null"`
	City      *City  `gorm:"foreignKey:CityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Model.TimeTrace
}

func newAddress(city *City, reference string) *Address {
	return &Address{
		Reference: reference,
		City:      city,
	}
}
