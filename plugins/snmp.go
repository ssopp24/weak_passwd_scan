package plugins

import (
	"github.com/gosnmp/gosnmp"

	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"
)

func ScanSNMP(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	gosnmp.Default.Target = s.Ip
	gosnmp.Default.Port = uint16(s.Port)
	gosnmp.Default.Community = s.Password
	gosnmp.Default.Timeout = vars.TimeOut

	err = gosnmp.Default.Connect()
	if err == nil {
		oids := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
		_, err := gosnmp.Default.Get(oids)
		if err == nil {
			result.Result = true
		}
	}

	return err, result
}
