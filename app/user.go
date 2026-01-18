package app

import (
	"github.com/SmurfsAtWork/lilpapa/app/models"
)

func (a *App) GetUserByUsername(username string) (models.User, error) {
	return a.repo.GetUserByUsername(username)
}
