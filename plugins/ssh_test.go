package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanSsh(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 22, Username: "ssopp24", Password: "zjw195126", Protocol: "ssh"}
	t.Log(plugins.ScanSsh(s))
}
