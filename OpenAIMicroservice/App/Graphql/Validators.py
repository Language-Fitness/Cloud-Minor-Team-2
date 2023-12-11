import base64
import json
import re
from App.Services.OpenAI.OpenAIAssistantManager import OpenAIAssistantManager
from graphql import GraphQLError


def validate_minimum_int(variable_name, value):
    if value < 1:
        raise GraphQLError(f"Amount at {variable_name} must be at least 1")


def validate_string(variable_name, value):
    if value is None or value.strip() == '':
        raise GraphQLError(f"{variable_name} must not be empty")


def validate_answer_options(answer_options):
    if not answer_options or not isinstance(answer_options, list):
        raise GraphQLError("Answer options must be a non-empty list.")

    if any(option is None or option.strip() == '' for option in answer_options):
        raise GraphQLError("Answer options must not contain empty or whitespace-only strings.")


def validate_token(token):
    if not is_valid_base64(token):
        raise GraphQLError("Invalid token!")

    try:
        ids_bytes = base64.b64decode(token)
        ids_json = ids_bytes.decode('utf-8')
        ids_dict = json.loads(ids_json)

        if not ids_dict.get('thread_id') or not ids_dict.get('thread_id').strip():
            raise GraphQLError("Invalid token: 'thread_id' is missing or empty")
        if not ids_dict.get('assistant_id') or not ids_dict.get('assistant_id').strip():
            raise GraphQLError("Invalid token: 'assistant_id' is missing or empty")

        validate_thread_id(ids_dict['thread_id'])

        return ids_dict['thread_id'], ids_dict['assistant_id']
    except Exception:
        raise GraphQLError("Invalid token!")


def validate_thread_id(thread_id):
    manager = OpenAIAssistantManager()
    manager.retrieve_messages(thread_id=thread_id)


def validate_assistant_id(assistant_id):
    manager = OpenAIAssistantManager()
    manager.retrieve_assistant(assistant_id=assistant_id)


def is_valid_base64(token):
    if not token or len(token) % 4 != 0:
        return False

    if not re.match('^[A-Za-z0-9+/]+={0,2}$', token):
        return False
    return True
