package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanMysql(t *testing.T) {
	service := models.ScanTask{Ip: "192.168.28.191", Port: 2400, Protocol: "mysql", Username: "", Password: ""}
	t.Log(plugins.ScanMysql(service))
}
