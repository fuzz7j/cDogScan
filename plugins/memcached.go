package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"net"
	"strings"
	"time"
)

func MemcachedScan(info *config.Info) (result bool, err error) {
	target := fmt.Sprintf("%s:%v", info.Host, info.Port)
	client, err := net.DialTimeout("tcp", target, time.Duration(info.Timeout)*time.Second)
	if err == nil {
		err = client.SetDeadline(time.Now().Add(time.Duration(info.Timeout) * time.Second))
		if err == nil {
			_, err = client.Write([]byte("stats\n"))
			if err == nil {
				rev := make([]byte, 1024)
				n, err := client.Read(rev)
				if err == nil {
					if strings.Contains(string(rev[:n]), "STAT") {
						res := fmt.Sprintf("Memcached Unauthorized:%v", target)
						log.Logsuccess(res)
						result = true
					}
					client.Close()
				}
			}
		}
	}
	return result, err
}
