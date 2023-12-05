# Dockerfile
FROM nginx:1.24
COPY ./nginx.conf /etc/nginx/nginx.conf
WORKDIR /code
COPY ./dist .
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]