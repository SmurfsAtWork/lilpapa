package app

import "github.com/SmurfsAtWork/lilpapa/app/models"

func (a *App) GetSmurfById(id uint) (models.Smurf, error) {
	return a.repo.GetSmurf(id)
}
