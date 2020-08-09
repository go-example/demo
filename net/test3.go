/**
 * Created by Administrator on 13-12-9.
 */
package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * 将string类型的ip地址转换为IP对象
 */
func main() {

	name := "192.168.1.97"

	ip := net.ParseIP(name)

	if ip == nil {
		fmt.Fprintf(os.Stderr, "Err:无效的地址")
		return
	}

	fmt.Fprintf(os.Stdout, "IP: %s %s\n", ip, ip.String())
	defaultMask := ip.DefaultMask()
	fmt.Fprintf(os.Stdout, "DefaultMask: %s %s\n", defaultMask, defaultMask.String())

	ones, bits := defaultMask.Size()
	fmt.Fprintf(os.Stdout, "ones: %d bits: %d\n", ones, bits)
}