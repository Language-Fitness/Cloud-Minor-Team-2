from services.keycloak.token_provider import TokenProvider
from utils.exceptions.security_exception import SecurityException


class Security:

    def __init__(self):
        self.token_provider = TokenProvider()

    def validate_token(self, token):
        is_active, error = self.token_provider.introspect_token(token)
        if error or not is_active:
            raise SecurityException("Unauthorized, Please use a valid token")

    def has_required_role(self, token, required_role):
        decoded_token, error = self.token_provider.decode_token(token)
        if error:
            raise SecurityException("Unauthorized, Please use a valid token")

        token_roles = self.extract_roles(decoded_token)
        if required_role not in token_roles:
            raise SecurityException("Unauthorized, insufficient privileges")

    @staticmethod
    def extract_token_from_header(info):
        request = info.context.get("request")
        auth_header = request.headers.get('Authorization')

        if auth_header is None:
            raise SecurityException("Authorization header is missing")

        parts = auth_header.split()

        if parts[0].lower() != 'bearer':
            raise SecurityException("Authorization header must start with Bearer")

        elif len(parts) == 1:
            raise SecurityException("Token not found")

        elif len(parts) > 2:
            raise SecurityException("Authorization header must be Bearer token")
        return parts[1]

    @staticmethod
    def extract_roles(token_data):
        roles = []
        if 'realm_access' in token_data:
            roles.extend(token_data['realm_access'].get('roles', []))
        if 'resource_access' in token_data:
            for resource in token_data['resource_access'].values():
                roles.extend(resource.get('roles', []))
        return roles
