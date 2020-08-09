package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func socks5Proxy(conn net.Conn) {
	defer conn.Close()

	var b [1024]byte

	n, err := conn.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("% x", b[:n])

	conn.Write([]byte{0x05, 0x00})

	n, err = conn.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("% x", b[:n])

	var addr string
	switch b[3] {
	case 0x01:
		sip := sockIP{}
		if err := binary.Read(bytes.NewReader(b[4:n]), binary.BigEndian, &sip); err != nil {
			log.Println("请求解析错误")
			return
		}
		addr = sip.toAddr()
	case 0x03:
		host := string(b[5 : n-2])
		var port uint16
		err = binary.Read(bytes.NewReader(b[n-2:n]), binary.BigEndian, &port)
		if err != nil {
			log.Println(err)
			return
		}
		addr = fmt.Sprintf("%s:%d", host, port)
	}

	server, err := client.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	go io.Copy(server, conn)
	io.Copy(conn, server)
}

type sockIP struct {
	A, B, C, D byte
}

func (ip sockIP) toAddr() string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", ip.A, ip.B, ip.C, ip.D, ip.PORT)
}

func socks5ProxyStart() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer server.Close()
	log.Println("开始接受连接")
	for {
		client, err := server.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("一个新连接")
		go socks5Proxy(client)
	}
}

var client *ssh.Client

func main() {
	b, err := ioutil.ReadFile("/home/myml/.ssh/id_rsa")
	if err != nil {
		log.Println(err)
		return
	}
	pKey, err := ssh.ParsePrivateKey(b)
	if err != nil {
		log.Println(err)
		return
	}
	config := ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(pKey),
		},
	}
	client, err = ssh.Dial("tcp", "host:22", &config)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("连接服务器成功")
	defer client.Close()
	client.Dial("tcp", ":888")
	socks5ProxyStart()
	return
}
