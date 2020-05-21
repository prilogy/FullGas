package main

import (
	"./helpers/env"
	h "./http"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	env.SetEnv()
	r := mux.NewRouter()

	r.HandleFunc("/", h.Index)
	r.HandleFunc("/tires", h.Tires)
	r.HandleFunc("/pads", h.Pads)
	r.HandleFunc("/chains", h.Chains)
	r.HandleFunc("/tires/cub/{cub}/type/{type}", h.Radius)
	r.HandleFunc("/tires/cub/{cub}/type/{type}/rFront/{rFront}/rBack/{rBack}", h.TiresPrice)
	r.HandleFunc("/tires/cub/{cub}/type/{type}/rFront/{rFront}", h.TiresFPrice)
	r.HandleFunc("/tires/cub/{cub}/type/{type}/rBack/{rBack}/spike/{spike}", h.TiresBPrice)


	r.PathPrefix("/src/css").Handler(http.StripPrefix("/src/css",
		http.FileServer(http.Dir("templates/src/css"))))
	r.PathPrefix("/src/scripts").Handler(http.StripPrefix("/src/scripts",
		http.FileServer(http.Dir("templates/src/scripts"))))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("templates/static"))))

	http.Handle("/", r)
	http.ListenAndServe(":8181", nil)
}
