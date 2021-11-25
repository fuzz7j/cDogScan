package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

func PostgresScan(info *config.Info) (result bool, err error) {
	starttime := time.Now().Unix()
	for _, user := range config.UserList["postgres"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", user, password, info.Host, info.Port, "postgres", "disable")
			db, err := sql.Open("postgres", dataSourceName)
			if err == nil {
				defer func() {
					err = db.Close()
				}()
				err = db.Ping()
				if err == nil {
					res := fmt.Sprintf("Postgres:%v:%v %v/%v", info.Host, info.Port, user, password)
					log.Logsuccess(res)
					result = true
				} else {
					res := fmt.Sprintf("[-]Postgres:%v:%v %v %v", info.Host, info.Port, user, password)
					log.LogError(res)
					if time.Now().Unix() - starttime > (int64(len(config.UserList["postgres"]) * len(config.Passwords)) * info.Timeout) {
						res := fmt.Sprintf("[Error]%v:%v", info.Host, info.Port)
						log.LogError(res)
						result = false
						return result,err
					}
				}
			}
		}
	}
	return result, err
}
