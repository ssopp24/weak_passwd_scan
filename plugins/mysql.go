package plugins

import (
	_ "github.com/netxfly/mysql"

	"weak_passwd_scan/models"

	"database/sql"
	"fmt"
)

func ScanMysql(service models.ScanTask) (err error, result models.ScanResult) {
	result.Task = service
	result.Result = false

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", service.Username,
		service.Password, service.Ip, service.Port, "mysql")
	db, err := sql.Open("mysql", dataSourceName)
	if err == nil {
		defer db.Close()

		err = db.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return err, result
}
