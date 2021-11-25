package log

import (
	"cDogScan/config"
	"fmt"
	"os"
	"time"
)

var Logtime int64
var Waittime int64 = 100
var LogErrtime int64

func Logsuccess(result string) {
	Logtime = time.Now().Unix()
	fmt.Printf("%v\n",result)
	if config.NoOutput == false {
		OutputToFile(result)
	}
}

func LogError(err string)  {
	if time.Now().Unix() - Logtime > Waittime && time.Now().Unix() - LogErrtime > Waittime {
		fmt.Printf("%v\n已完成: %v/%v\n",err, config.End, config.TargetNum)
		LogErrtime = time.Now().Unix()
	}
}

func OutputToFile(res string) {
	var text = []byte(res + "\n")
	filename := "result.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Create file error.", err)
	}
	_, _ = file.Write(text)
	file.Close()
}
