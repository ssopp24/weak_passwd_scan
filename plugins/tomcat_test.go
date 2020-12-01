package plugins_test

import (
	"testing"
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"
)

func TestScanTomcat(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 8080, Protocol: "tomcat", Username: "root", Password: "zjw195126"}
	t.Log(plugins.ScanTomcat(s))
}
