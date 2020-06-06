package http

import (
	"../helpers"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"net/http"
	"os"
	"strconv"
)

func TiresPrice(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	helpers.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	vars := mux.Vars(r)
	data := 0

	row := conn.QueryRow(context.Background(),
		"select price from _tires WHERE cub=$1 AND kit_unit=$2 AND radius_front=$3 AND radius_back=$4",
		vars["cub"], vars["type"], vars["rFront"], vars["rBack"])

	err = row.Scan(&data)
	helpers.ErrDefaultDetect(err, "QueryRow")

	Jdata, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s", Jdata)
}

func TiresFPrice(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	helpers.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	vars := mux.Vars(r)
	data := 0

	row := conn.QueryRow(context.Background(),
		"select price from _tires WHERE cub=$1 AND kit_unit=$2 AND radius_front=$3",
		vars["cub"], vars["type"], vars["rFront"])

	err = row.Scan(&data)
	helpers.ErrDefaultDetect(err, "QueryRow")

	w.WriteHeader(http.StatusOK)

	Jdata, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s", Jdata)
}

func TiresBPrice(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	helpers.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	vars := mux.Vars(r)
	data := 0

	spike, err := strconv.Atoi(vars["spike"])
	cub, err := strconv.Atoi(vars["cub"])
	_type, err := strconv.Atoi(vars["type"])
	rBack, err := strconv.Atoi(vars["rBack"])
	helpers.ErrDefaultDetect(err, "Первод в строку")

	row := conn.QueryRow(context.Background(),
		"select price from _tires WHERE cub=$1 AND kit_unit=$2 AND radius_back=$3 AND spike=$4",
		cub, _type, rBack, spike)

	err = row.Scan(&data)
	helpers.ErrDefaultDetect(err, "QueryRow")

	w.WriteHeader(http.StatusOK)

	Jdata, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s", Jdata)
}