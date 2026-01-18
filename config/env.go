package config

import (
	"os"

	"github.com/SmurfsAtWork/lilpapa/log"
)

var (
	values = envValues{}
)

func init() {
	values = envValues{
		Port:      getEnv("PORT"),
		GoEnv:     GoEnv(getEnv("GO_ENV")),
		JwtSecret: getEnv("JWT_SECRET"),
		BlobsDir:  getEnv("BLOBS_DIR"),
		AdminCredentials: struct {
			Username string
			Password string
		}{
			Username: getEnv("ADMIN_USERNAME"),
			Password: getEnv("ADMIN_PASSWORD"),
		},
		Sqlite3FilePath: getEnv("SQLITE3_FILE_PATH"),
	}
}

type GoEnv string

const (
	GoEnvProd GoEnv = "prod"
	GoEnvBeta GoEnv = "beta"
	GoEnvDev  GoEnv = "dev"
	GoEnvTest GoEnv = "test"
)

type envValues struct {
	Port             string
	GoEnv            GoEnv
	JwtSecret        string
	BlobsDir         string
	AdminCredentials struct {
		Username string
		Password string
	}
	Sqlite3FilePath string
}

// Env returns the thing's config values :)
func Env() envValues {
	return values
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatalln("The \"" + key + "\" variable is missing.")
	}
	return value
}
