package apis

import (
	"net/http"

	"github.com/SmurfsAtWork/lilpapa/actions"
)

type smurfSmurfMeApi struct {
	usecases *actions.Actions
}

func NewSmurfMeApi(usecases *actions.Actions) *smurfSmurfMeApi {
	return &smurfSmurfMeApi{
		usecases: usecases,
	}
}

func (u *smurfSmurfMeApi) HandleAuthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := parseContext(r.Context())
	if err != nil {
		handleErrorResponse(w, err)
		return
	}
}

func (m *smurfSmurfMeApi) HandleLogout(w http.ResponseWriter, r *http.Request) {
	sessionToken, ok := r.Header["Authorization"]
	if !ok {
		return
	}
	_ = m.usecases.InvalidateAuthenticatedSmurf(sessionToken[0])
}
