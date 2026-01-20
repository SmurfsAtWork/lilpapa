package app

import (
	"github.com/SmurfsAtWork/lilpapa/app/models"
)

func (a *App) GetSmurfById(id uint) (models.Smurf, error) {
	return a.repo.GetSmurf(id)
}

func (a *App) GetSmurfByNanoId(nanoId string) (models.Smurf, error) {
	return a.repo.GetSmurfByNanoId(nanoId)
}

func (a *App) CreateSmurf(smurf models.Smurf) (models.Smurf, error) {
	return a.repo.CreateSmurf(smurf)
}

func (a *App) UpdateSmurfPassword(id uint, newPassword string) error {
	return a.repo.UpdateSmurfPassword(id, newPassword)
}

func (a *App) DeleteSmurf(id uint) error {
	return a.repo.DeleteSmurf(id)
}
