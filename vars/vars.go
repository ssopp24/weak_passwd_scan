/*
utils包包含:
	常量定义
*/
package vars

import (
	"sync"
	"time"
)

var (
	TimeOut = 4 * time.Second
)

var (
	Mutex sync.Mutex

	// 标记特定服务的特定用户是否破解成功，成功的话不再尝试破解该用户
	SuccessHash map[string]bool
)

func init() {
	SuccessHash = make(map[string]bool)
}
