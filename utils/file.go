package utils

import (
	"bufio"
	"os"
	"strings"
	"weak_passwd_scan/logs"
)

/*
注意!
输入参数中字典路径需为绝对路径
*/
func ReadUserDict(userDict string) (users []string, err error) {
	file, err := os.Open(userDict)
	if err != nil {
		logs.Log.Println("[error]	Open error: ", err.Error())
		return users, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	users = append(users, "")

	return users, err
}

/*
注意!
输入参数中字典路径需为绝对路径
*/
func ReadPasswordDict(passDict string) (password []string, err error) {
	file, err := os.Open(passDict)
	if err != nil {
		logs.Log.Println("[error]	Open error: ", err.Error())
		return password, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			password = append(password, passwd)
		}
	}
	password = append(password, "")

	return password, err
}
