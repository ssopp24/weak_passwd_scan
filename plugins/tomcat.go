package plugins

import (
	"fmt"
	"weak_passwd_scan/models"

	"net/http"
)

func ScanTomcat(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	scanUrl := fmt.Sprintf("%v%v:%v%v", "http://", s.Ip, s.Port, "/manager/html")
	req, err := http.NewRequest("GET", scanUrl, nil)
	req.SetBasicAuth(s.Username, s.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err == nil && 200 == resp.StatusCode {
		result.Result = true
	}

	return err, result
}
