package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
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
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"
		c.Write([]byte(string(counter)))
	}
	c.Close()
}

func serverSetup1(arguments []string) {
	//arguments := os.Args
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