package env

import (
	er "golex/helpers/errCatch"
	"os"
)

func SetEnv() {
	err := os.Setenv("DATABASE_URL", "postgres://postgres:12345@localhost/fgmotoru")
	er.ErrDefaultDetect(err, "Ошибка установки ENV")
}