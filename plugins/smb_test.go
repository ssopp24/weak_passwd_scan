package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanSmb(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 139, Protocol: "smb", Username: "ssopp24", Password: "zjw195126"}
	t.Log(plugins.ScanSmb(s))
}
