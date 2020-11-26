package utils

import (
	"weak_passwd_scan/logs"
	"weak_passwd_scan/models"

	"github.com/go-redis/redis"

	"encoding/json"
)

// 从redis获取输入的json格式参数，转化参数格式为Task格式
func GetInputParameter(client *redis.Client, task *models.InputParam) (err error) {
	//测试：模拟构造json输入参数，并存进redis		begin
	//test1 := models.ScanParam{Protocol: "ssh", Port: 22, ThreadNum: 4, UserDict: "./dictionaries/ssh_default_username.txt", PasswdDict: "./dictionaries/ssh_default_password.txt"}
	//test2 := models.ScanParam{Protocol: "redis", Port: 6379, ThreadNum: 3, UserDict: "", PasswdDict: "./dictionaries/redis_default_password.txt"}
	//test3 := models.ScanParam{Protocol: "ftp", Port: 88, ThreadNum: 3, UserDict: "./dictionaries/ftp_default_username.txt", PasswdDict: "./dictionaries/ftp_default_password.txt"}
	//
	//var testServiceArr []models.ScanParam
	//testServiceArr = append(testServiceArr, test1)
	//testServiceArr = append(testServiceArr, test2)
	//testServiceArr = append(testServiceArr, test3)
	//
	//testInput := models.InputParam{TaskId: 1, Ip: "192.168.28.191", Num: len(testServiceArr), Item: testServiceArr}
	//
	//testInputJson, _ := json.Marshal(testInput)
	//client.LPush("list1", testInputJson)
	//测试：模拟构造json输入参数，并存进redis		end

	inputParam, err := client.BLPop(0, "list1").Result()
	if err != nil {
		logs.Log.Println("[error]	BLPop error: ", err.Error())
		return err
	}

	err = json.Unmarshal([]byte(inputParam[1]), task)
	if err != nil {
		logs.Log.Println("[error]	Unmarshal error: ", err.Error())
		return err
	}

	logs.Log.Println("[info]	Input parameters: ", *task)
	return nil
}
