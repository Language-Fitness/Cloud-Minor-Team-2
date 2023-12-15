import json
import os
from openai import OpenAI
from dotenv import load_dotenv
from App.Utils.Exceptions.AssistantAPIException import AssistantAPIException


class OpenAIAssistantManager:

    def __init__(self):
        # tijdelijk no touchy
        load_dotenv()
        api_key = os.getenv("OPENAI_API_KEY")

        if not api_key:
            raise ValueError("No OPENAI_API_KEY found in environment variables.")
        self.client = OpenAI(api_key=api_key)

    def create_assistant(self, assistant_json_object):
        try:
            return self.client.beta.assistants.create(
                instructions=assistant_json_object['instructions'],
                name=assistant_json_object['name'],
                tools=assistant_json_object['tools'],
                model=assistant_json_object['model']
            )
        except Exception:
            raise AssistantAPIException("Message could not be created. check fields json fields in assistant json")

    def delete_assistant(self, assistant_id):
        try:
            self.client.beta.assistants.delete(assistant_id)
        except Exception:
            raise AssistantAPIException(f"Assistant: {assistant_id} already deleted or does not exist")

    def create_thread(self):
        try:
            return self.client.beta.threads.create()
        except Exception as e:
            raise AssistantAPIException("An unexpected error occurred while creating thread.")

    def delete_thread(self, thread_id):
        try:
            self.client.beta.threads.delete(thread_id)
        except Exception:
            raise AssistantAPIException(f"Thread: {thread_id} already deleted or does not exist")

    def create_message_with_attachment(self, thread_id, message, file_id):
        try:
            return self.client.beta.threads.messages.create(
                thread_id=thread_id,
                role="user",
                content=message,
                file_ids=[file_id]
            )
        except Exception:
            raise AssistantAPIException("Message could not be created. check thread_id, message and file_id")

    def create_message(self, thread_id, message):
        try:
            return self.client.beta.threads.messages.create(
                thread_id=thread_id,
                role="user",
                content=message
            )
        except Exception:
            raise AssistantAPIException("Message could not be created. check thread_id and message")

    def create_file(self, filename, file_like_object):
        try:
            return self.client.files.create(
                file=(filename, file_like_object),
                purpose="assistants"
            )
        except Exception as e:
            raise AssistantAPIException("An unexpected error occurred while creating the file.")

    def delete_file(self, file_id):
        self.client.files.delete(file_id)

    def run_thread(self, thread_id, assistant_id):
        try:
            return self.client.beta.threads.runs.create(
                thread_id=thread_id,
                assistant_id=assistant_id
            )
        except Exception as e:
            raise AssistantAPIException("Thread could not be run. check thread_id and assistant_id")

    def retrieve_assistant(self, assistant_id):
        try:
            return self.client.beta.assistants.retrieve(
                assistant_id=assistant_id
            )

        except Exception:
            raise AssistantAPIException("No valid assistant id has been provided")

    def retrieve_messages(self, thread_id):
        try:
            return self.client.beta.threads.messages.list(
                thread_id=thread_id
            )
        except Exception:
            raise AssistantAPIException("No valid thread id has been provided")

    def load_assistant(self, assistant_file):
        try:
            with open(assistant_file, 'r') as file:
                return json.load(file)
        except FileNotFoundError:
            raise AssistantAPIException(f"Assistant file not found: {assistant_file}")
        except json.JSONDecodeError:
            raise AssistantAPIException(f"Invalid JSON format in file: {assistant_file}")
        except Exception as e:
            raise AssistantAPIException("An unexpected error occurred while loading assistant.")
