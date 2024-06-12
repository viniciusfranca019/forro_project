package Location

import (
	"forro_project/api/Base/Domain/Model"
)

type City struct {
	Model.Model
	Name         string `gorm:"type:varchar(100);not null"`
	StateID      string `gorm:"type:uuid;not null"`
	State        *State `gorm:"foreignKey:StateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CityIbgeCode string `gorm:"type:varchar(25);not null"`
}

func NewCity(state *State, name string, cityIbgeCode string) *City {
	return &City{
		Name:         name,
		CityIbgeCode: cityIbgeCode,
		State:        state,
	}
}
