package plugins

import (
	_ "github.com/denisenkom/go-mssqldb"

	"weak_passwd_scan/models"

	"database/sql"
	"fmt"
)

func ScanMssql(service models.ScanTask) (err error, result models.ScanResult) {
	result.Task = service
	result.Result = false

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", service.Ip,
		service.Port, service.Username, service.Password, "master")
	db, err := sql.Open("mssql", dataSourceName)
	if err == nil {
		defer db.Close()

		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return err, result
}
