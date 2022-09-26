package main

// INSTRUCTIONS: type 'go run *.go PORT_NUM'

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// This one would call the tcpC file and run it

func unicast_send(destination string, message string) {
	CONNECT := destination
	fmt.Println(CONNECT)
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(c, message+"\n")
}

/*
	func unicast_recieve(source string, message string) {
		listener, err := net.Listen("tcp", source)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
*/
func main() {
	//arguments[1] = port number and [2] is whether it will be server or client
	arguments := os.Args
	// if len(arguments) == 2 {
	// 	fmt.Println("Please provide port number, and whether this will be a server or client")
	// 	fmt.Println("Please do in the format: process number (s or c)")
	// 	return
	// }

	//this bottom piece takes a config file and reads it
	Dat, err := os.ReadFile("config.txt")
	x := string(Dat)
	fmt.Println(string(Dat))
	if err != nil {
		fmt.Println("err")
	}
	scanner := bufio.NewScanner(strings.NewReader(x))

	//making map for the id and port ... skip first is used so we do not add the min max to map
	skip_first := false
	id_map := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		id := strings.Split(line, " ")
		if skip_first {
			id_map[id[0]] = id[2]
		}
		skip_first = true
		//fmt.Println(scanner.Text())
	}
	fmt.Println(id_map)

	//creating servers and clients in a clique, O(n^2)
	if len(arguments) >= 2 && arguments[2] == "s" {
		serverArgs := []string{id_map[arguments[1]]}
		fmt.Println(serverArgs)
		serverSetup1(serverArgs)

	}
	if len(arguments) >= 2 && arguments[2] == "c" {
		//this is for reading user input sourced from linode tutorial
		//format of the user input would be "send 2 message", in this scenario process 2 would be sent a message
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')

			splitted := strings.Split(text[:len(text)-1], " ")
			if splitted[0] != "send" {
				fmt.Println("Error in command line...", splitted[0], "is not a valid command")
				continue
			}
			if len(splitted) < 3 {
				fmt.Println("Not enough arguements, please write as 'send ID message'")
				continue
			}
			process_destination := id_map[splitted[1]]
			// arguments[1] is the sender port id
			message := strings.Join(splitted[2:], " ") + " sent from " + arguments[1]
			fmt.Println(process_destination)
			fmt.Println(message)
			process := "127.0.0.1:" + process_destination
			unicast_send(process, message)

			// message, _ := bufio.NewReader(c).ReadString('\n')
			// fmt.Print("->: " + message)
			// if strings.TrimSpace(string(text)) == "STOP" {
			// 	fmt.Println("TCP client exiting...")
			// 	return
			// }
		}
	}

}
