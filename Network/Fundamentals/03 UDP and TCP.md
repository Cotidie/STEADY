# UDP & TCP
## UDP (User Datagram Protocol)
![UDP](https://i.imgur.com/C9cNJRP.jpg)  
 UDP is a Layer 4 protocol that doesn't required mutual connection, thereby provides faster connection speed and smaller datagram size. Header is only 8 bytes that only contain ports, length and checksum for 2 bytes each. VPN, Video Streaming, DNS, WebRTC, etc. prefers UDP over TCP since it's faster and stateless.  
 - UDP doesn't care datagram loss or error (One way delivery)
 - **Stateless**: No ordered datagram and no flow control (doesn't request re-send) 
### TCPDUMP
```bash
# Sniff udp traffic to or from Google DNS Server (8.8.8.8)
> sudo tcpdump -ni wlp0s20f3 'udp and (src 8.8.8.8 or dst 8.8.8.8)'
# Look for ip address via Google DNS Server
> nslookup naver.com 8.8.8.8

# ToS(Type of Service) indicates priority of the packet (0 is normal)
# Source: 10.0.0.48:33191, Query ID: 41500 for identifying the response.
IP (tos 0x0, ttl 64, id 64855, offset 0, flags [none], proto UDP (17), length 55)
    10.0.0.48.33191 > 8.8.8.8.53: 41500+ A? naver.com. (27)
IP (tos 0x0, ttl 122, id 27583, offset 0, flags [none], proto UDP (17), length 119)
    8.8.8.8.53 > 10.0.0.48.33191: 41500 4/0/0 naver.com. A 223.130.200.104, naver.com. A 223.130.195.95, naver.com. A 223.130.195.200, naver.com. A 223.130.200.107 (91)
```
  
## TCP ()

## Questions
**| Why is UDP/TCP header contained in IP packet?**  
It's because that datagram/packet is formed from higher layer to lower layer, attaching headers. 