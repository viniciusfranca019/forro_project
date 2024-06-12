package Location

import (
	"forro_project/api/Base/Domain/Model"
)

type Country struct {
	Model.Model
	Name         string  `gorm:"type:varchar(100);not null"`
	Abbreviation string  `gorm:"type:varchar(3);not null"`
	States       []State `gorm:"foreignKey:CountryID"`
}

func NewCountry(name string, abbreviation string) *Country {
	return &Country{
		Name:         name,
		Abbreviation: abbreviation,
	}
}
