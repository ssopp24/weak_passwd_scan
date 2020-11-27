/*定义日志输出位置与日志格式*/
package logs

import (
	"fmt"
	"log"
	"os"
)

var (
	Log log.Logger
)

func init() {
	/*
		注意!!!!
		log文件存储路径要改，建议改为绝对路径
	*/
	logFile, err := os.OpenFile("./logs/weak_passwd_scan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	Log.SetOutput(logFile)
	Log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	Log.SetPrefix("[weak_passwd_scan]")
}
