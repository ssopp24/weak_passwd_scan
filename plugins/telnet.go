package plugins

import (
	"github.com/ziutek/telnet"

	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"

	"fmt"
	"time"
)

var scanTelnetFlag = true

func checkErr(err error) {
	if err != nil {
		scanTelnetFlag = false
	}
}

func sendln(t *telnet.Conn, s string) {
	checkErr(t.SetWriteDeadline(time.Now().Add(vars.TimeOut)))
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	checkErr(err)
}

func expect(t *telnet.Conn, d ...string) {
	checkErr(t.SetReadDeadline(time.Now().Add(vars.TimeOut)))
	checkErr(t.SkipUntil(d...))
}

func ScanTelnet(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s

	t, err := telnet.DialTimeout("tcp", fmt.Sprintf("%v:%v", s.Ip, s.Port), vars.TimeOut)
	if nil == err {
		expect(t, "login: ")
		sendln(t, s.Username)
		expect(t, "Password: ")
		sendln(t, s.Password)
		expect(t, "$")
		t.Close()
	} else {
		scanTelnetFlag = false
	}

	result.Result = scanTelnetFlag

	return err, result
}
