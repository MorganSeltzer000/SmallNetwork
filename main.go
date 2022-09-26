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
var procDict = make(map[string][2]string) //map PID to [destIP,port]

func unicast_send(myPID string, destination string, message string) {
	fmt.Println("Sending to", destination)
	connection, err := net.Dial("tcp", destination)
	defer connection.Close()
	if err != nil {
		fmt.Printf("Unable to connect to process: %s\n", destination)
		return
	}
	startTime := time.Now().UnixMilli()
	//because we're simulating network delay, we print the message actually before we send
	//and then just have an error message if the sending actually fails
	fmt.Printf("Sent \"%s\" to process %s, system time is %d\n", message, destination, startTime)
	delay := int64(rand.Intn(defaultDelay[1]+defaultDelay[0]) - defaultDelay[0]) //so can compare w/ startTime
	//doing it this way, since context switching could happen
	for time.Now().UnixMilli()-startTime < delay {
	}
	n, err := fmt.Fprintf(connection, myPID+" "+message+"\n")
	if err != nil || (len(message)+len(myPID)+2) != n {
		fmt.Println(err)
	}
}

func unicast_receive(source, message string) {
	fmt.Printf("Received \"%s\" from process %s, system time is %s \n", message, source, time.Now().UnixMilli())

}

/*
 func simulate_process(procName) {
  go unicast_send()
  go unicast_receive()
 }
*/

func listener(port string) {
	l, err := net.Listen("tcp", ":"+port)
	fmt.Println("Supposed to listen on port", port)
	if err != nil {
		panic("Unable to listen on the port")
	}
	fmt.Println("Im listening!")
	//fmt.Println(listener) //just here temporarily so the variable is used
	connection, err := l.Accept()
	fmt.Println("Accepted the connection")
	if err != nil {
		fmt.Printf("Unable to connect to listener: %s\n", l)
		return
	}

	netData, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Printf("Unable to read from: %s\n", connection)
		return
	}
	str := strings.TrimSpace(string(netData))

	source := strings.Split(str, " ")[0]
	message := strings.Split(str, " ")[1]
	unicast_receive(source, message)
	connection.Close()
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

	i := 1 //because already read one line
	myPID := "-1"
	for ; scanner.Scan(); i++ {
		tmpSlice := strings.Split(scanner.Text(), " ")
		if i == linesToRead {
			if len(tmpSlice) != 3 {
				panic("the config line for this process had an odd formatting")
			}
			//set linesToRead to len, since we might skip some lines
			//+1 since we always skip the delay line
			myPID = tmpSlice[0]
		} else if len(tmpSlice) != 3 {
			fmt.Println("a config line had odd formatting, skipping")
		}
		procDict[tmpSlice[0]] = [2]string{tmpSlice[1], tmpSlice[2]}
	}
	if i < linesToRead { //never initialized, ran out of text
		panic("line_to_read is larger than the size of the config file")
	}

	go listener(procDict[myPID][1])
	var stdSlice []string
	stdScan := bufio.NewScanner(os.Stdin)
	for stdScan.Scan() {
		stdSlice = strings.Split(stdScan.Text(), " ")
		if len(stdSlice) < 3 || stdSlice[0] != "send" {
			fmt.Println("Incorrect args. Format as 'send ID MESSAGE'")
			continue
		}
		fmt.Println("Will be sending to", stdSlice[1])
		go unicast_send(myPID, procDict[stdSlice[1]][0]+":"+procDict[stdSlice[1]][1], strings.Join(stdSlice[2:], ""))
	}
	fmt.Println("at the end, truly")
	time.Sleep(1000)
}
