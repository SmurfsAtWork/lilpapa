package models

import "time"

type SmurfStatus struct {
	Id       uint `gorm:"primaryKey;autoIncrement"`
	SmurfId  uint `gorm:"not null;index"`
	Smurf    Smurf
	CpuUsage int64
	RamTotal int64
	RamUsed  int64

	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

func (SmurfStatus) TableName() string {
	return "smurf_status"
}
