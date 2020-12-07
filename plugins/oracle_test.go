package plugins_test

import (
	"testing"
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"
)

func TestScanOracle(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 1521, Protocol: "oracle", Username: "system", Password: "oracle"}
	t.Log(plugins.ScanOracle(s))
}
