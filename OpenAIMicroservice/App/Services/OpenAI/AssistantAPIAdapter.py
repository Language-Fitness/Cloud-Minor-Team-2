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

        return messsage.thread_id

    def Generate_Multiple_Choice_Questions(self, subject, amount_questions):
        assistant_json = self.assistant_manager.load_assistant(
            "Services/OpenAI/Assistants/MultipleChoiceQuestionAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        return messsage.thread_id

    def Generate_Explanation(self, question, given_answer):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/ExplanationAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        return messsage.thread_id

    def Generate_Answer(self, question, question_info):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/AnswerAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        thread = self.assistant_manager.create_thread()
        messsage = self.assistant_manager.create_message(thread.id, "test")

        return messsage.thread_id

    def Retrieve_Response(self, token):
        self.Decode_Token(token)

        messages = self.assistant_manager.retrieve_messages(thread_id)
        return messages

    def Encode_Token(self, token):
        return

    def Decode_Token(self):
        test = "f"
        testf = ""

        return test, testf
