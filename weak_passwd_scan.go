/*
程序主体流程控制:
	获取输入参数并进行格式转化
	执行扫描，保存破解成功后结果
	转化扫描结果格式，发送至redis
*/
package main

import (
	"weak_passwd_scan/logs"
	"weak_passwd_scan/models"
	"weak_passwd_scan/utils"

	"github.com/go-redis/redis"
)

func main() {
	for {
		//	连接redis
		/*
			注意!
			这里IP，端口，密码需修改
		*/
		rdb := redis.NewClient(&redis.Options{
			Addr:     "192.168.28.191:6379",
			Password: "",
			DB:       0,
		})
		//defer rdb.Close()
		logs.Log.Println("")

		//	获取输入参数并转化格式
		var task models.InputParam
		err := utils.GetInputParameter(rdb, &task)
		if err != nil {
			logs.Log.Println("[error]	GetInputParameter error: ", err.Error())
			panic(err.Error())
		}

		//执行扫描并保存破解成功后结果
		var scanResult []models.OutResult
		err = utils.Scan(task, &scanResult)
		if err != nil {
			logs.Log.Println("[error]	Scan error: ", err.Error())
			panic(err.Error())
		}

		// 将破解成功的协议转化为json格式并存入redis
		err = utils.SendOutputParameter(rdb, task, scanResult)
		if err != nil {
			logs.Log.Println("[error]	SendOutputParameter error: ", err.Error())
			panic(err.Error())
		}
	}
}
