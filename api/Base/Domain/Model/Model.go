package Model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Identity struct {
	ID string `gorm:"type:uuid;primaryKey" json:"id"`
}

type TimeTrace struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (i *Identity) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()

	if err != nil {
		panic("error on new id")
	}

	i.ID = id.String()
	return
}
