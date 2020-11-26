package plugins

import (
	"github.com/jlaffaye/ftp"

	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"

	"fmt"
)

func ScanFtp(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	conn, err := ftp.DialTimeout(fmt.Sprintf("%v:%v", s.Ip, s.Port), vars.TimeOut)
	if err == nil {
		err = conn.Login(s.Username, s.Password)
		if err == nil {
			defer conn.Logout()

			result.Result = true
		}
	}

	return err, result
}
