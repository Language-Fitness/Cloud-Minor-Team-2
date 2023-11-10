FROM node:20-alpine

WORKDIR /app

COPY supergraph/package*.json ./

#RUN npm config set cache ./cache
RUN npm install

#RUN chgrp -R 0 /app && chmod -R g=u /app
#RUN npm cache clean --force
#RUN npm cache verify

COPY supergraph/ .

EXPOSE 4000

CMD [ "npm", "start" ]