package Location

import (
	"forro_project/api/Base/Domain/Model"
)

type State struct {
	Model.Model
	Name         string   `gorm:"type:varchar(100);not null"`
	Abbreviation string   `gorm:"type:varchar(3);not null"`
	CountryID    string   `gorm:"type:uuid;not null"`
	Country      *Country `gorm:"foreignKey:CountryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cities       []City   `gorm:"foreignKey:StateID"`
}

func NewState(country *Country, name string, abbreviation string) *State {
	return &State{
		Country:      country,
		Name:         name,
		Abbreviation: abbreviation,
	}
}
