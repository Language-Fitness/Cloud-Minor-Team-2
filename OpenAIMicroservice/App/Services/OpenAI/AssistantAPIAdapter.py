import base64
import json

from App.Services.OpenAI.OpenAIAssistantManager import OpenAIAssistantManager


class AssistantAPIAdapter:
    def __init__(self):
        self.assistant_manager = OpenAIAssistantManager()

    def Generate_Open_Answer_Questions(self, subject, amount_questions):
        assistant_json = self.assistant_manager.load_assistant(
            "Services/OpenAI/Assistants/OpenAnswerQuestionsAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.Encode_Token(run.thread_id, run.assistant_id)
        return token

    def Generate_Multiple_Choice_Questions(self, subject, amount_questions):
        assistant_json = self.assistant_manager.load_assistant(
            "Services/OpenAI/Assistants/MultipleChoiceQuestionAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.Encode_Token(run.thread_id, run.assistant_id)
        return token

    def Generate_Explanation(self, question, given_answer):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/ExplanationAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.Encode_Token(run.thread_id, run.assistant_id)
        return token

    def Generate_Answer(self, question, question_info):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/AnswerAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.Encode_Token(run.thread_id, run.assistant_id)
        return token

    def Retrieve_Response(self, token):
        self.Decode_Token(token)

        messages = self.assistant_manager.retrieve_messages(thread_id)
        return messages

    def Encode_Token(self, thread_id, assistant_id):
        # Create a dictionary of the IDs
        ids_dict = {'thread_id': thread_id, 'assistant_id': assistant_id}
        # Convert the dictionary to a JSON string
        ids_json = json.dumps(ids_dict)
        # Encode the JSON string to bytes
        ids_bytes = ids_json.encode('utf-8')
        # Base64 encode the bytes
        encoded_ids = base64.b64encode(ids_bytes)
        # Convert bytes back to a string for easy storage or transmission
        return encoded_ids.decode('utf-8')

    def Decode_Token(self, token):
        # Decode the Base64 string to bytes
        ids_bytes = base64.b64decode(token)
        # Convert bytes back to a JSON string
        ids_json = ids_bytes.decode('utf-8')
        # Convert the JSON string back to a dictionary
        return json.loads(ids_json)
