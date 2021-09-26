package log

import (
	"cDogScan/config"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup
var Results = make(chan string)

func init() {
	go Savelog()
}

func Logsuccess(result string) {
	wg.Add(1)
	Results <- result
}

func Savelog() {
	for res := range Results {
		fmt.Println(res)
		if config.NoOutput == false {
			OutputToFile(res)
		}
		wg.Done()
	}

}

func OutputToFile(res string) {
	filename := "result.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Create file error.", err)
	}
	_, _ = file.WriteString(fmt.Sprintf("%v\n",res))
	file.Close()
}
