package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("listen err", err)
	}

	defer listen.Close()
	for {
		conn, conn_err := listen.Accept()
		if conn_err != nil {
			fmt.Println("accept err", conn_err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read err", err)
			break
		}
		echo := string(buf[:n])
		fmt.Println("recv:", echo)
		conn.Write([]byte(echo))
	}
}
