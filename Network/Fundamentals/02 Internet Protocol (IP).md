# Internet Protocol (IP)
## Overview
![ipv4](https://i.imgur.com/8ZY9szA.jpg)  
 IP is the address for host to host communication along with MAC address. It's the data that Layer 3 handles with packets, represented as **network.host/subnet** foramt.
- If destination ip resides in the same subnet, it only uses MAC address(L2) for communication
- If not, packet(L3) is delivered to the Default Gateway(Router)
  
## IP Packet
![IP Packet](https://i.imgur.com/YbOmhi9.png)  
- **Version**: IP Protocol's version (v4, v6)
- **IHL(HLEN)**: (Internet Header Length) Number of header's rows. Minimum 5.
- **Total Length**: Total size of a packet. 20 Bytes(Header) + user data (up to around 64kb)
- **Identification**: IP packet's unique ID, in case of fragmentation.
- **Flag/Fragmentation offset**: flags of status to fragmented or not and the offset of current fragmentation.
- **TTL**: Timeout limit. Routers decrease TTL when sending a packet.
- **Protocol**: Indicates which Protocols reside in user data (e.g. ICMP)

## ICMP (Internet Control Message Protocol)
![ICMP Datagram](https://download.huawei.com/mdl/image/download?uuid=9c21763198b14ae78941f0493a33605d)  
![ICMP code](https://static.javatpoint.com/tutorial/computer-network/images/icmp-protocol2.png)  
 ICMP is a Layer 3 protocol that adds observability over IP packets. When enabled, ICMP header and data reside in IP packet's data area. **Ping** and **Traceroute** are built upon ICMP.
 - When TTL is exceeded, a medium router sends back ICMP packet to the Source IP with error code.
 - **Ping**: `ping <destination>`. Measures latency and connection status.
 - **Traceroute**: `traceroute <destination>`. Finds routing path to the destination, by sending multiple pings with incremented TTL.