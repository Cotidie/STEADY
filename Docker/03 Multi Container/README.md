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


## Docker Compose
![docker compose](https://i.stack.imgur.com/zJxSM.png)  
Docker compose mainly solves annoying repetition of long commands/options for running containers in a multi-containerized project. Configuration for each container is specified in ```docker-compose.yaml```
- Compose doesn't replace ```Dockerfiles```
- Compose doesn't replace images or containers

### Install
| For linux: https://docs.docker.com/compose/install/linux/

### File Format
```yaml
# version of docker compose
# see: https://docs.docker.com/compose/compose-file/compose-versioning/
version: "3.8"
# Services: Container signatures for docker-compose to run
services:
  container:
    build: ...                        # build a custom image
    image: string                     # tag name
    container_name: string            # set container's name (optional)
    environment: map<string, string>  # pairs of env variables (-e)
    env_file: list<stirng>            # imports env variables from files (--env-file)
    volumes: list<string>             # volumes to mount (-v)
    networks: list<string>            # networks to use (--network)
    ports: list<string>               # ports to open (-p)
    stdin_open: bool                  # whether to open stdin of a container (-i)
    tty: bool                         # links terminal with host terminal (-t)
    depends_on: list<string>          # runs after these dependencies

# Named volumes should be specified before usage
volumes:
  data: <empty>
  ...

# Networks should be specified before usage
networks:
  goals: <empty>
  ...
```
- Named volumes, networks should be specified as top-level entries before use
- Env variables can be provided either with [key-value pairs](https://docs.docker.com/compose/environment-variables/#set-environment-variables-in-containers) or with a [env-file](https://docs.docker.com/compose/environment-variables/#the-env_file-configuration-option)
- If network is omitted, compose will auto-generate a network

### Build
```yaml
service:
  # Longer Format
  build:
    context: string(path)    # a path at which dockerfile is located
    dockerfile: string       # dockerfile to use
    args: map<string, any>   # key-value pairs to pass to dockerfile args
  # Shorter Format
  build: <path>              # a path where docker file is located 
```
If dockerfile access another folder, context should be parent folder that includes it

### Run
```bash
$ docker compose up (-d)
$ docker compose down (-v)
```
- `up -d` starts in detached mode
- `down -v` removes all named volumes used in the compose