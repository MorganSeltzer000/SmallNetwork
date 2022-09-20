package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var delayArray [2]int
var defaultDelay [2]int = [2]int{10, 1000}
var procSlice = make([][]string, 4) //starting cap 4

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

func listener() {
	fmt.Println("Im listening!")
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
		procSlice[i] = strings.Split(scanner.Text(), " ")
		if len(procSlice[i]) != 3 {
			fmt.Println("a config line had odd formatting, skipping")
			procSlice[i] = []string{"-1", "0.0.0.0", "-1"} //setting it to values that make it clear that its unreal
		}
	}
	if i < linesToRead { //never initialized, ran out of text
		panic("line_to_read is larger than the size of the config file")
	}
	fmt.Println("at the end")
	go listener()
	fmt.Println("at the end, truly")
}
