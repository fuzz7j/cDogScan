package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func MysqlScan(info *config.Info) (result bool, err error) {
	starttime := time.Now().Unix()
	flag, err := MysqlConn(info,"","")
	if flag == true && err == nil {
		return flag,nil
	}

	for _, user := range config.UserList["mysql"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			result, err := MysqlConn(info, user, password)
			if result == true && err == nil {
				return result, err
			} else {
					res := fmt.Sprintf("[-]MYSQL:%v:%v %v %v", info.Host, info.Port, user, password)
					log.LogError(res)
				if time.Now().Unix() - starttime > (int64(len(config.UserList["mysql"]) * len(config.Passwords)) * info.Timeout) {
					res := fmt.Sprintf("[Error]%v:%v", info.Host, info.Port)
					log.LogError(res)
					result = false
					return result,err
				}
			}
		}
	}
	return result, err
}

func MysqlConn(info *config.Info, user, pass string) (flag bool, err error) {
	flag = false
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/mysql?charset=utf8&timeout=%v",user, pass, info.Host, info.Port, time.Duration(info.Timeout) * time.Second)
	db, err := sql.Open("mysql", dataSourceName)
	if err == nil {
		db.SetConnMaxLifetime(time.Duration(info.Timeout) * time.Second)
		db.SetConnMaxIdleTime(time.Duration(info.Timeout) * time.Second)
		db.SetMaxIdleConns(0)
		defer db.Close()
		err = db.Ping()
		if err == nil {
			if user == "" && pass == "" {
				res := fmt.Sprintf("MYSQL:%v:%v NULL", info.Host, info.Port)
				log.Logsuccess(res)
				flag = true
			} else {
				res := fmt.Sprintf("MYSQL:%v:%v %v %v", info.Host, info.Port, user, pass)
				log.Logsuccess(res)
				flag = true
			}
		}
	}
	return flag, err
}


