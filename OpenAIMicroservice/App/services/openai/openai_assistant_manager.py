import json
import logging
import os
from openai import OpenAI, NotFoundError, AuthenticationError
from dotenv import load_dotenv
from utils.exceptions.assistant_api_exception import AssistantAPIException

import grpc
from App.proto.pb import OpenaiActions_pb2, OpenaiActions_pb2_grpc

class OpenAIAssistantManager:

    def __init__(self, bearer_token):
        self.client = None
        self.bearer_token = bearer_token

    def load_client(self):
        load_dotenv()
        api_key = os.getenv("OPENAI_API_KEY")

        self.get_school_from_bearer(self.bearer_token)

        if not api_key:
            raise AuthenticationError

        self.client = OpenAI(api_key=api_key)
    
    async def get_school_from_bearer(self, bearer):
        async with grpc.aio.insecure_channel("host.internal.docker:9050") as channel:
            stub = OpenaiActions_pb2_grpc.SchoolServiceStub(channel)
            response = await stub.GetKey(OpenaiActions_pb2.KeyRequest(bearerToken=bearer))
        print("Greeter client received: " + response.key)

    def create_assistant(self, assistant_json_object):
        try:
            self.load_client()

            return self.client.beta.assistants.create(
                instructions=assistant_json_object['instructions'],
                name=assistant_json_object['name'],
                tools=assistant_json_object['tools'],
                model=assistant_json_object['model']
            )
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except Exception as e:
            logging.error(f"Message could not be created. check fields json fields in assistant json. error: {e}")
            raise Exception(f"Message could not be created. check fields json fields in assistant json. error: {e}")

    def delete_assistant(self, assistant_id):
        try:
            self.load_client()

            self.client.beta.assistants.delete(assistant_id)
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except NotFoundError:
            logging.info(f"Assistant: {assistant_id} already deleted or does not exist")
        except Exception as e:
            logging.error(f"An unexpected error occurred while deleting assistant({assistant_id}). error: {e}")
            raise Exception(f"An unexpected error occurred while deleting assistant({assistant_id}). error: {e}")

    def create_thread(self):
        try:
            self.load_client()

            return self.client.beta.threads.create()
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except Exception as e:
            logging.error(f"An unexpected error occurred while creating thread. error: {e}")
            raise Exception(f"An unexpected error occurred while creating thread. error: {e}")

    def create_message_with_attachment(self, thread_id, message, file_id):
        try:
            return self.client.beta.threads.messages.create(
                thread_id=thread_id,
                role="user",
                content=message,
                file_ids=[file_id]
            )
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except Exception as e:
            logging.error(
                f"Message could not be created. thread_id: ({thread_id}), message: ({message}) and file_id: ({file_id}). error: {e}")
            raise Exception(
                f"Message could not be created. thread_id: ({thread_id}), message: ({message}) and file_id: ({file_id}). error: {e}")

    def create_message(self, thread_id, message):
        try:
            self.load_client()

            return self.client.beta.threads.messages.create(
                thread_id=thread_id,
                role="user",
                content=message
            )
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except Exception as e:
            logging.error(f"Message could not be created. check thread_id and message. error: {e}")
            raise Exception(f"Message could not be created. check thread_id and message. error: {e}")

    def create_file(self, filename, file_like_object):
        try:
            self.load_client()

            return self.client.files.create(
                file=(filename, file_like_object),
                purpose="assistants"
            )
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except Exception as e:
            logging.error(f"An unexpected error occurred while creating the file. error: {e}")
            raise Exception(f"An unexpected error occurred while creating the file. error: {e}")

    def delete_file(self, file_id):
        try:
            self.client.files.delete(file_id)
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except NotFoundError as e:
            logging.info(f"File: {file_id} already deleted or does not exist. error: {e}")
        except Exception as e:
            logging.error(f"An unexpected error occurred while deleting a file({file_id}). error: {e}")
            raise Exception(f"An unexpected error occurred while deleting a file({file_id}). error: {e}")

    def run_thread(self, thread_id, assistant_id):
        try:
            return self.client.beta.threads.runs.create(
                thread_id=thread_id,
                assistant_id=assistant_id
            )
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except Exception as e:
            raise Exception(f"Thread could not be run. check thread_id and assistant_id. error: {e}")

    def retrieve_messages(self, thread_id):
        try:
            self.load_client()

            return self.client.beta.threads.messages.list(
                thread_id=thread_id
            )
        except AuthenticationError:
            logging.error("Invalid OPENAI_API_KEY provided")
            raise AssistantAPIException(
                "Request could not be fulfilled. Please provide an valid openai API key in CMS.")
        except Exception as e:
            logging.error(f"No valid assistant id has been provided: {e}")
            raise Exception(f"No valid thread id has been provided. error: {e}")

    def load_assistant(self, assistant_file):
        try:
            with open(assistant_file, 'r') as file:
                return json.load(file)
        except FileNotFoundError as e:
            logging.error(f"Assistant file not found: {assistant_file}. error: {e}")
            raise Exception(f"Assistant file not found: {assistant_file}. error: {e}")
        except json.JSONDecodeError as e:
            logging.error(f"Invalid JSON format in file: {assistant_file}. error: {e}")
            raise Exception(f"Invalid JSON format in file: {assistant_file}. error: {e}")
        except Exception as e:
            logging.error(f"An unexpected error occurred while loading assistant. error: {e}")
            raise Exception(f"An unexpected error occurred while loading assistant. error: {e}")
