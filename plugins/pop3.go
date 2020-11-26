package plugins

import (
	"github.com/taknb2nch/go-pop3"

	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"

	"fmt"
	"net"
	"time"
)

func pop3DialTimeout(addr string, timeout time.Duration) (*pop3.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)

	if err != nil {
		return nil, err
	}

	return pop3.NewClient(conn)
}

func ScanPop3(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	client, err := pop3DialTimeout(fmt.Sprintf("%v:%v", s.Ip, s.Port), vars.TimeOut)
	if err == nil {
		defer func() {
			client.Quit()
			client.Close()
		}()

		if err = client.User(s.Username); err == nil {
			if err = client.Pass(s.Password); err == nil {
				result.Result = true
			}
		}
	}

	return err, result
}
