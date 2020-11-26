package plugins

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"

	"fmt"
	"net"
	"net/smtp"
)

func smtpDialTimeout(addr string) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, vars.TimeOut)
	if err != nil {
		return nil, err
	}

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func ScanSmtp(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	c, err := smtpDialTimeout(fmt.Sprintf("%v:%v", s.Ip, s.Port))
	if err == nil {
		auth := smtp.CRAMMD5Auth(s.Username, s.Password)
		err := c.Auth(auth)
		if err == nil {
			result.Result = true
		}
		c.Quit()
	}

	return err, result
}
