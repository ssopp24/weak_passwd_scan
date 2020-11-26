/*
utils包包含:
	从redis获取输入参数并进行格式转化模块
	扫描并保存破解成功后结果模块
	将输出结果进行格式转化并传递给redis模块

	字典文件读取模块
	不同协议相应字段MD5值计算模块x`
*/

package utils

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
	"weak_passwd_scan/logs"
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"
	"weak_passwd_scan/utils/hash"
	"weak_passwd_scan/vars"
)

func Scan(task models.InputParam, scanResult *[]models.OutResult) (err error) {
	ip := task.Ip
	for i := 0; i < task.Num; i++ {
		service := task.Item[i]
		err := GenerateTask(ip, service, scanResult)
		if err != nil {
			logs.Log.Println("[error]	GenerateTask error: ", err.Error())
			return err
		}
	}

	return nil
}

func GenerateTask(ip string, service models.ScanParam, scanResult *[]models.OutResult) (err error) {
	tasks := make([]models.ScanTask, 0)

	protocol := strings.ToUpper(service.Protocol)
	if protocol == "SNMP" || protocol == "REDIS" {
		passwdDict, pErr := ReadPasswordDict(service.PasswdDict)
		if pErr != nil {
			logs.Log.Println("[error]	ReadPasswordDict error: ", err.Error())
			return err
		}

		for _, passwd := range passwdDict {
			task := models.ScanTask{Ip: ip, Port: service.Port, Protocol: service.Protocol, Username: "", Password: passwd}
			tasks = append(tasks, task)
		}
	} else {
		usernameDict, uErr := ReadUserDict(service.UserDict)
		if uErr != nil {
			logs.Log.Println("[error]	ReadUserDict error: ", err.Error())
			return err
		}

		passwdDict, pErr := ReadPasswordDict(service.PasswdDict)
		if pErr != nil {
			logs.Log.Println("[error]	ReadPasswordDict error: ", err.Error())
			return err
		}

		for _, user := range usernameDict {
			for _, passwd := range passwdDict {
				task := models.ScanTask{Ip: ip, Port: service.Port, Protocol: service.Protocol, Username: user, Password: passwd}
				tasks = append(tasks, task)
			}
		}
	}

	wg := &sync.WaitGroup{}
	taskChan := make(chan models.ScanTask, service.ThreadNum*2)

	for i := 0; i < service.ThreadNum; i++ {
		go crackPassword(taskChan, wg, scanResult)
	}

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	waitTimeout(wg, vars.TimeOut*2)

	return nil
}

func crackPassword(taskChan chan models.ScanTask, wg *sync.WaitGroup, scanResult *[]models.OutResult) {
	for task := range taskChan {

		/*测试日志*/
		logs.Log.Printf("[info]	Ip: %v, Port: %v, [%v], UserName: %v, Password: %v, goroutineNum: %v", task.Ip, task.Port,
			task.Protocol, task.Username, task.Password, runtime.NumGoroutine())

		var k string
		protocol := strings.ToUpper(task.Protocol)

		if protocol == "REDIS" || protocol == "SNMP" {
			k = fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.Username)
		}

		h := hash.MakeTaskHash(k)
		if hash.CheckTaskHash(h) {
			wg.Done()
			continue
		}

		fn := plugins.ScanFuncMap[protocol]
		err, result := fn(task)
		SaveResult(err, result, scanResult)
		wg.Done()
	}
}

func SaveResult(err error, result models.ScanResult, sumScanResult *[]models.OutResult) {
	if err == nil && result.Result {
		var k string
		protocol := strings.ToUpper(result.Task.Protocol)

		if protocol == "REDIS" || protocol == "SNMP" {
			k = fmt.Sprintf("%v-%v-%v", result.Task.Ip, result.Task.Port, result.Task.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", result.Task.Ip, result.Task.Port, result.Task.Username)
		}

		h := hash.MakeTaskHash(k)
		isExist := hash.SetTaskHask(h)

		if true != isExist {
			vars.Mutex.Lock()
			result := models.OutResult{Protocol: result.Task.Protocol, Port: result.Task.Port, Username: result.Task.Username, Passwd: result.Task.Password}
			*sumScanResult = append(*sumScanResult, result)
			vars.Mutex.Unlock()
		}
	}
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(timeout):
		return true
	}
}
