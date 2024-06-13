package handlers

import "github.com/enzofalone/go-http-server/app/http"

func HandleHealth() string {
	return http.RespondWithStatus("pong", http.StatusOK)
}
