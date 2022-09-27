// sourced from Linodes tutorial on creating a tcp server
// link: https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var count = 0

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			fmt.Println("debug: break from server")
			break
		}
		currentTime := time.Now()
		fmt.Println(temp + " " + currentTime.Format("2006-01-02 15:04:05.0000"))
		counter := strconv.Itoa(count) + "\n"
		c.Write([]byte(string(counter)))
	}
	c.Close()
}

func serverSetup(arguments []string) {
	if len(arguments) != 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[0]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}
