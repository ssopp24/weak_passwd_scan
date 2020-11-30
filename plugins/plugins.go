/*
plugins包包含:
	各协议扫描函数实现
*/
package plugins

import "weak_passwd_scan/models"

type ScanFunc func(service models.ScanTask) (err error, result models.ScanResult)

var (
	ScanFuncMap map[string]ScanFunc
)

func init() {
	ScanFuncMap = make(map[string]ScanFunc)
	ScanFuncMap["FTP"] = ScanFtp
	ScanFuncMap["SSH"] = ScanSsh
	ScanFuncMap["SMB"] = ScanSmb
	ScanFuncMap["MSSQL"] = ScanMssql
	ScanFuncMap["MYSQL"] = ScanMysql
	ScanFuncMap["POSTGRESQL"] = ScanPostgres
	ScanFuncMap["REDIS"] = ScanRedis
	ScanFuncMap["MONGODB"] = ScanMongodb
	ScanFuncMap["SNMP"] = ScanSNMP
	ScanFuncMap["TELNET"] = ScanTelnet
	ScanFuncMap["SMTP"] = ScanSmtp
	ScanFuncMap["IMAP"] = ScanImap
	ScanFuncMap["POP3"] = ScanPop3
	ScanFuncMap["ORACLE"] = ScanOracle
	ScanFuncMap["TOMCAT"] = ScanTomcat
	ScanFuncMap["RDP"] = ScanRdp
	ScanFuncMap["DB2"] = ScanDb2
}
