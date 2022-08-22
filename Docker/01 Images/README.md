# Docker Image
![Image Layers](https://i.imgur.com/QHzlPeY.png)  
Images are read-only blueprint for containers and act as a starting point with ```docker run```. 
For useful commands, see: [docker build](#build), [docker run]()

## Image
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
```docker
# docker build -t <tag> <path>
docker build -t wonseok .
```

### Manage
