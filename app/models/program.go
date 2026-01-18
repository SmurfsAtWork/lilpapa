package models

import "time"

type ProgramArchitecture string

const (
	ProgramArchitectureAmd64   ProgramArchitecture = "amd64"
	ProgramArchitectureAarch64 ProgramArchitecture = "aarch64"
	ProgramArchitectureArm32   ProgramArchitecture = "arm32"
	ProgramArchitectureI386    ProgramArchitecture = "i386"
)

type ProgramOperatingSystem string

const (
	ProgramOperatingSystemLinux   ProgramOperatingSystem = "linux"
	ProgramOperatingSystemFreeBSD ProgramOperatingSystem = "freebsd"
	ProgramOperatingSystemOpenBSD ProgramOperatingSystem = "openbsd"
	ProgramOperatingSystemWindows ProgramOperatingSystem = "windows"
)

type Program struct {
	Id              uint                   `gorm:"primaryKey;autoIncrement"`
	PublicId        string                 `gorm:"index;unique;not null"`
	Name            string                 `gorm:"not null"`
	Architecture    ProgramArchitecture    `gorm:"not null;index"`
	OperatingSystem ProgramOperatingSystem `gorm:"not null;index"`
	SizeBytes       uint64                 `gorm:"not null"`
	DownloadPath    string                 `gorm:"not null"`
	BuildScript     string                 `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Program) TableName() string {
	return "programs"
}
