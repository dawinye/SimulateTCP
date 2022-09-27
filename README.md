# SimulateTCP (Alan Sikarov, Keith Kwan, Dawin Ye)

config.txt determines the minimum and maximum delay, as well as the number of processes and their corresponding ports. If you wish to add more ports, add additional lines in the config file. To run this simulation, open a new terminal window for every server or group of clients that you plan on using. 

Servers are ran with using 
```go run *.go (server number) s``` 
Clients are ran with using 
```go run *.go (client number) c```

Each process has one server and multiple clients, each client being able to send a message to one server. To send a message from a client, type in 
```send (server number) (message to send)``` 

We designed our code to have two files, main.go which acts as the main function and the client, and concTcpS.go which sets up our server to allow multiple clients to interact with it. We opted to not implement ```unicast_receive(source, message)``` as we received messages by reading from the connection passed in concTcpS.go. As for ```unicast_send(destination, message)```, we use the net module's dial method followed by a call to simulateDelay. The network delay "layer" is simulated by calling a goroutine which runs a for loop until a time between the min and max delay is reached, then it sends a dummy value to a channel that is initialized in unicast_send. Since reads/writes are blocking, unicast_send has no choice but to wait until simulateDelay finishes.  
