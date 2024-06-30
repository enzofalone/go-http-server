package handlers

import (
	"github.com/enzofalone/go-http-server/app/http"
)

func HandleEcho(req http.Request) string {
	str := req.Params["str"]
	id := req.QueryParams["id"]

	result := "str: " + str + " id: " + id + "\n"
	return http.RespondWithStatus(result, http.StatusOK)
}
