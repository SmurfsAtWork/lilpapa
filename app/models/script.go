package models

import "time"

type Script struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	ExecString string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Script) TableName() string {
	return "scripts"
}
