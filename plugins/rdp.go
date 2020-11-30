package plugins

/*
#cgo CFLAGS: -I/usr/include/freerdp2/ -I/usr/include/winpr2
#cgo LDFLAGS: -lfreerdp2 -lwinpr2

#include <stdio.h>
#include <freerdp/freerdp.h>

// 返回 0表示登录成功，返回 -1表示因各种原因导致的登录失败
int ScanRdpC(char *ip, int port, char *userName, char *passWord){
	freerdp *instance = 0;
	int ret = -1;

	instance = freerdp_new();
	if (NULL == instance || FALSE == freerdp_context_new(instance)){
        return ret;
    }

	instance->settings->Username = userName;
    instance->settings->Password = passWord;
    instance->settings->IgnoreCertificate = TRUE;
    instance->settings->AuthenticationOnly = TRUE;
    instance->settings->ServerHostname = ip;
    instance->settings->ServerPort = port;
    instance->settings->Domain = "";
    freerdp_connect(instance);

	int err = 0;
	err = freerdp_get_last_error(instance->context);
	if (0==err){
		ret = 0;
	}
	freerdp_disconnect(instance);
    freerdp_free(instance);

	return ret;
}
*/
import "C"
import "weak_passwd_scan/models"

func ScanRdp(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	ip := C.CString(s.Ip)
	port := C.int(s.Port)
	userName := C.CString(s.Username)
	passWord := C.CString(s.Password)

	// 注意:环境需提前安装freerdp软件
	ret := C.ScanRdpC(ip, port, userName, passWord)
	if ret == 0 {
		result.Result = true
	}

	return err, result
}
