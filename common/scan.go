package common

import (
	"cDogScan/config"
	"cDogScan/plugins"
	"strings"
	"sync"
)

func Scan(Info config.Info) {
	var ch = make(chan struct{}, config.Thread)
	var wg = sync.WaitGroup{}
	hosts, _ := ParseIP(Info.Host, config.Hostfile)
	if len(hosts) > 0 {
		AliveAddr := CheckAlive(hosts, Info.Port, Info.Timeout)
		for _, targetIP := range AliveAddr {
			if !config.NoScan {
				Info.Host, Info.Port = strings.Split(targetIP, ":")[0], strings.Split(targetIP, ":")[1]
				AddScan(Info.Port, Info, ch, &wg)
			}
		}
	}
	wg.Wait()
}

func AddScan(port string, Info config.Info, ch chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		mutex.Lock()
		mutex.Unlock()
		ScanFunc(port, &Info)
		wg.Done()
		mutex.Lock()
		config.End += 1
		mutex.Unlock()
		<-ch
	}()
	ch <- struct{}{}
}

func ScanFunc(port string, Info *config.Info) {
	call := plugins.ScanFuncMap[port]
	call(Info)
}
