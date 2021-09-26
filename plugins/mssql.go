package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"strings"
	"time"
)

func MssqlScan(info *config.Info) (result bool, err error) {
	for _, user := range config.UserList["mssql"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;encrypt=disable;timeout=%v", info.Host, info.Port, user, password, time.Duration(info.Timeout)*time.Second)
			db, err := sql.Open("mssql", dataSourceName)
			if err == nil {
				defer func() {
					err = db.Close()
				}()
				err = db.Ping()
				if err == nil {
					res := fmt.Sprintf("[MSSQL]%v:%v %v/%v", info.Host, info.Port, user, password)
					log.Logsuccess(res)
					result = true
				}
			}
		}
	}
	return result, err
}
