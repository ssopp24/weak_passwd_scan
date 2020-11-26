package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanFtp(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 88, Protocol: "ftp", Username: "ssopp24", Password: "zjw195126"}
	t.Log(plugins.ScanFtp(s))
}
