package models

import "time"

type RunnableType string

const (
	RunnableTypeProgram RunnableType = "program"
	RunnableTypeScript  RunnableType = "script"
)

type Runnable struct {
	Id        uint         `gorm:"primaryKey;autoIncrement"`
	Type      RunnableType `gorm:"not null"`
	ScriptId  uint
	Script    Script
	ProgramId uint
	Program   Program
	Args      string

	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

func (Runnable) TableName() string {
	return "runnables"
}
