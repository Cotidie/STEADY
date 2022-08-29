# Image and Container
![Image Layers](https://i.imgur.com/QHzlPeY.png)  
Images are read-only blueprint for containers and act as a starting point with ```docker run```.  
For useful commands, see: [docker build](#build), [docker run]()

## Images
| Command | Option | Description |
|:-------:|:------:|-------------|
|image    | ls     | List out all images downloaded |
|image    | rmi    | Removes images from disk |
- Images are read-only and have to be rebuilt whenever there are changes
- Images are layer-based upon ```Dockerfile``` and use cache on any layers with no changes

### Dockerfile
Dockerfile is a instructions for building custom image using ```docker build```. Each command is considered as a layer, if nothing has changed docker would use cache for faster build.
- ```EXPOSE``` is optional, as it will be specified in ```docker run -p``` option again
- ```CMD``` is only executed on each container instantiation
- ```COPY```, ```RUN npm install``` can be optimized by separating ```package.json```
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

# expose this port to the local computer
# docker run -p 3000:80 connects localhost:3000 to port 80 of this container
EXPOSE 80

# Always should be the last instruction
CMD ["node", "server.js"]
```

### Build
Builds an image from ```Dockerfile```. Images can be tagged by ```-t``` option for easy execution.
- Images can be listed using ```docker image ls```
- Images use cached layers instead of duplicating, indicated size is not actual volume size.
```docker
# docker build -t <tag> <path>
docker build -t wonseok .
```

## Containers
| Command | Option | Description |
|:-------:| ------ | ----------- |
| ps      | -a     | list out all containers including stopped status. Useful when no need to create a new container and just **start** again. |
| run   | -d     | detach from a terminal and run it in background. |
| run   | -it    | creates a terminal in a container and listens user input from stdin. if ```-t``` is omitted current stdin will be used. |
| run   | --rm   | removes a container when exiting or stopping |
| logs  | -f     | follow logs in real time. different from attach since ctrl-z/c doesn't stop the container |


### Run, Start, Restart
- **run**: Creates a container from image and starts it
- **start**: Starts any stopped(Exited) containers. Defaults to **detached** mode (ref. ```docker ps -a```)
  - ```-a``` only listens the output
  - ```-i``` allows user input

### Stop, Remove
- **stop**: Turn container's state into Exited and preserve its changes. This gives opportunity to commit it later.
- **rm**: Removes from existing containers. Loose all its state and changes.
