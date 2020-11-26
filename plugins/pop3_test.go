package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanPop3(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 155, Protocol: "pop3", Username: "test", Password: "zjw195126"}
	t.Log(plugins.ScanPop3(s))
}
