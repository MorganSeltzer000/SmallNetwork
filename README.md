# SmallNetwork
This is a NetWork Simulation Project implemented in Go. In this project, we will establish TCP connection with simulated random delay among many devices connected 
by using a configuration file.

There are 2 major parts of the file:
- Configuration.txt: The figuration file has 9 lines: the first line represents the minimum and maximum simulated delay. The rest of the file, per line, should
be understood as 3 parts: [ID][IP Address][Port Number] where each part was seperated by an empty space
- main.go
