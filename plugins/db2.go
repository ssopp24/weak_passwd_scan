package plugins

import "weak_passwd_scan/models"

func ScanDb2(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	result.Result = false

	return err, result
}
