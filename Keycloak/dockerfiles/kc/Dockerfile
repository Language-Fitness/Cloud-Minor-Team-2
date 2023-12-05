# Use the official Keycloak image from Quay.io
FROM quay.io/keycloak/keycloak:23.0.0

# Set environment variables
ENV KC_HOSTNAME=localhost
ENV KC_HOSTNAME_PORT=8888
ENV KC_HOSTNAME_STRICT_BACKCHANNEL="true"
ENV KC_DB=mariadb
ENV KC_DB_URL="jdbc:mariadb://mysql:3306/kc?characterEncoding=UTF-8"
ENV KC_DB_USERNAME=sa
ENV KC_DB_PASSWORD=secret
ENV KEYCLOAK_ADMIN=admin
ENV KEYCLOAK_ADMIN_PASSWORD=admin
ENV KC_HEALTH_ENABLED=true
ENV KC_LOG_LEVEL=info
ENV KC_HOSTNAME_STRICT_HTTPS=false
ENV HOSTNAME_STRICT=false
ENV KC_HOSTNAME_STRICT=false
ENV KC_HTTP_ENABLED=true
ENV KC_LOG=console
ENV KC_PROXY=none

# Set the working directory to /opt/jboss/keycloak
WORKDIR /opt/jboss/keycloak

# Expose the ports
EXPOSE 8888

# Define healthcheck
HEALTHCHECK --interval=15s --timeout=2s --retries=15 CMD curl -f http://localhost:8888/health/ready || exit 1

# Define the command to run on startup
CMD ["start-dev", "--http-port=8888"]
