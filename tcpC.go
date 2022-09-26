// create TCP client, sourced: https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/

// instructions:
// 1) run server main function with port number, e.g. go run tcpS.go 1234
// 2) run client main function in another terminal with port address, e.g. go run tcpC.go 127.0.0.1:1234

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func clientSetup(conn string, message string) {
	//arguments := os.Args
	// if len() != 1 {
	// 	fmt.Println("Please provide host:port.")
	// 	return
	// }

	CONNECT := conn
	fmt.Println(CONNECT)
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(c, message+"\n")
	// for {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print(">> ")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Fprintf(c, text+"\n")

	// 	message, _ := bufio.NewReader(c).ReadString('\n')
	// 	fmt.Print("->: " + message)
	// 	if strings.TrimSpace(string(text)) == "STOP" {
	// 		fmt.Println("TCP client exiting...")
	// 		return
	// 	}
	// }
}

func clientSetupCL() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
