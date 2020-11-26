package hash

import (
	"weak_passwd_scan/vars"
)

func MakeTaskHash(k string) string {
	hash := MD5(k)
	return hash
}

func CheckTaskHash(hash string) (isExist bool) {
	isExist = vars.SuccessHash[hash]

	return isExist
}

func SetTaskHask(hash string) (isExist bool) {
	vars.Mutex.Lock()
	if true == vars.SuccessHash[hash] {
		isExist = true
	} else {
		vars.SuccessHash[hash] = true
		isExist = false
	}
	vars.Mutex.Unlock()

	return isExist
}
