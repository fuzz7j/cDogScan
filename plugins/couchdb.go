package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func CouchDBScan(info *config.Info) (result bool, err error) {
	resp, err := http.Get(fmt.Sprintf("http://%v:%v/_utils/"))
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		if strings.Contains(string(body), "couchdb-logo") {
			res := fmt.Sprintf("CouchDB Unauthorized:%v:%v", info.Host, info.Port)
			log.Logsuccess(res)
			result = true
		}
		resp.Body.Close()
	}
	return result, err
}
