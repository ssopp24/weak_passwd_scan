/*
utils包包含:
	从redis获取输入参数并进行格式转化模块
	扫描并保存破解成功后结果模块
	将输出结果进行格式转化并传递给redis模块

	字典文件读取模块
	不同协议相应字段MD5值计算模块`
	oracle-sid猜测程序
*/

package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"weak_passwd_scan/logs"
	"weak_passwd_scan/models"
	"weak_passwd_scan/plugins"
	"weak_passwd_scan/utils/hash"
	"weak_passwd_scan/vars"
)

/*扫描函数最外层封装, 循环依次扫描本次Task包含的协议*/
func Scan(task models.InputParam, scanResult *[]models.OutResult) (err error) {
	vars.SuccessHash = make(map[string]bool)

	ip := task.Ip
	for i := 0; i < len(task.Item); i++ {
		service := task.Item[i]
		// 一次generateTask函数调用对应一种协议扫描
		// 目前Oracle扫描逻辑分为一类，其余协议扫描逻辑分为一类
		if "ORACLE" != strings.ToUpper(service.Protocol) {
			err := generateTask(ip, service, scanResult)
			if err != nil {
				logs.Log.Println("[error]	generateTask error: ", err.Error())
				return err
			}
		} else {
			sidNum, oracleSidArr := oracleSidGuess(ip, service.Port)
			if 0 == sidNum {
				continue
			} else {
				for j := 0; j < sidNum; j++ {
					vars.OracleGuessSid["oracleSid"] = oracleSidArr[j]
					err := generateTask(ip, service, scanResult)
					if err != nil {
						logs.Log.Println("[error]	generateTask error: ", err.Error())
						return err
					}
				}
			}
		}
	}

	return nil
}

/*组合Ip、Port、Protocol、Username、Password参数通过chan作为数据流通渠道，并发执行扫描*/
func generateTask(ip string, service models.ScanParam, scanResult *[]models.OutResult) error {
	tasks := make([]models.ScanTask, 0)

	protocol := strings.ToUpper(service.Protocol)
	if protocol == "SNMP" || protocol == "REDIS" {
		passwdDict, pErr := ReadPasswordDict(service.PasswdDict)
		if pErr != nil {
			logs.Log.Println("[error]	ReadPasswordDict error: ", pErr.Error())
			return pErr
		}

		for _, passwd := range passwdDict {
			task := models.ScanTask{Ip: ip, Port: service.Port, Protocol: service.Protocol, Username: "", Password: passwd}
			tasks = append(tasks, task)
		}
	} else {
		usernameDict, uErr := ReadUserDict(service.UserDict)
		if uErr != nil {
			logs.Log.Println("[error]	ReadUserDict error: ", uErr.Error())
			return uErr
		}

		passwdDict, pErr := ReadPasswordDict(service.PasswdDict)
		if pErr != nil {
			logs.Log.Println("[error]	ReadPasswordDict error: ", pErr.Error())
			return pErr
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
	wg.Wait()

	return nil
}

/*根据从chan获取的扫描任务数据，调用不同协议扫描函数，保存破解成功的结果*/
func crackPassword(taskChan chan models.ScanTask, wg *sync.WaitGroup, scanResult *[]models.OutResult) {
	for task := range taskChan {
		/*测试日志*/
		//logs.Log.Printf("[info]	Ip: %v, Port: %v, [%v], UserName: %v, Password: %v, goroutineNum: %v", task.Ip, task.Port,
		//	task.Protocol, task.Username, task.Password, runtime.NumGoroutine())

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

		/*
			注意!
			关注plugins.ScanFuncMap值，输入参数中协议名称和 代码中协议名称需对应（如sql_server，代码里为mssql）
		*/
		fn := plugins.ScanFuncMap[protocol]
		err, result := fn(task)
		saveResult(err, result, scanResult)
		wg.Done()
	}
}

func saveResult(err error, result models.ScanResult, sumScanResult *[]models.OutResult) {
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
		if !isExist {
			vars.Mutex.Lock()
			result := models.OutResult{Protocol: result.Task.Protocol, Port: result.Task.Port, Username: result.Task.Username, Passwd: result.Task.Password}
			*sumScanResult = append(*sumScanResult, result)
			vars.Mutex.Unlock()
		}
	}
}

func oracleSidGuess(ip string, port int) (int, []string) {
	/*
		注意!
		这里调用sidguess程序的路径值需修改，建议改为绝对路径，
		oracle sid字典的路径值需修改，建议改为绝对路径(字典文件可放在dictionaries文件夹下)
	*/
	cmd := exec.Command("/download/SIDGuesser/sidguess", "-i", ip, "-p", strconv.Itoa(port), "-d", "/download/SIDGuesser/oracle_sid.txt")
	buf, _ := cmd.Output()

	arr := bytes.Split(buf, []byte("\n"))
	arrSize := len(arr)

	var oracleSidArr []string
	var sidNum = 0
	for i := 0; i < arrSize; i++ {
		if bytes.Contains(arr[i], []byte("FOUND SID")) {
			containSidLine := bytes.Split(arr[i], []byte(" "))
			oracleSidArr = append(oracleSidArr, string(containSidLine[2]))
			sidNum += 1
		}
	}

	logs.Log.Println("[info]	sidNim: ", sidNum, " oracleArr: ", oracleSidArr)

	return sidNum, oracleSidArr
}
