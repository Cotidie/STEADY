# Docker

## Overview
### What is Docker 
Docker is a tool for creating and managing **containers**. A container is a unit of software that has code and dependencies. The same container always yields the exact same behavior and results. Docker simplifies the creation and management of such containers.
- can build and test in exactly the same environment
- can easily switch between projects
- a container(image) has a package of tools, env-settings and other metadata

### Containers vs VMs
![docker vs vm](https://i.imgur.com/Rr7enwF.jpg)  
- VM has OS overheads unlike Docker containers
- Docker uses OS built-in support for encapsulating containers
- Docker offers **images** as a sharing, re-building method
So docker containers can be said lightweigh virtual OS that uses built-in support of the host OS.

### Tools
- [Docker Desktop](https://docs.docker.com/desktop/): A virtual machine for Windows and MacOS to use Docker Engine
  - Linux doesn't need Docker Desktop because Linux provides all the supports for Docker Engine
- [Docker Hub](https://docs.docker.com/docker-hub/): A cloud for sharing container images
- [Docker Compose](https://docs.docker.com/compose/): An extension for managing more complex containers like multi-container projects