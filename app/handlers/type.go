package handlers

import "github.com/enzofalone/go-http-server/app/http"

type Handler func(http.Request) string
