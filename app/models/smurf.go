package models

import "time"

type Smurf struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Smurf) TableName() string {
	return "smurfs"
}
