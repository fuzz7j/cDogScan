package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"time"
)

func SnmpScan(info *config.Info) (result bool, err error) {
	params := &gosnmp.GoSNMP{
		Target:    info.Host,
		Port:      161,
		Community: "public",
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(info.Timeout) * time.Second,
	}
	err = params.Connect()
	defer params.Conn.Close()
	if err == nil {
		oids := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
		_, err2 := params.Get(oids)
		if err2 == nil {
			res := fmt.Sprintf("[SNMP]%v:%v", info.Host, 161)
			log.Logsuccess(res)
			result = true
		}
	}
	return result, nil
}