package env

import (
	"log"
	"os"
)

func SetEnv() {
	err := os.Setenv("DATABASE_URL", "postgres://postgres:12345@localhost/fgmotoru")
	if err != nil {
		log.Fatal("ListenAndServe Ошибка установки ENV : ", err)
	}
}