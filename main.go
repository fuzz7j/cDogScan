package main

import (
	"cDogScan/common"
	"cDogScan/config"
	"fmt"
)

func main() {
	var Info config.Info
	common.Flag(&Info)
	common.Parse(&Info)
	common.Scan(Info)
	if int(config.End) == config.TargetNum {
		fmt.Printf("已完成: %v/%v\n", config.End, config.TargetNum)
	}
}
