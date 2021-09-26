package plugins

import (
	"cDogScan/config"
	"cDogScan/log"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"strings"
	"time"
)

func SshScan(info *config.Info) (result bool, err error) {
	for _, user := range config.UserList["ssh"] {
		for _, password := range config.Passwords {
			password = strings.Replace(password, "{user}", user, -1)
			config := &ssh.ClientConfig{
				User: user,
				Auth: []ssh.AuthMethod{
					ssh.Password(password),
				},
				Timeout: time.Duration(info.Timeout) * time.Second,
				HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
					return nil
				},
			}
			client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", info.Host, info.Port), config)
			if err == nil {
				defer client.Close()
				session, err := client.NewSession()
				if err == nil {
					defer session.Close()
					res := fmt.Sprintf("[SSH]%v:%v %v/%v", info.Host, info.Port, user, password)
					log.Logsuccess(res)
					result = true
				}
			}
		}
	}
	return result, err
}
