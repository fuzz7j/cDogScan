package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

func PostgresScan(info *config.Info) (result bool, err error) {
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
					res := fmt.Sprintf("[Postgres]%v:%v %v/%v", info.Host, info.Port, user, password)
					log.Logsuccess(res)
					result = true
				}
			}
		}
	}
	return result, err
}
