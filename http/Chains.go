package http

import (
	"../helpers"
	"../models"
	"context"
	"github.com/jackc/pgx"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Chains(w http.ResponseWriter, r *http.Request)  {
	output := models.Chains()
	data := models.ChainsProduct{}

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

	marks := []string{"", "Motocross", "Road/offroad", "Street"}

	defer conn.Close(context.Background())

	var rows pgx.Rows
	marksQuery, ok := r.URL.Query()["mark"]
	var marksQ string
	if !ok || len(marksQuery[0]) < 1 {
		marksQ = ""
		rows, err = conn.Query(context.Background(),
			"select * from _chains LIMIT 20 OFFSET $1", 20*(key-1))
	}else {
		marksQ = marksQuery[0]
		rows, err = conn.Query(context.Background(),
			"select * from _chains WHERE usable=$1 LIMIT 20 OFFSET $2", marksQ, 20*(key-1))
	}

	helpers.ErrDefaultDetect(err, "QueryRow")

	iData := []models.ChainsProduct{}

	for rows.Next(){
		var model string
		err = rows.Scan(&data.Id, &data.Label, &model, &data.Mark, &data.Price, &data.Description)
		data.Model = strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}
			return r
		}, model)
		data.MarkName = marks[data.Mark]
		iData = append(iData, data)
		helpers.ErrDefaultDetect(err, "Row Scan")
	}
	defer rows.Close()

	if marksQ == ""{
		row := conn.QueryRow(context.Background(), "select COUNT(id) from _chains")
		err = row.Scan(&pages)
	}else{
		row := conn.QueryRow(context.Background(),
			"select COUNT(id) from _chains WHERE mark=$1 LIMIT 20 OFFSET $2", marksQ, 20*(key-1))
		err = row.Scan(&pages)
	}

	output.Chains = append(output.Chains, iData...)
	if marksQ != ""{
		output.PageMark = marksQ
	}
	temp := 1+pages/20
	output.Page = models.Pagination{
		Start: 1,
		End: temp,
		AllData: pages,
		Current: key,
		Link: "http://localhost:8181/chains/",
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

	tmpl := template.Must(template.New("chains").Funcs(fm).
		ParseFiles("templates/header.tmpl",
			"templates/chains.tmpl",
			"templates/footer.tmpl"))

	err = tmpl.ExecuteTemplate(w, "chains", output)

	helpers.ErrCatch(err, "Перевод в шаблон")
}