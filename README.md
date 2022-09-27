# SimulateTCP (Alan Sikarov, Keith Kwan, Dawin Ye)

config.txt determines the minimum and maximum delay, as well as the number of processes and their corresponding ports. If you wish to add more ports, add additional lines in the config file. 

To run this simulation, open a new terminal window for every server or group of clients that you plan on using. 

Servers are ran with using 
```go run *.go (server number) s``` //There should be 2 spaces

Clients are ran with using 
```go run *.go (client number) c``` //There should be 2 spaces

Each process has one server and multiple clients, each client being able to send a message to one server.

To send a message from a client, type in 
```send (server number) (message to send)``` //There should be 2 spaces
