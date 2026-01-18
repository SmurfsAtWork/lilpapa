package apis

import (
	"encoding/json"
	"net/http"

	"github.com/SmurfsAtWork/lilpapa/actions"
	"github.com/SmurfsAtWork/lilpapa/log"
)

type smurfLoginApi struct {
	usecases *actions.Actions
}

func NewSmurfLoginApi(usecases *actions.Actions) *smurfLoginApi {
	return &smurfLoginApi{
		usecases: usecases,
	}
}

func (e *smurfLoginApi) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var reqBody actions.LoginSmurfParams
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handleErrorResponse(w, err)
		return
	}

	payload, err := e.usecases.LoginSmurf(reqBody)
	if err != nil {
		log.Errorf("[EMAIL LOGIN API]: Failed to login smurf: %+v, error: %s\n", reqBody, err.Error())
		handleErrorResponse(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}
