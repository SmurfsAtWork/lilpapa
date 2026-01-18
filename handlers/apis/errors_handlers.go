package apis

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SmurfsAtWork/lilpapa/app"
	"github.com/SmurfsAtWork/lilpapa/log"
)

type errorResponse struct {
	ErrorId   string         `json:"error_id"`
	ExtraData map[string]any `json:"extra_data,omitempty"`
}

func handleErrorResponse(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	log.Errorf("error happened in api, %v\n", err)

	if papaError, ok := err.(app.PapaError); ok {
		if papaError.ExposeToClients() {
			w.WriteHeader(papaError.ClientStatusCode())
			_ = json.NewEncoder(w).Encode(errorResponse{
				ErrorId:   strings.ToLower(papaError.Error()),
				ExtraData: papaError.ExtraData(),
			})
			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(errorResponse{
		ErrorId: "internal-server-error",
	})
}
