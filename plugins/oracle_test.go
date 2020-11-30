package plugins_test

import (
	"testing"
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"
)

func TestScanOracle(t *testing.T) {
	s := models.ScanTask{Ip: "", Port: 22, Protocol: "", Username: "", Password: ""}
	t.Log(plugins.ScanOracle(s))
}
