package plugins

import (
	"fmt"
	"net"
	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"

	"golang.org/x/crypto/ssh"
)

func ScanSsh(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	config := &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		Timeout: vars.TimeOut,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", s.Ip, s.Port), config)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		errRet := session.Run("echo xsec")
		if err == nil && errRet == nil {
			defer session.Close()
			result.Result = true
		}
	}
	return err, result
}
