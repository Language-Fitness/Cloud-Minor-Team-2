FROM node:20-alpine

WORKDIR /app

COPY supergraph/package*.json ./

RUN npm install

COPY supergraph/ .

EXPOSE 4000

CMD [ "npm", "start" ]