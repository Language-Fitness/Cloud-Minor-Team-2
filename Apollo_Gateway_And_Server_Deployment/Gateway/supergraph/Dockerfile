FROM node:20-alpine

ENV NPM_CONFIG_CACHE=/app/.npm

WORKDIR /app

COPY package*.json ./

RUN npm install

RUN chmod -R 755 /app/.npm

COPY . .

EXPOSE 4000

CMD [ "node", "index.js" ]