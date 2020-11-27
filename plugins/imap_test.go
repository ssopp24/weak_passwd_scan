package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanImap(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 143, Protocol: "imap", Username: "ssopp24", Password: "zjw195126"}
	t.Log(plugins.ScanImap(s))
}
