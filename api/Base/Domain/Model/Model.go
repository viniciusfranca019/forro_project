package Model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        string         `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV7()

	if err != nil {
		panic("error on new id")
	}

	m.ID = id.String()
	return
}
