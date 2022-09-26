# SmallNetwork
This is a NetWork Simulation Project implemented in Go. In this project, we will establish TCP connection with simulated random delay among many devices connected 
by using a configuration file.

There are 2 major parts of the file:
- config.txt: The configuration file contains 9 lines: the first line is the minimum and maximum simulated delay. The other lines contain three information: ID, IP Address, and Port Number. It should look like the following

```
min_Delay max_Delay
[ID][IP][Port Number]
.....
......
```
- main.go: This is the main executable file that contains all code. It contains the following functions. 

The program can be started by running in the following
- 1: Open 2 or more terminals: Since we are simulating a network traffic, we probably want to see connections between more than 2 nodes.
- 2: In each terminal, run the following:
```
go run main.go [n]
```
where [n] represents the nth config.txt file. Windows User might see a prompted screen for Windows Security Alert, proceed by clicking Allow access.

In each terminal you opened, there should be a message prompted saying 
```
Supposed to listen on port x
Im listening!
```
where x is the associated port number that is on the line n, from previous input. 
- 3: In any of the terminal that you have opened that has ID x, you can type the following 
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

- 4: If correctly, on the terminal assoacited with ID y(can be identified via Port Number printout in Last step), it should appear the following
```
Accepted the connection
Received "hello" from process [x], system time is [system time]
```
where x is the terminal you send message from, and system time is the current system time.

