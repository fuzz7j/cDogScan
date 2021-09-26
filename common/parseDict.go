package common

import (
	"bufio"
	"cDogScan/config"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Parse(Info *config.Info)  {
	ParseUser(Info)
	ParsePassword(Info)
	ParseScantype(Info)
}

func ParseUser(Info *config.Info) {
	if Info.Username != "" {
		users := strings.Split(Info.Username, ",")
		for _, user := range users {
			Info.Usernames = append(Info.Usernames, user)
		}
		for name := range config.UserList {
			config.UserList[name] = Info.Usernames
		}
	}
	if config.Userfile != "" {
		users, err := ReadFile(config.Userfile)
		if err == nil {
			for _, user := range users {
				Info.Usernames = append(Info.Usernames, user)
			}
		}
		for name := range config.UserList {
			config.UserList[name] = Info.Usernames
		}
	}
}

func ParsePassword(Info *config.Info)  {
	if Info.Password != "" {
		pwds := strings.Split(Info.Password, ",")
		for _, pwd := range pwds {
			Info.Passwords = append(Info.Passwords, pwd)
		}
		config.Passwords = Info.Passwords
	}
	if config.Userfile != "" {
		pwds, err := ReadFile(config.Passfile)
		if err == nil {
			for _, pwd := range pwds {
				Info.Passwords = append(Info.Passwords, pwd)
			}
		}
		config.Passwords = Info.Passwords
	}
}

func ReadFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("读取字典失败")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var res []string
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			res = append(res, text)
		}
	}
	return res, nil
}

func ParseScantype(Info *config.Info) {
	var ports string
	Info.Scantypes = strings.Split(Info.Scantype,",")
	if Info.Scantype != "all" {
		for _, Info.Scantype = range Info.Scantypes {
			_, ok := config.PortList[Info.Scantype]
			if !ok {
				fmt.Printf("Error: Uknown \"%v\" Please check -m\n", Info.Scantype)
				os.Exit(1)
			}
			port, _ := config.PortList[Info.Scantype]
			ports = strings.TrimSuffix(fmt.Sprintf("%v,%v",port,ports),",")
		}
		Info.Port = ports
	}
}