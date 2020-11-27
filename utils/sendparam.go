package utils

import (
	"encoding/json"
	"weak_passwd_scan/logs"
	"weak_passwd_scan/models"

	"github.com/go-redis/redis"
)

func SendOutputParameter(client *redis.Client, inputParam models.InputParam, outResult []models.OutResult) (err error) {
	outParam := models.OutParam{TaskId: inputParam.TaskId, Ip: inputParam.Ip, Num: len(outResult), Item: outResult}
	/*
		注意!
		输出参数与之前约定字段名称不同，参考models中OutParam与OutResult
	*/
	outParamJson, _ := json.Marshal(outParam)

	/*
		注意!
		key名字要改
	*/
	err = client.LPush("list2", outParamJson).Err()
	if err != nil {
		logs.Log.Println("[error]	LPush error: ", err.Error())
		return err
	}

	logs.Log.Println("[info]	Output parameters: ", outParam)
	return nil
}
