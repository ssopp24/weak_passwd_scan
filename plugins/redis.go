package plugins

import (
	"github.com/go-redis/redis"

	"weak_passwd_scan/models"
	"weak_passwd_scan/vars"

	"fmt"
)

func ScanRedis(s models.ScanTask) (err error, result models.ScanResult) {
	result.Task = s
	opt := redis.Options{Addr: fmt.Sprintf("%v:%v", s.Ip, s.Port),
		Password: s.Password, DB: 0, DialTimeout: vars.TimeOut}
	client := redis.NewClient(&opt)
	defer client.Close()
	_, err = client.Ping().Result()
	if err == nil {
		result.Result = true
	}
	return err, result
}
