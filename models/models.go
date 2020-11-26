/*全局数据结构定义*/
package models

type InputParam struct {
	TaskId int
	Ip     string
	Num    int
	Item   []ScanParam
}

type ScanParam struct {
	Protocol   string
	Port       int
	ThreadNum  int
	UserDict   string
	PasswdDict string
}

type ScanTask struct {
	Ip       string
	Port     int
	Protocol string
	Username string
	Password string
}

type OutResult struct {
	Protocol string
	Port     int
	Username string
	Passwd   string
}

type OutParam struct {
	TaskId int
	Ip     string
	Num    int
	Item   []OutResult
}

type ScanResult struct {
	Task   ScanTask
	Result bool
}
