package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// This one would call the tcpC file and run it
/*
func unicast_send(destination string, message string) {
	connection, err := net.Dial("tcp", destination)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func unicast_recieve(source string, message string) {
	listener, err := net.Listen("tcp", source)
	if err != nil {
		fmt.Println(err)
		return
	}

}
*/
func main() {
	//we kept this if we need to run commands in terminal from within main
	cmnd := exec.Command("main.exe", "arg")
	//cmnd.Run() // and wait
	cmnd.Start()

	//this bottom piece takes a config file and reads it ... not sure if relative path works but can probably find a way to generalize it
	Dat, err := os.ReadFile("/Users/keith/DS/goScrap/config.txt")
	x := string(Dat)
	fmt.Println(string(Dat))
	if err != nil {
		fmt.Println("err")
	}
	scanner := bufio.NewScanner(strings.NewReader(x))
	for scanner.Scan() {
		line := scanner.Text()
		id := strings.Split(line, " ")
		fmt.Println(id[0])
		//fmt.Println(scanner.Text())
	}
	//this is for reading user input sourced from linode tutorial
	//format of the user input would be "send 2 message", in this scenario process 2 would be sent a message
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		fmt.Println(text)

		splitted := strings.Split(text, " ")
		if splitted[0] != "send" {
			fmt.Println("Error in command line...", splitted[0], "is not a valid command")
			continue
		}
		if len(splitted) < 3 {
			fmt.Println("Not enough arguements, please write as 'send ID message'")
			continue
		}
		process_destination := splitted[1]
		message := strings.Join(splitted[2:], "")
		fmt.Println(process_destination)
		fmt.Println(message)

		//function call to unicast send would be in here I believe
		//unicast_send(process_destination, message)

		// message, _ := bufio.NewReader(c).ReadString('\n')
		// fmt.Print("->: " + message)
		// if strings.TrimSpace(string(text)) == "STOP" {
		// 	fmt.Println("TCP client exiting...")
		// 	return
		// }
	}
}
