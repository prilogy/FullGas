package http

import (
	"FullGas/helpers"
	"FullGas/models"
	"html/template"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	server := os.Getenv("SERVER")
	data := models.Index()

	templates, err := template.ParseFiles(
		server + "templates/header.tmpl",
		server + "templates/index.tmpl",
		server + "templates/footer.tmpl")

	helpers.ErrCatch(err, "Парсинг файлов")
	
	tmpl := templates.Lookup("index.tmpl")
	err = tmpl.ExecuteTemplate(w, "index", data)

	helpers.ErrCatch(err, "Перевод в шаблон")
}