# Network Concept
## Terms
- **interface**: a connection or medium between a computer and a network, such as ethernet or wifi.
- **MTU**: (Maximum Transmission Unit) a single maximum packet size that a network device can receive at one time.
- **File Descriptor**: a file descriptor is a unique identifier to a opened file or socket. It does a critical role of reading and writing data to a file/socket. Maximum number of file descriptor in memory is 4096.
- **TCB**: Transmission Control Block, something like PCB, stores connection status about a socket. It handles info such as the current sequence number and ack number.

## Congestion
 Congestion happens when a router has no more space/buffer to consume a packet and drops the packet. It alerts the congestion(packet loss) to the sender.  