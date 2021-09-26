package common

import (
	"bufio"
	"errors"
	"net"
	"os"
	"strconv"
	"strings"
)

/*
192.168.0.1
192.168.0.1/24
192.168.0.1-192.168.255.255
*/

func ParseIP(ip string, filename string) (host []string, err error) {
	if ip != "" {
		host, err = ChooseModel(ip)
	}
	if filename != "" {
		host, err = ParseIpFile(filename)
	}
	host = RemoveDuplicate(host)
	return host, err
}

func ChooseModel(ip string) ([]string, error) {
	switch {
	case strings.Contains(ip, "-"):
		return ParseIPA(ip)
	default:
		testip := net.ParseIP(ip)
		if testip == nil {
			return nil, errors.New("Please check ip. for example: 127.0.0.1 | 127.0.0.1-255")
		}
		return []string{ip}, nil
	}
}

func ParseIPA(ip string) ([]string, error) {
	ips := strings.Split(ip, "-")
	var IPList []string
	if len(ips[1]) < 4 {
		ip2, err := strconv.Atoi(ips[1])
		if ip2 > 255 || err != nil {
			return nil, errors.New("IP地址输入有误")
		}
		SplitIP := strings.Split(ips[0], ".")
		ip1, err := strconv.Atoi(SplitIP[3])
		if ip1 > ip2 || err != nil {
			return nil, errors.New("IP地址输入有误")
		}
		PrefixIP := strings.Join(SplitIP[0:3], ".")
		for i := ip1; i <= ip2; i++ {
			IPList = append(IPList, PrefixIP+"."+strconv.Itoa(i))
		}
	}
	return IPList, nil
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

func ParseIpFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("read hostfile error")
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