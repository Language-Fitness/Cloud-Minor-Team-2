# Use the official MariaDB image from Docker Hub
FROM docker.io/mariadb:11

# Set environment variables
ENV MARIADB_DATABASE=kc
ENV MARIADB_ROOT_PASSWORD=secret
ENV MARIADB_PASSWORD=secret
ENV MARIADB_USER=sa

# Expose the default MySQL port
EXPOSE 3306

# Define the volume for data persistence
VOLUME /var/lib/mysql

# Set the working directory to /var/lib/mysql
WORKDIR /var/lib/mysql