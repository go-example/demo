package main

import (
	"fmt"
	"net"
)

func main() {
	ip,_ := LocalIPv4s()
	fmt.Println(ip)
}

// LocalIPs return all non-loopback IPv4 addresses
func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
			fmt.Println(ipnet.Mask)
		}
	}

	return ips, nil
}


