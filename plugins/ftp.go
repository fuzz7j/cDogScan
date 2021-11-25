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
	starttime := time.Now().Unix()
	flag, err := FtpConn(info,"anonymous","")
	if flag == true && err == nil {
		return flag,nil
	} else {
		res := fmt.Sprintf("[-]FTP:%v:%v %v %v", info.Host, info.Port, "anonymous", err)
		log.LogError(res)
	}

	for _, user := range config.UserList["ftp"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			flag, err := FtpConn(info, user, password)
			if flag == true && err == nil {
				return flag,nil
			} else {
				res := fmt.Sprintf("[-]FTP:%v:%v %v %v", info.Host, info.Port, user, password)
				log.LogError(res)
				if time.Now().Unix() - starttime > (int64(len(config.UserList["ftp"]) * len(config.Passwords)) * info.Timeout) {
					res := fmt.Sprintf("[Error]%v:%v ", info.Host, info.Port)
					log.LogError(res)
					result = false
					return result,err
				}
			}
		}
	}
	return result, err
}

func FtpConn(info *config.Info, user, pass string) (flag bool, err error) {
	flag = false
	conn, err := ftp.DialTimeout(fmt.Sprintf("%v:%v", info.Host, info.Port), time.Duration(info.Timeout) * time.Second)
	if err == nil {
		err = conn.Login(user,pass)
		if err == nil {
			flag = true
			res := fmt.Sprintf("FTP:%v:%v %v %v", info.Host, info.Port, user, pass)
			dirs, err := conn.List("")
			if err == nil {
				if len(dirs) > 0 {
					for i := 0; i < len(dirs); i++ {
						if len(dirs[i].Name) > 50 {
							res += "\n   [->]" + dirs[i].Name[:50]
						} else {
							res += "\n   [->]" + dirs[i].Name
						}
						if i == 5 {
							break
						}
					}
				}
			}
			log.Logsuccess(res)
		}
	}
	return flag, err
}