package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var delayArray [2]int
var defaultDelay [2]int = [2]int{10, 1000}
var procSlice = make([][]string, 0, 4) //starting cap 4

/*
	func unicast_send(destination string, message string) {
		connection, err := net.Dial("tcp", destination)
		if err != nil {
			fmt.Printf("Unable to connect to process: %s", destination)
		}
		fmt.Fprintf(connection, message)
		//todo
	}

	func unicast_receive(source, message string) {
		listener, err = net.Listen("tcp", PORT)
		receiveTime := time.Now().UnixMilli()
		fmt.Printf("Recieved at %d", receiveTime)
	}

	func simulate_process(procName) {
		go unicast_send()
		go unicast_receive()
	}
*/

func listener(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	fmt.Println("Supposed to listen on port", port)
	if err != nil {
		panic("Unable to listen on the port")
	}
	fmt.Println("Im listening!")
	fmt.Println(listener)
}

func main() {
	if len(os.Args) < 2 {
		panic("Incorrect args: should be ./[PROGNAME] line_to_read")
	}
	linesToRead, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Incorrect args: should be ./[PROGNAME] line_to_read")
	}
	file, err := os.Open("config.txt")
	if err != nil {
		panic("Unable to open config file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	delayLine := strings.Split(scanner.Text(), " ")
	delayArray[0], err = strconv.Atoi(delayLine[0])
	if err != nil {
		fmt.Println("Unable to read delay line. Using defaults")
		delayArray[0] = defaultDelay[0]
		fmt.Println("Default currently is:", delayArray[0])
	}
	delayArray[1], err = strconv.Atoi(delayLine[1])
	if err != nil {
		fmt.Println("Unable to read delay line. Using defaults")
		delayArray[0] = defaultDelay[1]
		fmt.Println("Default currently is:", delayArray[1])
	}

	i := 0
	for ; scanner.Scan(); i++ {
		tmpSlice := strings.Split(scanner.Text(), " ")
		if i == linesToRead {
			if len(tmpSlice) != 3 {
				panic("the config line for this process had an odd formatting")
			}
			//set linesToRead to len, since we might skip some lines
			//+1 since we always skip the delay line
			linesToRead = len(tmpSlice) + 1
		} else if len(tmpSlice) != 3 {
			fmt.Println("a config line had odd formatting, skipping")
		}
		procSlice = append(procSlice, tmpSlice)
	}
	if i < linesToRead { //never initialized, ran out of text
		panic("line_to_read is larger than the size of the config file")
	}

	fmt.Println("at the end", linesToRead, procSlice[linesToRead])
	go listener(procSlice[linesToRead][2])
	fmt.Println("at the end, truly")
	time.Sleep(1000)
}
