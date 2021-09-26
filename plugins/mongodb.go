package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"gopkg.in/mgo.v2"
	"strings"
	"time"
)

func MongodbScan(info *config.Info) (result bool, err error) {
	for _, user := range config.UserList["mongodb"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			url := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", user, password, info.Host, info.Port, "test")
			session, err := mgo.DialWithTimeout(url, time.Duration(info.Timeout)*time.Second)
			if err == nil {
				defer session.Close()
				err = session.Ping()
				if err == nil {
					res := fmt.Sprintf("[MongoDB]%v:%v %v/%v", info.Host, info.Port, user, password)
					log.Logsuccess(res)
					result = true
				}
			}
		}
	}
	return result, err
}
