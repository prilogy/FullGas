package http

import (
	er "../helpers/errCatch"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"net/http"
	"os"
)

func Radius(w http.ResponseWriter, r *http.Request)  {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	er.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	vars := mux.Vars(r)
	type Tire struct {
		Id			int
		RadiusFront	int
		RadiusBack	int
		Spike		int
	}

	data := Tire{}
	tires := []Tire{}

	rows, err := conn.Query(context.Background(),
		"select id,radius_front,radius_back,spike from _tires WHERE kit_unit=$1 AND cub=$2",
		vars["type"], vars["cub"])
	er.ErrDefaultDetect(err, "QueryRow")
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.RadiusFront, &data.RadiusBack, &data.Spike)
		tires = append(tires, data)
	}
	er.ErrDefaultDetect(err, "Row Scan")
	w.WriteHeader(http.StatusOK)

	Jdata, _ := json.Marshal(tires)
	fmt.Fprintf(w, "%s", Jdata)
}