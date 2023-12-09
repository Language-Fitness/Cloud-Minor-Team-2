import json
import os
from openai import OpenAI
from dotenv import load_dotenv


class OpenAIAssistantManager:

    def __init__(self):
        # tijdelijk no touchy
        load_dotenv()
        api_key = os.getenv("OPENAI_API_KEY")

        if not api_key:
            raise ValueError("No OPENAI_API_KEY found in environment variables.")
        self.client = OpenAI(api_key=api_key)

    def create_assistant(self, assistant_json_object):
        return self.client.beta.assistants.create(
            instructions=assistant_json_object['instructions'],
            name=assistant_json_object['name'],
            tools=assistant_json_object['tools'],
            model=assistant_json_object['model']
        )

    def delete_assistant(self, assistant_id):
        try:
            self.client.beta.assistants.delete(assistant_id)
        except Exception:
            print("Assistant already deleted")

    def create_thread(self):
        return self.client.beta.threads.create()

    def delete_thread(self, thread_id):
        self.client.beta.threads.delete(thread_id)

    def create_message_with_attachment(self, thread_id, message, file_id):
        return self.client.beta.threads.messages.create(
            thread_id=thread_id,
            role="user",
            content=message,
            file_ids=[file_id]
        )

    def create_message(self, thread_id, message):
        return self.client.beta.threads.messages.create(
            thread_id=thread_id,
            role="user",
            content=message
        )

    def create_file(self, filename):
        return self.client.files.create(
            file=open(filename, "rb"),
            purpose="assistants"
        )

    def delete_file(self, file_id):
        self.client.files.delete(file_id)

    def run_thread(self, thread_id, assistant_id):
        return self.client.beta.threads.runs.create(
            thread_id=thread_id,
            assistant_id=assistant_id
        )

    def retrieve_assistant(self, assistant_id):
        try:
            return self.client.beta.assistants.retrieve(
                assistant_id=assistant_id
            )

        except Exception:
            raise Exception("No valid assistant id")

    def retrieve_messages(self, thread_id):
        try:
            return self.client.beta.threads.messages.list(
                thread_id=thread_id
            )
        except Exception:
            raise Exception("No valid thread id")

    def load_assistant(self, assistant_file):
        try:
            with open(assistant_file, 'r') as file:
                return json.load(file)
        except FileNotFoundError:
            print(f"The file {assistant_file} was not found.")
        except json.JSONDecodeError:
            print(f"The file {assistant_file} could not be decoded.")
        except Exception as e:
            print(f"An error occurred: {e}")
