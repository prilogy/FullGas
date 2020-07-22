package http

import (
	help "FullGas/helpers"
	"FullGas/models"
	"context"
	"github.com/jackc/pgx"
	"html/template"
	"net/http"
	"os"
)

func Tires(w http.ResponseWriter, r *http.Request)  {
	server := os.Getenv("SERVER")
	data := models.Tires()

	//Догрузка из БД
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	help.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select cub, radius_front, radius_back from _tires")
	defer rows.Close()
	help.ErrDefaultDetect(err, "QueryRow")

	iData := make([]int, 0, 5)
	var temp, radiusF, radiusB int
	data.TiresInside.RadiusInside = make(map[int][]int)

	for rows.Next(){
		err = rows.Scan(&temp, &radiusF, &radiusB)

		if radiusF != 0{
			data.TiresInside.RadiusInside[temp] = append(data.TiresInside.RadiusInside[temp], radiusF)
		}
		if radiusB != 0{
			data.TiresInside.RadiusInside[temp] = append(data.TiresInside.RadiusInside[temp], radiusB)
		}

		iData = append(iData, temp)
		help.ErrDefaultDetect(err, "Row Scan")
	}

	for k, v := range data.TiresInside.RadiusInside{
		data.TiresInside.RadiusInside[k] = help.Unique(v)
	}

	data.TiresInside.Cub = append(data.TiresInside.Cub, help.Unique(iData)...)

	templates, err := template.ParseFiles(
		server + "templates/header.tmpl",
		server + "templates/tires.tmpl",
		server + "templates/footer.tmpl")

	tmpl := templates.Lookup("tires.tmpl")
	err = tmpl.ExecuteTemplate(w, "tires", data)

	help.ErrCatch(err, "Перевод в шаблон")
}