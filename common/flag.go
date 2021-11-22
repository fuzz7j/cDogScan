package common

import (
	"cDogScan/config"
	"flag"
)

func Banner() {
	banner := "\n       _____               _____                 \n      |  __ \\             / ____|                \n   ___| |  | | ___   __ _| (___   ___ __ _ _ __  \n  / __| |  | |/ _ \\ / _` |\\___ \\ / __/ _` | '_ \\ \n | (__| |__| | (_) | (_| |____) | (_| (_| | | | |\n  \\___|_____/ \\___/ \\__, |_____/ \\___\\__,_|_| |_|\n                     __/ |                       \n                    |___/                        \n"
	version := "                                    Version 1.0.1\n"
	link := "               https://github.com/fuzz7j/cDogScan\n"
	print(banner,link, version)
}

func Flag(Info *config.Info) {
	Banner()
	flag.StringVar(&Info.Host, "i", "", "ip address,for example: 192.168.0.1 | 192.168.0.1-255 | 192.168.0.1-192.168.255.255 | 192.168.0.1/24")
	flag.StringVar(&Info.Port, "p", config.DefaultPorts, "port,for example: 22 | 1-65535")
	flag.StringVar(&Info.Username, "user", "", "username")
	flag.StringVar(&Info.Password, "pass", "", "password")
	flag.Int64Var(&Info.Timeout, "t", 3, "timeout")
	flag.StringVar(&Info.Scantype, "m", "all", "scan type ,for example: -m ssh | -m ssh,ftp,mysql")
	flag.IntVar(&config.Thread, "T", 600, "thread")
	flag.StringVar(&config.Hostfile, "f", "", "ip file, for example: -f ip.txt")
	flag.StringVar(&config.Userfile, "userfile", "", "username dict, for example: -userfile user.txt")
	flag.StringVar(&config.Passfile, "passfile", "", "password dict, for example: -passfile pass.txt")
	flag.BoolVar(&config.NoOutput, "nooutput", false, "not output result")
	flag.BoolVar(&config.NoScan, "no", false, "disable models, just scan port")
	flag.Parse()
}
