package sqlite

import (
	"time"

	"github.com/SmurfsAtWork/lilpapa/app/models"
	"github.com/SmurfsAtWork/lilpapa/config"
	"github.com/SmurfsAtWork/lilpapa/evy"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/schema"
)

var migratableModels = []schema.Tabler{
	new(models.User),
	new(models.Program),
	new(models.Script),
	new(models.Runnable),
	new(models.Smurf),
	new(models.SmurfConfig),
	new(models.SmurfCommand),
	new(models.SmurfLog),
	new(models.SmurfStatus),
	new(evy.EventPayload),
}

func Migrate() error {
	dbConn, err := dbConnector()
	if err != nil {
		return err
	}

	for _, table := range migratableModels {
		err = dbConn.Debug().AutoMigrate(table)
		if err != nil {
			return err
		}
	}

	err = dbConn.Exec("PRAGMA encoding=\"UTF-8\";").Error
	if err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(config.Env().AdminCredentials.Password), bcrypt.DefaultCost)

	superMechman := models.User{
		Username:  config.Env().AdminCredentials.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	_ = dbConn.Create(&superMechman)

	return nil
}
