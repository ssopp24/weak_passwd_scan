package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanMongodb(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 27017, Protocol: "mongodb", Username: "myUserAdmin", Password: "Zjw195126@"}
	t.Log(plugins.ScanMongodb(s))
}
