package app

import (
	"time"

	"github.com/SmurfsAtWork/lilpapa/app/models"
)

type Repository interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id uint) (models.User, error)
	GetUserByUsername(username string) (models.User, error)

	CreateSmurf(smurf models.Smurf) (models.Smurf, error)
	GetSmurf(id uint) (models.Smurf, error)
	GetSmurfByNanoId(nanoId string) (models.Smurf, error)
	UpdateSmurfPassword(id uint, newPassword string) error
	DeleteSmurf(id uint) error

	UpsertSmurfConfig(smurfId uint, smurfConfig models.SmurfConfig) (models.SmurfConfig, error)
	GetSmurfConfig(smurfId uint) (models.SmurfConfig, error)

	CreateSmurfCommand(smurfCommand models.SmurfCommand) (models.SmurfCommand, error)
	GetSmurfCommand(smurfId uint) (models.SmurfCommand, error)
	GetSmurfCommands(smurfId uint) ([]models.SmurfCommand, error)
	DeleteSmurfCommand(smurfId, commandId uint) error
	DeleteSmurfCommands(smurfId uint) error

	CreateSmurfLog(smurfLog models.SmurfLog) (models.SmurfLog, error)
	GetSmurfLogs(smurfId uint, since time.Time) ([]models.SmurfLog, error)

	CreateSmurfStat(smurfStat models.SmurfStatus) (models.SmurfStatus, error)
	GetSmurfStats(smurfId uint, since time.Time) ([]models.SmurfStatus, error)

	CreateProgram(program models.Program) (models.Program, error)

	CreateScript(script models.Script) (models.Script, error)

	CreateRunnable(runnable models.Runnable) (models.Runnable, error)
}
