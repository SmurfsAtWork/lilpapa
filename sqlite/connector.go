package sqlite

import (
	"os"
	"path"

	"github.com/SmurfsAtWork/lilpapa/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB = nil

func dbConnector() (*gorm.DB, error) {
	sqlite3DirPath := path.Dir(config.Env().Sqlite3FilePath)
	sqlite3FilePath := config.Env().Sqlite3FilePath

	if _, err := os.Stat(sqlite3DirPath); !os.IsExist(err) {
		_ = os.Mkdir(sqlite3DirPath, 0755)
	}
	if _, err := os.Stat(sqlite3FilePath); instance == nil || err != nil {
		var err error
		instance, err = gorm.Open(sqlite.Open(sqlite3FilePath), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	return instance, nil
}
