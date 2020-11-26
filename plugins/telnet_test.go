package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanTelnet(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 28, Protocol: "telnet", Username: "ssopp24", Password: "zjw195126"}
	t.Log(plugins.ScanTelnet(s))
}
