package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func startWebService() {
	r := mux.NewRouter()

	http.ListenAndServe(cfg.Listen, r)
}
