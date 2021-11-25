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
	starttime := time.Now().Unix()
	flag, err := RedisConn(info, "")
	if flag == true && err == nil {
		return flag,nil
	}
	for _, password := range config.Passwords {
		password = strings.Replace(password, "{user}", "redis", -1)
		result, err := RedisConn(info, password)
		if result == true && err == nil {
			return result, err
		} else {
			res := fmt.Sprintf("[-]Redis:%v:%v %v", info.Host, info.Port, password)
			log.LogError(res)
			if time.Now().Unix() - starttime > (int64(len(config.Passwords)) * info.Timeout) {
				res := fmt.Sprintf("[Error]%v:%v", info.Host, info.Port)
				log.LogError(res)
				result = false
				return result,err
			}
		}
	}
	return result, err
}

func RedisConn(info *config.Info, password string) (flag bool, err error) {
	opt := redis.Options{Addr: fmt.Sprintf("%v:%v", info.Host, info.Port), Password:password, DB: 0, DialTimeout: time.Duration(info.Timeout) * time.Second}
	client := redis.NewClient(&opt)
	_, err = client.Ping().Result()
	if err == nil {
		if password == "" {
			res := fmt.Sprintf("Redis Unauthorized:%v:%v", info.Host, info.Port)
			log.Logsuccess(res)
		} else {
			res := fmt.Sprintf("Redis:%v:%v %v", info.Host, info.Port, password)
			log.Logsuccess(res)
		}
		flag = true
		defer func() {
			if client != nil {
				_ = client.Close()
			}
		}()
	}
	return flag, err
}