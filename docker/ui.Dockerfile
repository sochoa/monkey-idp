FROM node:14
RUN mkdir -p /ui
WORKDIR /ui
COPY ui/ /ui
RUN npm install
EXPOSE 8080
CMD [ "yarn", "start" ]
