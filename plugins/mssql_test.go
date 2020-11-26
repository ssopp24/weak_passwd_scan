package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanMssql(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 1433, Protocol: "mssql", Username: "sa", Password: "Zjw195126@"}
	t.Log(plugins.ScanMssql(s))
}
