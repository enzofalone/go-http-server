package handlers

import "github.com/enzofalone/go-http-server/app/http"

func HandleHealth(req http.Request) string {
	return http.RespondWithStatus("pong", http.StatusOK)
}
