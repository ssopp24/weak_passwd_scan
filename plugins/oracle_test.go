package plugins_test

import (
	"testing"
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"
)

/*
	注意!
	运行此测试文件时，需修改oracle.go文件，参考oracle.go的17行
*/
func TestScanOracle(t *testing.T) {
	s := models.ScanTask{Ip: "192.168.28.191", Port: 1521, Protocol: "oracle", Username: "system", Password: "oracle"}
	t.Log(plugins.ScanOracle(s))
}
