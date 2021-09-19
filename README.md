# Simple Chat Server
<br  />
<p  align="center">
  
<summary>Table of Contents</summary>

<ol>

  

<li><a  href="#about-the-task">About The Task</a></li>

<li><a  href="#how-to-run">Installation</a></li>
<li><a  href="#reference">Reference</a></li>

</ol>

# About The Task
* implement a tcp server for fragmented packets server functions:
    * send messages to all clients connected.
    * send messages to clients specified (you should tag the client)

* Test case:
(There are 10 clients, all of which can get broadcast messages
Client 1 sends a directed broadcast message to client 2)

* Solution:
  1. Setup tcp server 
  2. Start with basic start, stop, broadcast and listen to clients
  3. Setup tcp client that connects to server with partcular name 
  4. Create UI for client
# How to run
  1. Run the server that listens 5000 port 
  ```sh
go run server/cmd/main.go
```
  2. Run the client 
  ```sh
    go run ui/cmd/main.go
    ```
  3. Enter your name
  4. Send message 
  5. From server client can be tagged using
    ```sh
    tag clientname message
    ```

## References
  1. https://gist.github.com/drewolson/3950226
  2. https://github.com/marcusolsson/tui-go/blob/master/example/login/main.go
  3. https://github.com/marcusolsson/tui-go/blob/master/example/chat/main.go
  4. https://github.com/firstrow/tcp_server/blob/master/tcp_server.go
