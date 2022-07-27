package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanMongodb(t *testing.T) {
	s := models.ScanTask{Ip: "127.0.0.1", Port: 27017, Protocol: "mongodb", Username: "admin", Password: "123456"}
	t.Log(plugins.ScanMongodb(s))
}
