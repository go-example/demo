package main

import (
	"fmt"
	"log"
	"net"
)

func main()  {
	inter, err :=  net.InterfaceByName("eth0")
	if err != nil {
		log.Fatalf("无法获取信息: %v", err)
	}

	// 获取Mac地址
	fmt.Println("Mac addr: ", inter.HardwareAddr.String())

	addrs, err := inter.Addrs()
	if err != nil {
		log.Fatalln(err)
	}

	// 获取IP地址，子网掩码
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				fmt.Println("IP:", ip.IP)
				fmt.Println("Mask:", ip.Mask)
			}
		}
	}
}
