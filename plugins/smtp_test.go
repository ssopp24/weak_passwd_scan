package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanSmtp(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 29, Protocol: "smtp", Username: "ssopp24@ssopp24.com.cn", Password: "zjw195126"}
	t.Log(plugins.ScanSmtp(s))
}
