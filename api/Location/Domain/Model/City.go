package Location

import (
	"forro_project/api/Base/Domain/Model"
)

type City struct {
	Model.Identity
	Name    string `gorm:"type:varchar(100);not null"`
	StateID string `gorm:"type:uuid;not null"`
	State   *State `gorm:"foreignKey:StateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Model.TimeTrace
}

func NewCity(state *State, name string) *City {
	return &City{
		Name:  name,
		State: state,
	}
}
