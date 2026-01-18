package models

import "time"

type SmurfCommand struct {
	Id                    uint `gorm:"primaryKey;autoIncrement"`
	SmurfId               uint `gorm:"not null;index"`
	Smurf                 Smurf
	RunnableId            uint
	Runnable              Runnable
	RunnableHealthCheckId uint
	RunnableHealthCheck   Runnable

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SmurfCommand) TableName() string {
	return "smurf_commands"
}
