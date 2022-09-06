# Volume
![Docker Volume](https://i.imgur.com/vmVA5eC.png)  
 Docker container is an isolated operating system from host OS and has its own file system. Containers lose all information when it's stopped. As for persistent data, docker needs to share file system with host OS by ```Docker Volume```.

## Table of Contents
  - [Container Layer](#container-layer)
  - [Docker Volume](#docker-volume)
    - [Anonymous](#anonymous)
    - [Named](#named)
    - [Bind Mounts](#bind-mounts)

## Container Layer
![Container Layer](https://i.imgur.com/n8cf8Qx.png)  
Container Layer exists on each container on top of read-only layers(commands). It stored temporary files that a process tries to modify it files, following Copy-On-Write mechanism.

## Docker Volume
Volume exists in two types: **Anonymous** and **Named**. Volumes are located on host machine, but hidden by Docker. Users don't have access to volumes' location unless they are **Bind Mounts**.
- Container file system is overwritten by volumes/mounts
- More specific volume/mount will survive

### Anonymous
- Gives a random name to a volume created
- Directly attached to a container
- Not automatically removed when not ```run --rm``` => use ```docker volume prune```
- Normally used as preserving container folder from being overwritten
```docker
# Dockerfile

# Creates an anonnymous volume
# Same as docker run -v /app/feedback
VOLUME [ "/app/feedback" ]
```

### Named
Named ones can only be specified with ```docker run -v <name>:<container path>``` and users can choose which volume to mount on that specific contaier path.
- if \<name> doesn't exist, new volume will be created
- ```docker volume ls``` to list available volumes

### Bind Mounts

