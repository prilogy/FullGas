package main

import (
	help "FullGas/helpers"
	h "FullGas/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main(){
	r := mux.NewRouter()

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	server := os.Getenv("SERVER")
	os.Setenv("SERVER", "")


	r.HandleFunc("/", h.Index)
	r.HandleFunc("/tires", h.Tires)
	r.HandleFunc("/pads", h.Pads)
	r.HandleFunc("/chains", h.Chains)
	r.HandleFunc("/tires/cub/{cub}/type/{type}", h.Radius)
	r.HandleFunc("/tires/cub/{cub}/type/{type}/rFront/{rFront}/rBack/{rBack}", h.TiresPrice)
	r.HandleFunc("/tires/cub/{cub}/type/{type}/rFront/{rFront}", h.TiresFPrice)
	r.HandleFunc("/tires/cub/{cub}/type/{type}/rBack/{rBack}/spike/{spike}", h.TiresBPrice)
	r.HandleFunc("/product/{product}/id/{id}", h.CreateOrder)
	//r.HandleFunc("/product/tires/cub/{cub}/rFront/{rFront}/rBack/{rBack}/type/{type}/spike/{spike}/price/{price}", h.CreateOrderTires)
	r.HandleFunc("/orderId/{orderId}", h.SendEmail)

	r.PathPrefix("/src/css").Handler(http.StripPrefix("/src/css",
		http.FileServer(http.Dir(server + "templates/src/css"))))
	r.PathPrefix("/src/scripts").Handler(http.StripPrefix("/src/scripts",
		http.FileServer(http.Dir(server + "templates/src/scripts"))))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir(server + "templates/static"))))
	r.PathPrefix("/templates").Handler(http.StripPrefix("/templates",
		http.FileServer(http.Dir(server + "templates"))))

	fmt.Println("Hello! It's work!")
	http.Handle("/", r)
	err := http.ListenAndServe(":8000", nil)

	help.ErrDefaultDetect(err, "Запуск сервера")
}
