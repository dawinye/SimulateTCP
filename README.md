# SimulateTCP (Alan Sikarov, Keith Kwan, Dawin Ye)

### Setup
config.txt determines the minimum and maximum delay, as well as the number of processes and their corresponding ports. If you wish to add more ports, add additional lines in the config file. To run this simulation, open a new terminal window for every server or client that you plan on using. 

### Instructions
Servers are ran with using 
```go run *.go (server id) s```. 
Clients are ran with using 
```go run *.go (client id) c```.
Clients can be terminated by typing "STOP", whereas ```CTRL+C``` will stop both servers and clients that are running in terminal

Each process has one server and one client, with each client being able to send a message to any server. To send a message from a client, type in 
```send (destination server id) (message to send)``` 

### Justification
We designed our code to have two files, main.go which acts as the main function and the client, and concTcpS.go which sets up our server to allow multiple clients to interact with it. We decided to include the client into main.go because otherwise, we would have to pass in the min and max delay into the client setup function, whereas we can reuse the readConfig() function in main.go and assign the min and max delay in unicast_send() locally. We opted to not implement ```unicast_receive(source, message)``` as we receive messages by reading from the connection passed in concTcpS.go. As for ```unicast_send(destination, message)```, we use the net module's dial method followed by a call to simulateDelay. The network delay "layer" is simulated by calling a goroutine which runs a for loop until a time between the min and max delay is reached, then it sends a dummy value to a channel that is initialized in unicast_send. Since reads/writes are blocking, unicast_send has no choice but to wait until simulateDelay finishes. Lastly, we chose to have one server and client per process, rather than having n clients per process with each client being able to dial to one specific server. In our implementation, we chose to have each client be able to dial to every server, saving resources by redialing to different servers after sending a previous message. 

