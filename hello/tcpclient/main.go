package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sent := bufio.NewReader(os.Stdin)
	for {
		readStr, err := sent.ReadString('\n')
		if err != nil {
			log.Fatal("read stdin err -> ", err)
			break
		}
		inputInfo := strings.Trim(readStr, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			break
		}
		_, send_err := conn.Write([]byte(string(inputInfo[:])))
		if send_err != nil {
			log.Fatal("send err -> ", send_err)
			break
		}

		buf := make([]byte, 1024)
		n, recv_err := conn.Read(buf[:])
		if recv_err != nil {
			log.Fatal("recv err -> ", recv_err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
