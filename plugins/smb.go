package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"github.com/stacktitan/smb/smb"
	"strings"
	"time"
)

func SmbScan(info *config.Info) (result bool, err error) {
	for _, user := range config.UserList["smb"] {
		for _, pass := range config.Passwords {
			pass = strings.Replace(pass, "{user}", user, -1)
			flag, err := doWithTimeOut(info, user, pass)
			if flag == true && err == nil {
				res := fmt.Sprintf("[SMB]%v:%v %v/%v", info.Host, info.Port, user, pass)
				log.Logsuccess(res)
				result = true
			}
		}
	}
	return result, err
}

func SmbConn(info *config.Info, user string, pass string, signal chan struct{}) (flag bool, err error) {
	flag = false
	Host, Username, Password := info.Host, user, pass
	options := smb.Options{
		Host:        Host,
		Port:        445,
		User:        Username,
		Password:    Password,
		Domain:      "",
		Workstation: "",
	}
	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			flag = true
		}
	}
	signal <- struct{}{}
	return flag, err
}

func doWithTimeOut(info *config.Info, user string, pass string) (flag bool, err error) {
	signal := make(chan struct{})
	go func() {
		flag, err = SmbConn(info, user, pass, signal)
	}()
	select {
	case <-signal:
		return flag, err
	case <-time.After(time.Duration(info.Timeout) * time.Second):
		return false, err
	}
}