package main

import (
	"github.com/gorilla/mux"
	"golex/helpers/env"
	"net/http"
	h "webKR/http"
)

func main(){
	env.SetEnv()
	r := mux.NewRouter()

	r.HandleFunc("/", h.Main)
	r.HandleFunc("/autoMan", h.AutoMan)

	r.PathPrefix("/src/css").Handler(http.StripPrefix("/src/css",
		http.FileServer(http.Dir("templates/src/css"))))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("templates/static"))))

	http.Handle("/", r)
	http.ListenAndServe(":8181", nil)
}
