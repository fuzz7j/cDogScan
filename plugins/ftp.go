package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"github.com/jlaffaye/ftp"
	"strings"
	"time"
)

func FtpScan(info *config.Info) (result bool, err error) {
	for _, user := range config.UserList["ftp"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			conn, err := ftp.DialTimeout(fmt.Sprintf("%v:%v", info.Host, info.Port), time.Duration(info.Timeout)*time.Second)
			if err == nil {
				err = conn.Login(user, password)
				if err == nil {
					defer func() {
						err = conn.Logout()
					}()
					res := fmt.Sprintf("[FTP]%v:%v %v/%v", info.Host, info.Port, user, password)
					log.Logsuccess(res)
					result = true
				}
			}
		}
	}
	return result, err
}
