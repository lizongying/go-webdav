package utils

import (
	"log"
	"net"
	"strings"
)

func Lan() (lanIp string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {

			if strings.HasPrefix(i.Name, "VMnet") || strings.HasPrefix(i.Name, "VirtualBox") {
				continue
			}

			var addrs []net.Addr
			addrs, err = i.Addrs()
			if err != nil {
				log.Println(err)
				continue
			}

			for _, a := range addrs {
				if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && !ipNet.IP.IsLinkLocalUnicast() && ipNet.IP.To4() != nil {
					ip := ipNet.IP.String()
					if ip == "127.0.0.1" {
						continue
					}
					lanIp = ip
					break
				}
			}
		}
	}
	return
}
