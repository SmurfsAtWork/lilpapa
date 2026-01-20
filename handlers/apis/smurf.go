package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SmurfsAtWork/lilpapa/actions"
)

type smurfApi struct {
	usecases *actions.Actions
}

func NewSmurfApi(usecases *actions.Actions) *smurfApi {
	return &smurfApi{
		usecases: usecases,
	}
}

func (u *smurfApi) HandleCreateSmurf(w http.ResponseWriter, r *http.Request) {
	ctx, err := parseContext(r.Context())
	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	var params actions.CreateSmurfParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		handleErrorResponse(w, err)
		return
	}
	params.ActionContext = ctx

	payload, err := u.usecases.CreateSmurf(params)
	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func (u *smurfApi) HandleUpdateSmurfPassword(w http.ResponseWriter, r *http.Request) {
	ctx, err := parseContext(r.Context())
	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	smurfId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		handleErrorResponse(w, &ErrBadRequest{
			FieldName: "id",
		})
		return
	}

	var params actions.UpdateSmurfPasswordParams
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		handleErrorResponse(w, err)
		return
	}
	params.ActionContext = ctx
	params.SmurfId = uint(smurfId)

	payload, err := u.usecases.UpdateSmurfPassword(params)
	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func (u *smurfApi) HandleDeleteSmurf(w http.ResponseWriter, r *http.Request) {
	ctx, err := parseContext(r.Context())
	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	smurfId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		handleErrorResponse(w, &ErrBadRequest{
			FieldName: "id",
		})
		return
	}

	params := actions.DeleteSmurfParams{
		ActionContext: ctx,
		SmurfId:       uint(smurfId),
	}

	payload, err := u.usecases.DeleteSmurf(params)
	if err != nil {
		handleErrorResponse(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}
