# Multi Container
## Network
![Docker Network](https://docs.mirantis.com/containers/v3.0/dockeree-ref-arch/_images/cnm.png)  
Containers can freely make http request and response to World Wide Web using host's network. 
### Host
Docker parses a special host ```host.docker.internal``` into the host machine's address.
```js
mongoose.connect(
  'mongodb://host.docker.internal:27017/swfavorites',
  (err) => {
    ...
  }
);
```
### Container
- hosts can be container names if they are in the same network
- hosts can be specified using container's own ip (```docker container inspect```)
```js
$ docker network create <network name>
$ docker run --network <network name> <container>

// connect to 'mongo' container
mongoose.connect(
  'mongodb://mongo:27014/swfavorites',
  (err) => {
    ...
  }
);
```