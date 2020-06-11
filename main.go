package main

import (
	h "FullGas/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main(){
	err := os.Setenv("DATABASE_URL", "postgres://postgres:12345@localhost/fgmotoru")
	if err != nil {
		log.Fatal("ListenAndServe Ошибка установки ENV : ", err)
	}
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

	fmt.Println("Hello! It's work!")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
