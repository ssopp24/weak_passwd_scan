package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanPostgres(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 5432, Protocol: "postgres", Username: "postgres", Password: "Zjw195126@"}
	t.Log(plugins.ScanPostgres(s))
}
