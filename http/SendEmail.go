package http

import (
	help "FullGas/helpers"
	"FullGas/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
)

func CreateOrder(w http.ResponseWriter, r *http.Request)  {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	help.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	vars := mux.Vars(r)
	var data int

	row := conn.QueryRow(context.Background(),
		"INSERT INTO orders (product, product_id) VALUES ($1, $2) RETURNING (id)", vars["product"], vars["id"])

	err = row.Scan(&data)
	help.ErrDefaultDetect(err, "QueryExec")

	Jdata, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s", Jdata)
}

func SendEmail(w http.ResponseWriter, r *http.Request)  {
	server := os.Getenv("SERVER")
	output := models.SendEmail()
	order := models.SendOrder{}

	if err := r.ParseForm(); err != nil {
		fmt.Println("ParseForm() err: ", err)
	}
	name := r.FormValue("firstName")
	phone := r.FormValue("phone")

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	help.ErrDefaultDetect(err, "DataBase Connection")
	defer conn.Close(context.Background())

	vars := mux.Vars(r)


	row := conn.QueryRow(context.Background(),
		"select product, product_id from orders WHERE id=$1", vars["orderId"])

	err = row.Scan(&order.Product, &order.ProductId)
	order.Id = vars["orderId"]

	var message string
	if order.Product == "tires" {
		var data struct{
			Id 			int
			Cub			int
			RadiusFront	int
			RadiusBack	int
			KitUnit		int
			Spike		int
			Price		int
		}

		row := conn.QueryRow(context.Background(),
			"select * from _tires WHERE id=$1", order.ProductId)

		err = row.Scan(&data.Id, &data.Cub, &data.RadiusFront, &data.RadiusBack, &data.KitUnit, &data.Spike, &data.Price)
		fmt.Println(data)
		message = "Новый заказ и это резина! \nИмя клиента: " + name + "\nНомер телефона: " + phone +
			"\nId товара: " + strconv.Itoa(data.Id) + "\nКубатура: " + strconv.Itoa(data.Cub) + "\nРадиус передний: " +
			strconv.Itoa(data.RadiusFront) + "\nРадиус задний: " + strconv.Itoa(data.RadiusBack) + "\nТип комплекта: " +
			strconv.Itoa(data.KitUnit) + "\nШипы: " + strconv.Itoa(data.Spike) + "\nЦена: " + strconv.Itoa(data.Price)
	}else if order.Product == "pads" {
		var data struct{
			Id 			int
			Mark		int
			Model		int
			Years		int
			Price		int
		}

		row := conn.QueryRow(context.Background(),
			"select * from _pads WHERE id=$1", order.ProductId)

		err = row.Scan(&data.Id, &data.Mark, &data.Model, &data.Years, &data.Price)
		fmt.Println(data)
		message = "Новый заказ и это колодки! \nИмя клиента: " + name + "\nНомер телефона: " + phone +
			"\nId товара: " + strconv.Itoa(data.Id) + "\nМодель: " + strconv.Itoa(data.Model) + "\nГоды модели: " +
			strconv.Itoa(data.Years) + "\nЦена: " + strconv.Itoa(data.Price)
	}else if order.Product == "chains" {
		var data struct{
			Id 			int
			Label		string
			Model		string
			Usable		int
			Price		int
			Desc		string
		}

		row := conn.QueryRow(context.Background(),
			"select * from _chains WHERE id=$1", order.ProductId)

		err = row.Scan(&data.Id, &data.Label, &data.Model, &data.Usable, &data.Price, &data.Desc)
		fmt.Println(data)
		message = "Новый заказ и это цепи! \nИмя клиента: " + name + "\nНомер телефона: " + phone +
			"\nId товара: " + strconv.Itoa(data.Id) + "\nМодель: " + data.Model + "\nНомер применяемости: " +
			strconv.Itoa(data.Usable) + "\nЦена: " + strconv.Itoa(data.Price)
	}

	help.ErrDefaultDetect(err, "QueryRow")

	// user we are authorizing as
	from := "aolychkin@gmail.com"

	// use we are sending email to
	to := "aolychkin@gmail.com"

	// server we are authorized to send email through
	host := "smtp.gmail.com"

	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", from, "Xz4lbm777%", host)

	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style

	if err := smtp.SendMail(host+":587", auth, from, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
	}
	fmt.Println("Email Sent! \n" + message)

	tmpl := template.Must(template.New("sendEmail").
		ParseFiles(server + "templates/header.tmpl",
			server + "templates/sendEmail.tmpl",
			server + "templates/footer.tmpl"))

	output.OrderData = order
	err = tmpl.ExecuteTemplate(w, "sendEmail", output)
}