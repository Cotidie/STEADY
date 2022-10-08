# RabbitMQ

## Install
### Using Docker
```docker
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.10-management
```
### Ubuntu (22.04)
| see: https://www.rabbitmq.com/download.html  
| refer to: https://www.cherryservers.com/blog/how-to-install-and-start-using-rabbitmq-on-ubuntu-22-04

```bash
# Install prerequisites
apt-get install curl gnupg apt-transport-https -y

# Add Erlang, RabbitMQ, PackageCloud signing keys
curl -1sLf "https://keys.openpgp.org/vks/v1/by-fingerprint/0A9AF2115F4687BD29803A206B73A36E6026DFCA" | sudo gpg --dearmor | sudo tee /usr/share/keyrings/com.rabbitmq.team.gpg > /dev/null
curl -1sLf "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0xf77f1eda57ebb1cc" | sudo gpg --dearmor | sudo tee /usr/share/keyrings/net.launchpad.ppa.rabbitmq.erlang.gpg > /dev/null
curl -1sLf "https://packagecloud.io/rabbitmq/rabbitmq-server/gpgkey" | sudo gpg --dearmor | sudo tee /usr/share/keyrings/io.packagecloud.rabbitmq.gpg > /dev/null

# Add repositories to apt list
sudo vim /etc/apt/sources.list.d/rabbitmq.list

deb [signed-by=/usr/share/keyrings/net.launchpad.ppa.rabbitmq.erlang.gpg] http://ppa.launchpad.net/rabbitmq/rabbitmq-erlang/ubuntu jammy main
deb-src [signed-by=/usr/share/keyrings/net.launchpad.ppa.rabbitmq.erlang.gpg] http://ppa.launchpad.net/rabbitmq/rabbitmq-erlang/ubuntu jammy main
deb [signed-by=/usr/share/keyrings/io.packagecloud.rabbitmq.gpg] https://packagecloud.io/rabbitmq/rabbitmq-server/ubuntu/ jammy main
deb-src [signed-by=/usr/share/keyrings/io.packagecloud.rabbitmq.gpg] https://packagecloud.io/rabbitmq/rabbitmq-server/ubuntu/ jammy main

# Update apt repositories
sudo apt-get update -y

# Install erlang packages
apt-get install -y erlang-base \
    erlang-asn1 erlang-crypto erlang-eldap erlang-ftp erlang-inets \
    erlang-mnesia erlang-os-mon erlang-parsetools erlang-public-key \
    erlang-runtime-tools erlang-snmp erlang-ssl \
    erlang-syntax-tools erlang-tftp erlang-tools erlang-xmerl

# Finally, install RabbitMQ
apt-get install rabbitmq-server -y --fix-missing
```

## Setting Up
### Environment Variables
| see: https://www.rabbitmq.com/configure.html#supported-environment-variables  
RabbitMQ takes environment variables for setting up basic configuration. Environment variables are prioritized over `rabbitmq-env.conf` or `rabbitmq-env-conf.bat`, and will override env-variables in the file.
```bash
# Port to open to clients (Default: 5682)
export RABBITMQ_NODE_PORT=5682
# Port to open to nodes and CLI tools
export RABBITMQ_DIST_PORT=20000+$RABBITMQ_NODE_PORT
export RABBITMQ_NODENAME=rabbit$HOSTNAME
# Database parent folder for nodes
# Each mnesia(server's database) folder must belongs to one node
export RABBITMQ_MNESIA_BASE=/var/lib/rabbitmq/mnesia/
# Dedicated message store for current node
export RABBITMQ_MNESIA_DIR=$RABBITMQ_MNESIA_BASE/$RABBITMQ_NODENAME
# Config file to read (.conf .config)
export RABBITMQ_CONFIG_FILE=/etc/rabbitmq/rabbitmq.conf
```

### Configuration
![channel](https://i.imgur.com/3XO6PZM.png)  
| See: https://www.rabbitmq.com/configure.html#config-items
```bash
# A new style format snippet. This format is used by rabbitmq.conf files.
ssl_options.cacertfile           = /path/to/ca_certificate.pem
ssl_options.certfile             = /path/to/server_certificate.pem
ssl_options.keyfile              = /path/to/server_key.pem
ssl_options.verify               = verify_peer
ssl_options.fail_if_no_peer_cert = true
```
Configuration file `rabbitmq.conf` is responsible for settings such as TCP, TLS, constraints, authentication...etc. Below are most important config entries.
| Entry | Default | Desc |
| ----- | ------- | ---- |
| cluster_name | "" | Defines the name of a origin that messages belong to |
| listeners.tcp.default | $RABBITMQ_NODE_PORT | A port number to open to clients |
| management.tcp.port | 15672 | Port to open for management web service |
| hearbeat | 60 (s) | Duration of time the TCP connection is considered to be unreachable. The smaller value will be used between server's and client's |
| frame_max | 131072 (bytes) | Maximum size of each message frame. Larger size will increase throughput, smaller size will improve latency. |
| channel_max | 0 (unlimited) | Number of channels to use to communicate. Increased size will use more memory (footprint) |

### Plugins
RabbitMQ supports various plugins such as system monitoring, message type extension, ans so on. Plugins can be enabled or disabled regardless of whether it is at runtime or not.
```bash
# To see a list of available plugins
$ sudo rabbitmq-plugins list
# To enable a plugin 
$ sudo rabbitmq-plugins enable <plugin name>
```
#### Management Web
![management-plugin](https://i.imgur.com/R3sesTw.png)
```bash
$ sudo rabbitmq-plugins enable rabbitmq-management
```
`rabbitmq-management` plugin provides visualized web page for easy management. Default user id and pw is `guest` for both. id/pw can be set in admin tab. The address to access the web board is `localhost:15682`, which port number comes from `management.tcp.port` config.

### Run Server
```bash
# Start a RabbitMQ node
$ sudo rabbitmq-server start
# View status of running rabbitmq nodes
$ sudo rabbitmqctl status
# Stop running node
$ sudo rabbitmq-server stop
```