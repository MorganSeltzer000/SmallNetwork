# SmallNetwork
This is a NetWork Simulation Project implemented in Go. In this project, we will establish TCP connection with simulated random delay among many devices connected 
by using a configuration file.

There are 2 major parts of the file:
- config.txt: The first line is the minimum and maximum simulated delay. The other lines contain three information: ID, IP Address, and Port Number. It looks like the following

```
min_Delay max_Delay
[ID][IP][Port Number]
.....
......
```
- main.go: This is the main executable file that contains all code. It contains the following functions.
  - main: This parses the config file and command-line arguments (see below paragraph), and calls listener and unicast_send in goroutines
  - listener: This creates a server on the port specified in the config file, and accepts the connections.
  - reader: This reads the message from the connection. This is in a separate function so listener can accept simultaneous connections without blocking to read.
  - unicast_recieve: This prints out the message, including the current time for the user to see the delay between the send time in one process and the recieve time in another
  - unicast_send: This sends the message to the specified destination with the specified message

The program can be started by running in the following
- Open 2 or more terminals: Since we are simulating a network traffic, we probably want to see connections between more than 2 nodes.
- In each terminal, run the following:
```
go run main.go [n]
```
where [n] represents the nth line of config.txt file. Windows User might see a prompted screen for Windows Security Alert, proceed by clicking Allow access.

In each terminal you opened, there should be a message prompted saying 
```
Supposed to listen on port x
Im listening!
```
where x is the associated port number that is on the line n, from previous input. 
- In any of the terminal that you have opened that has ID x, you can type the following 
```
send y "MESSAGE"
```
where "MESSAGE" is the message that you want to send to ID y. Note the input y is the [ID] associated with n that you have inputed from Step 2.
Then the following should appear
```
Will be sending to y   
Sending to [IP assoacited with y]
Sent "hello" to process [IP assoacited with y], system time is [current system time]
```
where everything inside [] is determined at the user's end.

- If correctly, on the terminal assoacited with ID y (can be identified via Port Number printout in Last step), it should appear the following
```
Accepted the connection
Received "hello" from process [x], system time is [system time]
```
where x is the terminal you send message from, and system time is the current system time.

At the end of communication, the user can terminate the communication by ```Ctrl + C```(for Windows Users).Then the following message shall appear.
```
at the end, truly
(exit status xxxxxxx) 
```

The line ```exit status xxxxxx ``` should only appear on the sending side.

# Assumptions
This code is able to handle simultaneous connections, as long as there are enough ports for each connection

We assume there is enough memory on the machine to store the config file in memory

We assume the config file is written using ipv4 (however, the program overall still works if some config lines are formatted incorrectly, they are just skipped)

We assume that the ip address on the line in the config file specified by the command-line argument is the address of the computer running that process
