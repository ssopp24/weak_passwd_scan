package utils

import (
	"weak_passwd_scan/logs"
	"weak_passwd_scan/models"

	"github.com/go-redis/redis"

	"encoding/json"
)

// 从redis获取输入的json格式参数，转化参数格式为Task格式
func GetInputParameter(client *redis.Client, task *models.InputParam) (err error) {
	/*
		注意!
		key名字要改
	*/
	inputParam, err := client.BLPop(0, "list1").Result()
	if err != nil {
		logs.Log.Println("[error]	BLPop error: ", err.Error())
		return err
	}

	/*
		注意!
		反序列化时可能会丢失一些字段，若丢失，则需要更改输入参数字段名与models.go中InputParam与ScanParam相同
	*/
	err = json.Unmarshal([]byte(inputParam[1]), task)
	if err != nil {
		logs.Log.Println("[error]	Unmarshal error: ", err.Error())
		return err
	}

	logs.Log.Println("[info]	Input parameters: ", *task)
	return nil
}
