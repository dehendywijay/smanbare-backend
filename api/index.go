package handler

import (
	"gin-app/cmd/server"
	"net/http"
)

var r = server.SetupRouter()

func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}