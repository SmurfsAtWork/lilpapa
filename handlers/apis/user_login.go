package apis

import (
	"encoding/json"
	"net/http"

	"github.com/SmurfsAtWork/lilpapa/actions"
	"github.com/SmurfsAtWork/lilpapa/log"
)

type adminLoginApi struct {
	usecases *actions.Actions
}

func NewAdminLoginApi(usecases *actions.Actions) *adminLoginApi {
	return &adminLoginApi{
		usecases: usecases,
	}
}

func (a *adminLoginApi) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var reqBody actions.LoginUserParams
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handleErrorResponse(w, err)
		return
	}

	payload, err := a.usecases.LoginUser(reqBody)
	if err != nil {
		log.Errorf("[EMAIL LOGIN API]: Failed to login admin: %+v, error: %s\n", reqBody, err.Error())
		handleErrorResponse(w, err)
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}
