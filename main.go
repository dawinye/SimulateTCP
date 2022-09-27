// INSTRUCTIONS: type 'go run *.go [process_id] [s or c]'
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func read_config() (int, int, map[string]string) {
	min_delay := 0
	max_delay := 0
	Dat, err := os.ReadFile("config.txt")
	x := string(Dat)
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
			continue
		}
		min_delay, err = strconv.Atoi(id[0])
		max_delay, err = strconv.Atoi(id[1])
		skip_first = true
	}
	return min_delay, max_delay, id_map
}

func simulateDelay(minDelay int, maxDelay int, ch chan<- string) {
	rand.Seed(time.Now().UnixNano())
	delay := float64(rand.Intn(maxDelay-minDelay)+minDelay) / float64(1000)
	for start := time.Now(); time.Since(start) <= time.Duration(delay*float64(time.Second)); {
	}
	ch <- "hello"
	close(ch)
}

func unicast_send(destination string, message string) {
	CONNECT := destination
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	min, max, _ := read_config()
	channel := make(chan string)

	go simulateDelay(min, max, channel)
	_ = <-channel

	fmt.Fprintf(c, message+"\n")
}

func main() {
	//arguments[1] = port number and [2] is whether it will be server or client
	arguments := os.Args
	_, _, id_map := read_config()
	currentTime := time.Now()
	currentTime.Format("2006-01-02 15:04:05.0000")

	if _, id_valid := id_map[arguments[1]]; !id_valid || len(arguments) < 2 || !(arguments[2] == "s" || arguments[2] == "c") {
		fmt.Println("Please format in \"go run *.go [process_id] [c or s]\". Make sure the ID is valid")
		return
	}
	//creating servers and clients in a clique, O(n^2)
	if arguments[2] == "s" {
		serverArgs := []string{id_map[arguments[1]]}
		serverSetup1(serverArgs)

	}
	if arguments[2] == "c" {
		sender_id := arguments[1]
		//this is for reading user input sourced from linode tutorial
		//format of the user input would be "send 2 message", in this scenario process 2 would be sent a message
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')

			splitted := strings.Split(text[:len(text)-1], " ")
			if splitted[0] == "STOP" {
				fmt.Println("TCP client exiting...")
				return
			}
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
			raw_message := strings.Join(splitted[2:], " ")
			message_to_send := "Received \"" + raw_message + "\" from process " + sender_id + ", system time is "
			process := "127.0.0.1:" + process_destination
			currentTime = time.Now()
			unicast_send(process, message_to_send)

			fmt.Println("Sent " + raw_message + " from process " + sender_id + ", system time is " + currentTime.Format("2006-01-02 15:04:05.0000"))
		}
	}
}
