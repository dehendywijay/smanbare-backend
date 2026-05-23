package cmd

import (
	"gin-app/cmd"
	"net/http"
)

var r = cmd.SetupRouter()

func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}