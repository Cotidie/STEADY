# Image and Container
![Image Layers](https://i.imgur.com/QHzlPeY.png)  
Images are read-only blueprint for containers and act as a starting point with ```docker run```. Images then can be shared either by attaching Dockerfile to codebase or pushing itself on **Docker Hub**.  
For useful commands, see: [docker build](#build), [docker run](#containers)

## Table of Content
  - [Images](#images)
    - [Dockerfile](#dockerfile)
    - [Build](#build)
    - [Inspect](#inspect)
  - [Containers](#containers)
    - [Run, Start, Restart](#run-start-restart)
    - [Stop, Remove](#stop-remove)
    - [Copy](#copy)
  - [Docker Hub](#docker-hub)
    - [To Upload](#to-upload)
    - [To Download](#to-download)

## Images
| Command | Option | Description |
|:-------:|:------:|-------------|
|image    | ls     | List out all images downloaded |
|image    | rmi    | Removes images from disk |
|build    | --build-arg| pass a value to a dockerfile argument |
- Images are read-only and have to be rebuilt whenever there are changes
- Images are layer-based upon ```Dockerfile``` and use cache on any layers with no changes

### Dockerfile
Dockerfile is a instructions for building custom image using ```docker build```. Each command is considered as a layer, if nothing has changed docker would use cache for faster build.
- ```EXPOSE``` is optional, as it will be specified in ```docker run -p``` option again
- ```CMD``` is only executed on each container instantiation
- ```COPY```, ```RUN npm install``` can be optimized by separating ```package.json```
- ```ENV``` variables can passed by file with --env-file option
```docker
# baseImage: on which this containaner is based
FROM node

# set current working directory inside container
WORKDIR /app

# copy <source> <dest>: which files should go inside container
# source '.' means a folder with Dockerfile
COPY . /app

# After copying files, execute this command
RUN npm install

# Arg defines a variable available inside Dockerfile
# default value can be overriden by --build_arg <var>=<val> option
ARG DEFAULT_PORT=80 

# Env defines environment variable available inside a container
# value can be overriden by -e <var>=<val>
ENV PORT $DEFULAT_PORT

# expose this port to the local computer
# docker run -p 3000:80 connects localhost:3000 to port 80 of this container
EXPOSE $DEFAULT_PORT

# Always should be the last instruction
# executed when running a container
CMD ["node", "server.js"]
```

### Build
Builds an image from ```Dockerfile```. Images can be tagged by ```-t``` option for easy execution.
- Images can be listed using ```docker image ls```
- Images use cached layers instead of duplicating, indicated size is not actual volume size.
- TAG in ```-t``` option defaults to latest
- Images can be renamed by ```docker tag <previous> REPOSITORY:TAG```
- ```.dockerignore``` excludes files specified in it, separated by new lines
```docker
# docker build -t <REPOSITORY:TAG> <path>
docker build -t wonseok:tag .
```

### Inspect
```docker image inspect <id>``` exposes basic configuration about an image; env-variables, id, layers, os etc. 
- New containers will apply ```Config```, not ```ContainerConfig```
- The number of layers can be different from the number of commands in ```Dockerfile``` as it can have base image.
```json
{
  "Id": "sha256:0c45fbfb1ecce397c2cb823d8b3a03392186eb0669780e8a2eacc8c68d3be2dd",
  "Config": {
    "Env": [
      "USER=wonseok",
      ...
    ],
    "CMD": [...],
    "WorkingDir": "/app"
  },
  "Os": "linux",
  "RootFS": {
    "Layers": [...]
  }
}
```

## Containers
| Command | Option | Description |
|:-------:| ------ | ----------- |
| ps      | -a     | list out all containers including stopped status. Useful when no need to create a new container and just **start** again. |
| run   | -d     | detach from a terminal and run it in background. |
| run   | -it    | creates a terminal in a container and listens user input from stdin. if ```-t``` is omitted current stdin will be used. |
| run   | --rm   | removes a container when exiting or stopping |
| run   | --env-file| loads env variables from file separated with new lines
| logs  | -f     | follow logs in real time. different from attach since ctrl-z/c doesn't stop the container |

```docker
docker run --rm -p 3000:80 --name goals-app IMAGE:TAG
```

### Run, Start, Restart
- **run**: Creates a container from image and starts it
- **start**: Starts any stopped(Exited) containers. Defaults to **detached** mode (ref. ```docker ps -a```)
  - ```-a``` only listens the output
  - ```-i``` allows user input

### Stop, Remove
- **stop**: Turn container's state into Exited and preserve its changes. This gives opportunity to commit it later.
- **rm**: Removes from existing containers. Loose all its state and changes.

### Copy
Normally used for getting logs from container, not good solution for updating code since running containers can be affected by it.
```docker
# Copy everything into /app folder of CONTAINER
docker cp ./. CONTAINER:/app
# Copy text.txt from CONTAINER to current folder
docker CONTAINER:/app/text.txt ./
```

## Docker Hub
Docker hub is a public repository for sharing docker images. docker uses ```push``` for uploading and ```pull``` for downloading.
### To Upload
1. Create a repositry on [Docker Hub](https://hub.docker.com/)
2. ```docker login```
3. build or rename an image to fit `username/repository` format
4. ```docker pull username/repository:tag```
### To Download
- ```docker pull username/repository:tag```
- ```docker run repository:tag``` will try to pull from Docker Hub if not in local.