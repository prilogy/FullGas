package http

import (
	er "golex/helpers/errCatch"
	"golex/models"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	data := models.Index()

	templates, err := template.ParseFiles(
		"templates/header.tmpl",
		"templates/index.tmpl",
		"templates/footer.tmpl")

	tmpl := templates.Lookup("index.tmpl")
	err = tmpl.ExecuteTemplate(w, "index", data)

	er.ErrCatch(err, "Перевод в шаблон")
}