package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"net"
	"strings"
	"time"
)

func ZookeeperScan(info *config.Info) (result bool, err error) {
	target := fmt.Sprintf("%v:%v", info.Host, info.Port)
	conn, err := net.DialTimeout("tcp", target, time.Duration(info.Timeout)*time.Second)
	if err == nil {
		err = conn.SetDeadline(time.Now().Add(time.Duration(info.Timeout) * time.Second))
		if err == nil {
			_, err := conn.Write([]byte("envi\n"))
			if err == nil {
				rev := make([]byte, 1024)
				n, err := conn.Read(rev)
				if err == nil {
					if strings.Contains(string(rev[:n]), "Environment") {
						res := fmt.Sprintf("[Zookeeper unauthorized]%v", target)
						log.Logsuccess(res)
						result = true
					}
					conn.Close()
				}
			}
		}
	}
	return result, err
}
