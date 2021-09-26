package main

import (
	"cDogScan/common"
	"cDogScan/config"
)

func main() {
	var Info config.Info
	common.Flag(&Info)
	common.Parse(&Info)
	common.Scan(Info)
}
