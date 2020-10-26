# Peuthe Socket
Peuthe(in acheh language is same with neuputhe, ex : "Neu puthe sigam nyan bek lhe tat peugah broh putoh" is meaning "Tell someone"), Peuthe Socket is a Websocket written by golang
This socket will "tell" all client connected with Peuthe if handler triggered by another user.

## How it work??
When you send payload, peuthe socket will send for all client who connected with peuthe socket.

## How to run
- make sure you have a client for connected with peuthe socket (connect to `ws://localhost:8084/ws`).
- send payload using http handler, you can check file `send.sh` how to using handler with data.


# Contributor
- Fajrul Aulia
