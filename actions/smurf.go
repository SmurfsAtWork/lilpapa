package actions

import (
	"github.com/SmurfsAtWork/lilpapa/app/models"
)

type Smurf struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type CreateSmurfParams struct {
	ActionContext
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateSmurfPayload struct {
	Id     uint   `json:"id"`
	NanoId string `json:"nano_id"`
}

func (a *Actions) CreateSmurf(params CreateSmurfParams) (CreateSmurfPayload, error) {
	newSmurf, err := a.app.CreateSmurf(models.Smurf{
		Name:     params.Name,
		Password: params.Password,
	})
	if err != nil {
		return CreateSmurfPayload{}, err
	}

	return CreateSmurfPayload{
		Id:     newSmurf.Id,
		NanoId: newSmurf.NanoId,
	}, nil
}

type UpdateSmurfPasswordParams struct {
	ActionContext
	SmurfId     uint   `json:"smurf_id"`
	NewPassword string `json:"new_password"`
}

type UpdateSmurfPasswordPayload struct {
}

func (a *Actions) UpdateSmurfPassword(params UpdateSmurfPasswordParams) (UpdateSmurfPasswordPayload, error) {
	err := a.app.UpdateSmurfPassword(params.SmurfId, params.NewPassword)
	if err != nil {
		return UpdateSmurfPasswordPayload{}, err
	}

	return UpdateSmurfPasswordPayload{}, nil
}

type DeleteSmurfParams struct {
	ActionContext
	SmurfId uint `json:"smurf_id"`
}

type DeleteSmurfPayload struct {
}

func (a *Actions) DeleteSmurf(params DeleteSmurfParams) (DeleteSmurfPayload, error) {
	err := a.app.DeleteSmurf(params.SmurfId)
	if err != nil {
		return DeleteSmurfPayload{}, err
	}

	return DeleteSmurfPayload{}, nil
}
