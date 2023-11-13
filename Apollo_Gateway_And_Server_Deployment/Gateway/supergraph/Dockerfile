FROM node:20-alpine

WORKDIR /app

ENV NPM_CONFIG_CACHE=/app/.npm

COPY package*.json ./

#RUN npm config set cache ./cache
RUN npm install

#RUN chgrp -R 0 /app && chmod -R g=u /app
#RUN npm cache clean --force
#RUN npm cache verify

COPY . .

EXPOSE 4000

CMD [ "npm", "start" ]