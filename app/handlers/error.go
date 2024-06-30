package handlers

import "github.com/enzofalone/go-http-server/app/http"

func NotFound(req http.Request) string {
	return http.RespondWithStatus("", http.StatusNotFound)
}

func GenericError(req http.Request) string {
	return http.RespondWithStatus("", http.StatusInternalServerError)
}

func BadRequest(req http.Request) string {
	return http.RespondWithStatus("", http.StatusBadRequest)
}
