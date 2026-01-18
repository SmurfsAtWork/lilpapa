package models

import "time"

type SmurfLog struct {
	Id      uint `gorm:"primaryKey;autoIncrement"`
	SmurfId uint `gorm:"not null;index"`
	Smurf   Smurf
	Text    string

	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

func (SmurfLog) TableName() string {
	return "smurf_logs"
}
