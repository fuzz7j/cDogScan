package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func MysqlScan(info *config.Info) (result bool, err error) {
	for _, user := range config.UserList["mysql"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",user, password, info.Host, info.Port, "mysql")
			db, err := sql.Open("mysql", dataSourceName)
			if err != nil {
					return result, err
			}
			err = db.Ping()
			if err == nil {
				res := fmt.Sprintf("[MYSQL]%v:%v %v/%v", info.Host, info.Port, user, password)
				log.Logsuccess(res)
				result = true
			}
		}
	}
	return result, err
}


