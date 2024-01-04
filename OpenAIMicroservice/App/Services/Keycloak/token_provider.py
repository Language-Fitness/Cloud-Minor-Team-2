import os

import requests
import json
import base64

class TokenProvider:
    def __init__(self):
        self.endpoint = os.getenv("KEYCLOAK_HOST")
        self.client_id = os.getenv("KEYCLOAK_CLIENT_ID")
        self.client_secret = os.getenv("KEYCLOAK_CLIENT_SECRET")


    def introspect_token(self, bearer_token):
        auth_header = self.generate_basic_auth_header()

        req_body = f"token_type_hint=requesting_party_token&token={bearer_token}"
        headers = {
            "Authorization": auth_header,
            "Content-Type": "application/x-www-form-urlencoded"
        }

        try:
            response = requests.post(self.endpoint, data=req_body, headers=headers)
            response.raise_for_status()
        except requests.RequestException as e:
            return False, e

        try:
            introspection_response = response.json()
        except json.JSONDecodeError as e:
            return False, e

        active = introspection_response.get("active", False)
        return active, None

    def decode_token(self, token):
        parts = token.split(".")
        if len(parts) != 3:
            return None, ValueError("Invalid token format")

        try:
            decoded_payload = base64.urlsafe_b64decode(parts[1] + "==")
        except base64.binascii.Error as e:
            return None, e

        try:
            claims = json.loads(decoded_payload)
        except json.JSONDecodeError as e:
            return None, e

        return claims, None

    def generate_basic_auth_header(self):
        auth_string = f"{self.client_id}:{self.client_secret}"
        auth_header = base64.b64encode(auth_string.encode()).decode()
        return f"Basic {auth_header}"

