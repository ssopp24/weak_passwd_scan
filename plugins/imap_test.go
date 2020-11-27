package plugins_test

import (
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"

	"testing"
)

func TestScanImap(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.164", Port: 143, Protocol: "imap", Username: "imapuner", Password: "qscvhuksd"}
	t.Log(plugins.ScanImap(s))
}
