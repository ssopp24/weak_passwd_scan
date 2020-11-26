package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanSNMP(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 161, Username: "", Password: "public", Protocol: "snmp"}
	t.Log(plugins.ScanSNMP(s))
}
