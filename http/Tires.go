package http

import (
	er "FullGas/helpers/errCatch"
	f "FullGas/helpers/function"
	"FullGas/models"
	"context"
	"github.com/jackc/pgx"
	"html/template"
	"net/http"
	"os"
)

func Tires(w http.ResponseWriter, r *http.Request)  {
	data := models.Tires()

	//Догрузка из БД
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	er.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select cub from _tires")
	defer rows.Close()
	er.ErrDefaultDetect(err, "QueryRow")

	iData := make([]int, 0, 5)
	var temp int

	for rows.Next(){
		err = rows.Scan(&temp)
		iData = append(iData, temp)
		er.ErrDefaultDetect(err, "Row Scan")
	}

	data.TiresInside.Cub = append(data.TiresInside.Cub, f.Unique(iData)...)

	templates, err := template.ParseFiles(
		"templates/header.tmpl",
		"templates/tires.tmpl",
		"templates/footer.tmpl")

	tmpl := templates.Lookup("tires.tmpl")
	err = tmpl.ExecuteTemplate(w, "tires", data)

	er.ErrCatch(err, "Перевод в шаблон")
}