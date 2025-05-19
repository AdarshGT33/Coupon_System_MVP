package models

import (
	"time"

	"github.com/google/uuid"
)

type Medicine struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	MedicineName  string
	MedicineBrand string
	Category      string
	Price         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
