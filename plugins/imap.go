package plugins

import (
	"github.com/emersion/go-imap/client"

	"weak_passwd_scan/models"

	"fmt"
)

func ScanImap(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	c, err := client.Dial(fmt.Sprintf("%v:%v", s.Ip, s.Port))
	if err == nil {
		defer c.Logout()

		if err := c.Login(s.Username, s.Password); err == nil {
			result.Result = true
		}
	}

	return err, result
}
