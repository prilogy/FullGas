package http

import (
	"FullGas/helpers"
	"FullGas/models"
	"context"
	"github.com/jackc/pgx"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Stars(w http.ResponseWriter, r *http.Request)  {
	server := os.Getenv("SERVER")
	output := models.Stars()
	data := models.StarsProduct{}

	keys, ok := r.URL.Query()["page"]
	var key int
	if !ok || len(keys[0]) < 1 {
		key = 1
	}else {
		key, _ = strconv.Atoi(keys[0])
	}

	//Догрузка из БД
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	helpers.ErrDefaultDetect(err, "DataBase Connection")
	//Количество страниц
	var pages int

	//marks := []string{"", "Honda", "Kawasaki", "KTM", "Suzuki", "Husqvarna", "YAMAHA"}

	defer conn.Close(context.Background())

	var rows pgx.Rows
	marksQuery, ok := r.URL.Query()["mark"]
	var marksQ string
	var MarksQ  string
	if !ok || len(marksQuery[0]) < 1 {
		marksQ = ""
		MarksQ = ""
		rows, err = conn.Query(context.Background(),
			"select * from _stars LIMIT 18 OFFSET $1", 18*(key-1))
	}else {
		marksQ = marksQuery[0]
		MarksQ = strings.ToUpper(marksQ)
		rows, err = conn.Query(context.Background(),
			"select * from _stars WHERE mark=$1 OR mark=$2 LIMIT 18 OFFSET $3", marksQ, MarksQ, 18*(key-1))
	}

	helpers.ErrDefaultDetect(err, "QueryRow")

	iData := []models.StarsProduct{}

	for rows.Next(){
		err = rows.Scan(&data.Id, &data.Brand, &data.Mark, &data.Model, &data.Years, &data.Price, &data.Side, &data.Img)
		//data.Mark = marks[data.Mark]
		iData = append(iData, data)
		helpers.ErrDefaultDetect(err, "Row Scan")
	}
	defer rows.Close()

	if marksQ == ""{
		row := conn.QueryRow(context.Background(), "select COUNT(id) from _stars")
		err = row.Scan(&pages)
	}else{
		row := conn.QueryRow(context.Background(),
			"select COUNT(id) from _stars WHERE mark=$1 OR mark=$2", marksQ, MarksQ)
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
		Link: "http://localhost:8000/pads/",
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

	tmpl := template.Must(template.New("stars").Funcs(fm).
		ParseFiles(server + "templates/header.tmpl",
			server + "templates/stars.tmpl",
			server + "templates/footer.tmpl"))

	err = tmpl.ExecuteTemplate(w, "stars", output)

	helpers.ErrCatch(err, "Перевод в шаблон")
}