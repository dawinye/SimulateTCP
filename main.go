package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// This one would call the tcpC file and run it
func unicast_send(destination, message) {

}

func unicast_recieve(source, message) {

}

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

}
