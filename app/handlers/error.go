package handlers

import "github.com/enzofalone/go-http-server/app/http"

func NotFound() string {
	return http.RespondWithStatus("", http.StatusNotFound)
}

func GenericError() string {
	return http.RespondWithStatus("", http.StatusInternalServerError)
}

func BadRequest() string {
	return http.RespondWithStatus("", http.StatusBadRequest)
}
