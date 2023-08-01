FROM node:14

WORKDIR /app

COPY package.json /app/

RUN npm install

COPY . /app

EXPOSE 80

# For /app bind mount
VOLUME [ "/app/node_modules", "/app/temp" ]

CMD ["node", "server.js"]