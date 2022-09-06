# Volume
![Docker Volume](https://i.imgur.com/vmVA5eC.png)  
 Docker container is an isolated operating system from host OS and has its own file system. Containers lose all information when it's stopped. As for persistent data, docker needs to share file system with host OS by ```Docker Volume```.

## Table of Contents
- [Container Layer](#container-layer)
- [Docker Volume](#docker-volume)
  - [Anonymous](#anonymous)
  - [Named](#named)
  - [Bind Mounts](#bind-mounts)
- [Errors](#errors)
  - [Read-only file system: unknown](#read-only-file-system-unknown)

## Container Layer
![Container Layer](https://i.imgur.com/n8cf8Qx.png)  
Container Layer exists on each container on top of read-only layers(commands). It stored temporary files that a process tries to modify it files, following Copy-On-Write mechanism.

## Docker Volume
Volume exists in two types: **Anonymous** and **Named**. Volumes are located on host machine, but hidden by Docker. Users don't have access to volumes' location unless they are **Bind Mounts**.
- Container file system is overwritten by volumes/mounts
- More specific volume/mount will survive
- Moundpoint of a volume can be checked with ```docker volume inspect``` (note: virtual path)

### Anonymous
```docker
### Dockerfile

# Creates an anonnymous volume
# Same as docker run -v /app/feedback
VOLUME [ "/app/feedback" ]

### bash
docker run <container path>
```
- Gives a random name to a volume created
- Directly attached to a container
- Not automatically removed when not ```run --rm``` => use ```docker volume prune```
- Normally used as preserving container folder from being overwritten

### Named
```docker
docker run -v <name>:<container path>
```
Named ones can only be specified with ```docker run -v <name>:<container path>``` and users can choose which volume to mount on that specific contaier path.
- if \<name> doesn't exist, new volume will be created
- ```docker volume ls``` to list available volumes
- Survives container shutdown and can be shared across containers

### Bind Mounts
```docker
docker run -v <absolute path>:<container path>[:access]
```
Bind mounts is a one on one mapping between host path and container path. Users then can freely access and edit the container folder. One mounted, linked container path will be overwritten by a host folder, giving a room for testing and development environment.
- When read-only, anonymous volumes inside should be specified with -v option in ```docker run```
- For production, all data should be built as a snapshot(image)

## Errors
### Read-only file system: unknown
```docker
# Dockerfile
...
# For /app bind mount
VOLUME [ "/app/node_modules", "/app/temp" ]
```
```bash
$ docker run --rm --name feedback-app -p 3000:80 -v "/home/wonseok/repositories/STEADY/Docker/02 Volumes/note-app:/app:ro" -v feedback:/app/feedback feedback-app:volume

docker: Error response from daemon: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: error mounting "/var/lib/docker/volumes/6fcda7108259a3402f6b170446ce9ceaba8a6dead8bd0e0622da67621b91840f/_data" to rootfs at "/app/node_modules": mkdir /var/lib/docker/overlay2/9649fc5ed930f4f94ab9188bbde185a72d092b0e01fb228376fe94d1c8a5a8e7/merged/app/node_modules: read-only file system: unknown.
```
- **Cause**: Anonymous volume resides in read-only bind mount
- **Solution**: Should specify in run command to propery override read-only attribute.
```bash
$ docker run --rm --name feedback-app -p 3000:80 -v "/home/wonseok/repositories/STEADY/Docker/02 Volumes/note-app:/app:ro" -v feedback:/app/feedback -v /app/node_modules feedback-app:volume
```