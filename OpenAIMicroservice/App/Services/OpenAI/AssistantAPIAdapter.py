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
        thread_id, assistant_id = self.Decode_Token(token)

        messages = self.assistant_manager.retrieve_messages(thread_id)
        self.assistant_manager.delete_thread(thread_id)
        self.assistant_manager.delete_assistant(assistant_id)

        return messages

    def Encode_Token(self, thread_id, assistant_id):
        ids_dict = {'thread_id': thread_id, 'assistant_id': assistant_id}
        ids_json = json.dumps(ids_dict)
        ids_bytes = ids_json.encode('utf-8')
        encoded_ids = base64.b64encode(ids_bytes)

        return encoded_ids.decode('utf-8')

    def Decode_Token(self, token):
        ids_bytes = base64.b64decode(token)
        ids_json = ids_bytes.decode('utf-8')
        ids_dict = json.loads(ids_json)
        return ids_dict['thread_id'], ids_dict['assistant_id']
