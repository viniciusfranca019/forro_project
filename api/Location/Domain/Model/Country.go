package Location

import (
	"forro_project/api/Base/Domain/Model"
)

type Country struct {
	Model.Identity
	Name         string  `gorm:"type:varchar(100);not null"`
	Abbreviation string  `gorm:"type:varchar(3);not null"`
	States       []State `gorm:"foreignKey:CountryID"`
	Model.TimeTrace
}

func NewCountry(name string, abbreviation string) *Country {
	return &Country{
		Name:         name,
		Abbreviation: abbreviation,
	}
}
