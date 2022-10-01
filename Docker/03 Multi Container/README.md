# Multi Container
## Network
![Docker Network](https://docs.mirantis.com/containers/v3.0/dockeree-ref-arch/_images/cnm.png)  
Containers can freely make http request and response to World Wide Web using host's network. 
### Host
Docker parses host names provided with ```--add-host <custom>:<ip|host-gateway>``` option in ```docker run```
- In Windows/Mac, a special host ```host.docker.internal``` matches to ```localhost```
- If with ```docker run --net=host```, localhost can be used
- ```host-gateway``` is a special keyword for ```localhost``` in --add-host option
```js
// Only works in Windows/Mac, otherwise --add-host should be passed
mongoose.connect(
  'mongodb://host.docker.internal:27017/swfavorites',
  (err) => {
    ...
  }
);
```
### Container
If containers are in the same network, they can communicate with their container name without any open port. Docker network doesn't work with Javascript because it's using networking features provided by a browser, not OS.
- hosts can be container names if they are in the same network
- hosts can be specified using container's own ip (```docker container inspect```)
- Javascript should use ```localhost:port``` to communicate to a container
```js
// backend/app.js
$ docker network create goals
// id and password can be set for security
// for persistent data, use internal /data/db 
$ docker run --rm --name mongo --network goals \
      -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
	    -e MONGO_INITDB_ROOT_PASSWORD=secret \
      -v mongo:/data/db \  // named mount
      mongo
// port should be open for javascript access
$ docker run --rm --name goals-backend --network goals -p 90:80  node-goals 

// connect to 'mongo' container
// authSource field should be provided if the host contains id/pw
mongoose.connect(
  'mongodb://mongoadmin:secret@mongo:27014/swfavorites?authSource=admin',
  (err) => {
    ...
  }
);
```

```js
// frontend/App.js
$ docker run --name goals-frontend -p 3000:3000 --rm node-goals-front

// backend container should open port and link it to host's port 90
// host by container name won't work even in the same network (bc browser) 
try {
  const response = await fetch('http://localhost:90/goals');
  ...
}
```