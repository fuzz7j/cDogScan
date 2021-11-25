package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ElasticScan(info *config.Info) (result bool, err error) {
	resp, err := http.Get(fmt.Sprintf("http://%v:%v/_cat", info.Host, info.Port))
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		if strings.Contains(string(body), "/_cat/master") {
			res := fmt.Sprintf("Elastic Unauthorized:%v:%v", info.Host, info.Port)
			log.Logsuccess(res)
			result = true
		}
		resp.Body.Close()
	}
	return result, err
}
