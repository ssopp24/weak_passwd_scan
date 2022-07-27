package plugins

import (
	"gopkg.in/mgo.v2"

	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"

	"fmt"
)

func ScanMongodb(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", s.Username, s.Password, s.Ip, s.Port, "admin")
	session, err := mgo.DialWithTimeout(url, vars.TimeOut)
	if err == nil {
		defer session.Close()

		err = session.Ping()
		if err == nil {
			result.Result = true
		}
	}

	return err, result
}
