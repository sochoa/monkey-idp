FROM node:14
RUN mkdir -p /src
WORKDIR /src
COPY ui/package.json ./
RUN npm install
EXPOSE 8080
CMD [ "node", "server.js" ]
