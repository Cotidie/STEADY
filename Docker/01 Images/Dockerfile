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
EXPOSE 80

# Always should be the last instruction
CMD ["node", "server.js"]