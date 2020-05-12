package http

import (
	er "FullGas/helpers/errCatch"
	"FullGas/models"
	"context"
	"github.com/jackc/pgx"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func Pads(w http.ResponseWriter, r *http.Request)  {
	output := models.Pads()
	data := models.PadsProduct{}

	keys, ok := r.URL.Query()["page"]
	var key int
	if !ok || len(keys[0]) < 1 {
		key = 1
	}else {
		key, _ = strconv.Atoi(keys[0])
	}

	//Догрузка из БД
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	er.ErrDefaultDetect(err, "DataBase Connection")
	//Количество страниц
	var pages int

	marks := []string{"", "KTM", "Husqvarna", "Husaberg", "BMW", "Gas-Gas",
		"Sherco", "TM Racing", "YAMAHA", "Honda", "SUZUZKI", "Kawasaki"}

	defer conn.Close(context.Background())

	var rows pgx.Rows
	marksQuery, ok := r.URL.Query()["mark"]
	var marksQ string
	if !ok || len(marksQuery[0]) < 1 {
		marksQ = ""
		rows, err = conn.Query(context.Background(),
			"select * from _pads LIMIT 18 OFFSET $1", 18*(key-1))
	}else {
		marksQ = marksQuery[0]
		rows, err = conn.Query(context.Background(),
			"select * from _pads WHERE mark=$1 LIMIT 18 OFFSET $2", marksQ, 18*(key-1))
	}

	er.ErrDefaultDetect(err, "QueryRow")

	iData := []models.PadsProduct{}

	for rows.Next(){
		err = rows.Scan(&data.Id, &data.Mark, &data.Model, &data.Years, &data.Img, &data.Price)
		data.MarkName = marks[data.Mark]
		iData = append(iData, data)
		er.ErrDefaultDetect(err, "Row Scan")
	}
	defer rows.Close()

	if marksQ == ""{
		row := conn.QueryRow(context.Background(), "select COUNT(id) from _pads")
		err = row.Scan(&pages)
	}else{
		row := conn.QueryRow(context.Background(),
			"select COUNT(id) from _pads WHERE mark=$1", marksQ)
		err = row.Scan(&pages)
	}

	output.Pads = append(output.Pads, iData...)
	if marksQ != ""{
		output.PageMark = marksQ
	}
	temp := 1+pages/20
	output.Page = models.Pagination{
		Start: 1,
		End: temp,
		AllData: pages,
		Current: key,
		Link: "http://localhost:8181/pads/",
	}

	var fm = template.FuncMap{
		"Iter": func(count int) []int {
			var Items []int
			for i := 1; i <= count; i++ {
				Items = append(Items, i)
			}
			return Items
		},
	}

	tmpl := template.Must(template.New("pads").Funcs(fm).
		ParseFiles("templates/header.tmpl",
		"templates/pads.tmpl",
		"templates/footer.tmpl"))

	err = tmpl.ExecuteTemplate(w, "pads", output)

	er.ErrCatch(err, "Перевод в шаблон")
}