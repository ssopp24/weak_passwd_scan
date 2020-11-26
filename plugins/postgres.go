package plugins

import (
	"weak_passwd_scan/models"

	"database/sql"
	"fmt"
)

func ScanPostgres(service models.ScanTask) (err error, result models.ScanResult) {
	result.Task = service
	result.Result = false

	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", service.Username,
		service.Password, service.Ip, service.Port, "postgres", "disable")
	db, err := sql.Open("postgres", dataSourceName)

	if err == nil {
		defer db.Close()

		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}
	return err, result
}
