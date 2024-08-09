# KEIJI-BUS

# about

- Keiji-Bus is a message bus that forms a critical component of the Keiji task scheduling system. It provides a TCP-based server that handles message queuing and processing through a buffered channel.

## Features

## tcp-server

The TCP server in keiji-bus is responsible for handling incoming connections on two separate ports:

**PUSH PORT**: This port is used to receive messages. Incoming messages are pushed to the message queue by the server.

**PULL PORT**: This port is used to retrieve messages. Messages are pulled from the queue and sent to the client connected to this port.

The server starts a separate goroutine for each port to listen for incoming connections, ensuring that message handling is concurrent and non-blocking.

## queue

The message queue in keiji-bus is implemented as a buffered channel:

**Push Operation**: The server reads data from the connection on the PUSH port, unmarshals the data into a message structure, and then pushes the message to the queue.

**Pull Operation**: The server pulls one message at a time from the queue and writes it to the connection that requested it on the PULL port.

## installation

`go install https://github.com/aodr3w/keiji-bus@latest`

## Usage

`keiji bus`

you should see the following message in the terminal:

```
#% keiji-bus
2024/08/09 16:41:18 waiting for termination signal
2024/08/09 16:41:18 Server started at :8005
2024/08/09 16:41:18 Server started at :8006

```
**sending && receiving a message**

- `send`

```
# ~ % echo '{"cmd": "startTask", "taskID": "12345"}' | nc localhost 8005
OK%

```   

- `receive`

```
# ~ % nc localhost 8006 
{"cmd":"startTask","taskID":"12345"}%

```


## Integration

keiji-bus enables interprocess communication in the keiji task scheduling system. it allows the main application to send task and system control messages to the scheduler in a decoupled and reliable manner

License
keiji-bus is open-source software licensed under the MIT License. See the LICENSE file for more details.

This README now better reflects the functionality and usage of the keiji-bus based on the code provided. It provides a clear description of the system's components, how they work, and how users can interact with them.