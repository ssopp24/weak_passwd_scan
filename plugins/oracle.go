package plugins

import (
	_ "github.com/sijms/go-ora"

	"database/sql"
	"fmt"

	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"
)

func ScanOracle(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	oracleLoginInfo := fmt.Sprintf("oracle://%v:%v@%v:%v/%v", s.Username, s.Password, s.Ip, s.Port, vars.OracleGuessSid["oracleSid"])
	conn, err := sql.Open("oracle", oracleLoginInfo)
	if err == nil {
		err = conn.Ping()
		if err == nil {
			conn.Close()
			result.Result = true
		}
	}

	return err, result
}
