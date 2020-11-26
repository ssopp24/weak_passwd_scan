package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanRedis(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 6379, Password: "Zjw195126@"}
	t.Log(plugins.ScanRedis(s))
}
