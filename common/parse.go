package common

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func ParseIP(ip string, filename string) (host []string, err error) {
	if ip != "" {
		host, err = ChooseModel(ip)
	}
	if filename != "" {
		hosts, _ := ParseIPFile(filename)
		for _, ip := range hosts {
			res, _ := ChooseModel(ip)
			for _, tempip := range res {
				host = append(host, tempip)
			}
		}
	}
	host = RemoveDuplicate(host)
	return host, err
}

func ChooseModel(ip string) ([]string, error) {
	switch {
	case strings.Contains(ip, "-"):
		ips := strings.Split(ip, "-")
		if len(ips[1]) < 4 {
			return ParseIPA(ip)
		} else {
			return ParseIPB(ip)
		}
	case strings.Contains(ip,"/"):
		return ParseIPD(ip)
	default:
		testip := net.ParseIP(ip)
		if testip == nil {
			return nil, errors.New("Please check ip. for example: 127.0.0.1 | 127.0.0.1-255")
		}
		return []string{ip}, nil
	}
}

// 192.168.0.1-255

func ParseIPA(ip string) ([]string, error) {
	ips := strings.Split(ip, "-")
	var IPList []string
	if len(ips[0]) < 16 {
		if len(ips[1]) < 4 {
			ip2, err := strconv.Atoi(ips[1])
			if ip2 > 255 || err != nil {
				return nil, errors.New("IP地址输入错误")
			}
			SplitIp := strings.Split(ips[0], ".")
			ip1, err := strconv.Atoi(SplitIp[3])
			if ip1 > ip2 || err != nil {
				return nil, errors.New("IP地址输入错误")
			}
			PrefixIP := strings.Join(SplitIp[0:3], ".")
			for i := ip1; i <= ip2; i++ {
				IPList = append(IPList, PrefixIP + "." + strconv.Itoa(i))
			}
		}
	}
	return IPList, nil
}

// 192.168.0.1-192.168.0.255
// 192.168.0.1-192.168.255.255

func ParseIPB(ip string) ([]string, error) {
	var IPList []string
	ips := strings.Split(ip, "-")
	if len(ips[0]) < 16 && len(ips[1]) < 16 {
		SplitIP1 := strings.Split(ips[0], ".")
		SplitIP2 := strings.Split(ips[1], ".")
		if len(SplitIP1) != 4 || len(SplitIP2) != 4 {
			return nil, errors.New("IP地址输入错误")
		}
		if SplitIP1[0] == SplitIP2[0] && SplitIP1[1] == SplitIP2[1] && SplitIP1[2] == SplitIP2[2] {
			if SplitIP1[3] == SplitIP2[3] {
				return nil, errors.New("IP地址输入错误")
			} else if SplitIP1[3] != SplitIP2[3] {
				res, err := ParseIPA(fmt.Sprintf("%v-%v", ips[0], SplitIP2[3]))
				if err != nil {
					return nil, errors.New("IP地址输入错误")
				}
				return res, nil
			}
		}

		if SplitIP1[0] == SplitIP2[0] && SplitIP1[1] == SplitIP2[1] {
			if SplitIP1[2] != SplitIP2[2] {
				IPList, err := ParseIPC(SplitIP1, SplitIP2)
				if err != nil {
					return nil, errors.New("IP地址输入错误")
				}
				return IPList, nil
			}
		}
	}
	return IPList, nil
}

// 192.168.1.0-192.168.255.255

func ParseIPC(SplitIP1, SplitIP2 []string) ([]string, error) {
	var IPList []string
	PrefixIP := strings.Join(SplitIP1[0:2], ".")
	ip1, _ := strconv.Atoi(SplitIP1[2])
	ip2, _ := strconv.Atoi(SplitIP2[2])
	for i := ip1; i < ip2; i++ {
		for i2 := 1; i2 <= 255; i2++ {
			IPList = append(IPList, PrefixIP+"." + strconv.Itoa(i) + "." + strconv.Itoa(i2))
		}
	}
	tempips, err := ParseIPA(fmt.Sprintf("%v.%v.%v-%v", PrefixIP, SplitIP2[2], 1, SplitIP2[3]))
	if err != nil {
		return nil, errors.New("IP地址输入错误")
	}
	for _, tempip := range tempips {
		IPList = append(IPList, tempip)
	}
	return IPList, nil
}

// 192.168.0.1/24

func ParseIPD(ip string) ([]string, error) {
	var res []string
	ips := strings.Split(ip,"/")
	if ips[1] != "24" {
		return nil, errors.New("IP地址输入错误")
	}
	SplitIP := strings.Split(ips[0],".")
	res, err := ParseIPA(fmt.Sprintf("%v.%v.%v.%v-%v", SplitIP[0],SplitIP[1],SplitIP[2],1,255))
	if err != nil {
		return nil, errors.New("IP地址输入错误")
	}
	return res,nil
}

func ParsePort(port string) []int {
	var ports []int
	slices := strings.Split(port, ",")
	for _, port := range slices {
		port = strings.Trim(port, " ")
		upper := port
		if strings.Contains(port, "-") {
			ranges := strings.Split(port, "-")
			if len(ranges) < 2 {
				continue
			}
			startPort, _ := strconv.Atoi(ranges[0])
			endPort, _ := strconv.Atoi(ranges[1])
			if startPort < endPort {
				port = ranges[0]
				upper = ranges[1]
			} else {
				port = ranges[1]
				upper = ranges[0]
			}
		}
		start, _ := strconv.Atoi(port)
		end, _ := strconv.Atoi(upper)
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}
	}
	return ports
}

func ParseIPFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Read hostfile error")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var IPList []string
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			IPList = append(IPList, text)
		}
	}
	return IPList, nil
}

func RemoveDuplicate(host []string) []string {
	result := make([]string, 0, len(host))
	temp := map[string]struct{}{}
	for _, item := range host {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}