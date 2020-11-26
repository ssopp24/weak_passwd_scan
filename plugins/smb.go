package plugins

import (
	"github.com/stacktitan/smb/smb"

	"weak_passwd_scan/models"
)

func ScanSmb(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	options := smb.Options{
		Host:        s.Ip,
		Port:        s.Port,
		User:        s.Username,
		Password:    s.Password,
		Domain:      "",
		Workstation: "",
	}

	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()

		if session.IsAuthenticated {
			result.Result = true
		}
	}
	return err, result
}
