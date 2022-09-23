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

var delayArray [2]int
var defaultDelay [2]int = [2]int{10, 1000}
var procSlice = make([][]string, 0, 4) //starting cap 4

func unicast_send(destination string, message string) {
	connection, err := net.Dial("tcp", destination)
	if err != nil {
		fmt.Printf("Unable to connect to process: %s", destination)
		return
	}
	startTime := time.Now().UnixMilli()
	delay := int64(rand.Intn(defaultDelay[1]+defaultDelay[0]) - defaultDelay[0]) //so can compare w/ startTime
	//doing it this way, since context switching could happen
	for time.Now().UnixMilli()-startTime < delay {
	}
	n, err := fmt.Fprintf(connection, message)
	if err != nil || len(message) != n {
		fmt.Printf("Did not send entire message to process %s", destination)
	}
}

/*
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
	var stdSlice []string
	stdScan := bufio.NewScanner(os.Stdin)
	for stdScan.Scan() {
		stdSlice = strings.Split(stdScan.Text(), " ")
		if len(stdSlice) < 3 || stdSlice[0] != "send" {
			fmt.Println("Incorrect args. Format as 'send ID MESSAGE'")
			continue
		}
		go unicast_send(stdSlice[1], strings.Join(stdSlice[2:], ""))
	}
	fmt.Println("at the end, truly")
	time.Sleep(1000)
}
