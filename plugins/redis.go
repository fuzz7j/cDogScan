package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

func RedisScan(info *config.Info) (result bool, err error) {
	for _, password := range config.Passwords {
		password = strings.Replace(password, "{user}", "redis", -1)
		opt := redis.Options{Addr: fmt.Sprintf("%v:%v", info.Host, info.Port), Password:password, DB: 0, DialTimeout: time.Duration(info.Timeout) * time.Second}
		client := redis.NewClient(&opt)
		_, err = client.Ping().Result()
		if err != nil {
			return result, err
		}
		res := fmt.Sprintf("[Redis]%v:%v %v", info.Host, info.Port, password)
		log.Logsuccess(res)
		result = true
		defer func() {
			if client != nil {
				_ = client.Close()
			}
		}()
	}
	return result, err
}
