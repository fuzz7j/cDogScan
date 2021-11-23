package common

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type IPAddr struct {
	Ip   string
	Port int
}

var (
	mutex     sync.Mutex
	AliveAddr []string
)

func CheckAlive(hosts []string, ports string, timeout int64) []string {
	var wg sync.WaitGroup
	port := ParsePort(ports)
	addrs := make(chan IPAddr, len(hosts)*len(ports))
	res := make(chan string, len(hosts)*len(ports))
	go func() {
		for found := range res {
			AliveAddr = append(AliveAddr, found)
			wg.Done()
		}
	}()

	for i := 0; i < config.Thread; i++ {
		go func() {
			for addr := range addrs {
				Connect(addr, res, timeout, &wg)
				wg.Done()
			}
		}()
	}

	for _, port := range port {
		for _, host := range hosts {
			wg.Add(1)
			addrs <- IPAddr{host, port}
		}
	}
	wg.Wait()
	close(addrs)
	close(res)
	return AliveAddr
}

func Connect(addr IPAddr, respondingHosts chan<- string, reqTimeout int64, wg *sync.WaitGroup) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", addr.Ip, addr.Port), time.Duration(reqTimeout)*time.Second)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err == nil {
		address := addr.Ip + ":" + strconv.Itoa(addr.Port)
		log.Logsuccess(address)
		respondingHosts <- address
		wg.Add(1)
	}
}
