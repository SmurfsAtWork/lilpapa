package models

import "time"

type SmurfLogLevel string

const (
	SmurfLogLevelInfo    SmurfLogLevel = "INFO"
	SmurfLogLevelWarning SmurfLogLevel = "WARN"
	SmurfLogLevelError   SmurfLogLevel = "ERROR"
)

type SmurfConfig struct {
	Id                    uint `gorm:"primaryKey;autoIncrement"`
	SmurfId               uint `gorm:"not null;index"`
	Smurf                 Smurf
	ActiveCommandId       uint
	ActiveCommand         SmurfCommand
	CommandUpdateInterval time.Duration
	ConfigUpdateInterval  time.Duration
	LogsLevel             SmurfLogLevel
	LogsPushInterval      time.Duration
	StatsPushInterval     time.Duration

	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

func (SmurfConfig) TableName() string {
	return "smurf_configs"
}
