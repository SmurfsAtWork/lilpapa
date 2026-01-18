package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SmurfsAtWork/lilpapa/app"
)

type ErrorResponse struct {
	ErrorId   string         `json:"error_id"`
	ExtraData map[string]any `json:"extra_data,omitempty"`
}

func HandleErrorResponse(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	if papaError, ok := err.(app.PapaError); ok {
		if papaError.ExposeToClients() {
			w.WriteHeader(papaError.ClientStatusCode())
			_ = json.NewEncoder(w).Encode(ErrorResponse{
				ErrorId:   strings.ToLower(papaError.Error()),
				ExtraData: papaError.ExtraData(),
			})
			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		ErrorId: "internal-server-error",
	})
}
