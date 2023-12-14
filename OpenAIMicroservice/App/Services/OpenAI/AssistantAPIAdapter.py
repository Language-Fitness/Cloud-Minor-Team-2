import base64
import json

from App.Services.OpenAI.OpenAIAssistantManager import OpenAIAssistantManager


class AssistantAPIAdapter:
    def __init__(self):
        self.assistant_manager = OpenAIAssistantManager()

    # def generate_open_answer_questions(self, subject, level, amount_questions):
    #     assistant_json = self.assistant_manager.load_assistant(
    #         "Services/OpenAI/Assistants/OpenAnswerQuestionsAssistant.json")
    #     assistant = self.assistant_manager.create_assistant(assistant_json)
    #
    #     request = f"onderwerp: {subject}, nederlands niveau: {level}, aantal vragen: {amount_questions}"
    #
    #     thread = self.assistant_manager.create_thread()
    #     self.assistant_manager.create_message(thread.id, request)
    #
    #     run = self.assistant_manager.run_thread(thread.id, assistant.id)
    #     token = self.encode_token(run.thread_id, run.assistant_id)
    #     return token

    def generate_multiple_choice_questions(self, question_subject, question_level, amount_questions):
        assistant_json = self.assistant_manager.load_assistant(
            "Services/OpenAI/Assistants/MultipleChoiceQuestionAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        request = f"onderwerp: {question_subject}, nederlands niveau: {question_level}, aantal vragen: {amount_questions}"

        thread = self.assistant_manager.create_thread()
        self.assistant_manager.create_message(thread.id, request)

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token

    def read_multiple_choice_questions_from_file(self, file_data, filename):

        file = self.assistant_manager.create_file(file_data)

        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/ExplanationAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        request = f"bestandsnaam: {filename}, Geef de multiple choice vragen uit het bestand"

        thread = self.assistant_manager.create_thread()
        self.assistant_manager.create_message_with_attachment(thread.id, request, file.id)

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token

    def generate_explanation(self, question_subject, question_text, given_answer, correct_answer):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/ExplanationAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        request = f"onderwerp van de vraag: {question_subject}, de vraag zelf : {question_text}, gegeven antwoord: {given_answer}, het correcte antwoord: {correct_answer}"

        thread = self.assistant_manager.create_thread()
        self.assistant_manager.create_message(thread.id, request)

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token

    def generate_multiple_choice_answer(self, question_level, question_subject, question_text, answer_options):
        assistant_json = self.assistant_manager.load_assistant("Services/OpenAI/Assistants/AnswerAssistant.json")
        assistant = self.assistant_manager.create_assistant(assistant_json)

        request = f"onderwerp: {question_subject}, nederlands niveau: {question_level}, de vraag zelf: {question_text}, antwoord opties: {answer_options}"

        thread = self.assistant_manager.create_thread()
        self.assistant_manager.create_message(thread.id, request)

        run = self.assistant_manager.run_thread(thread.id, assistant.id)
        token = self.encode_token(run.thread_id, run.assistant_id)
        return token



    # def retrieve_open_answer_questions(self, token):
    #     try:
    #         thread_id, assistant_id = self.decode_token(token)
    #         messages = self.assistant_manager.retrieve_messages(thread_id)
    #     except Exception as e:
    #         raise Exception("Please enter a valid token!")
    #
    #     try:
    #         json_data = self.get_last_message(messages)
    #         json_data = json_data.replace('```json', '').replace('```', '')
    #
    #         json_data_dict = json.loads(json_data)
    #         json_data_dict["status"] = "success"
    #
    #         self.assistant_manager.delete_assistant(assistant_id)
    #
    #         return json_data_dict
    #
    #     except json.JSONDecodeError:
    #         raise Exception("Response still pending, please wait.")
    #
    #     except Exception as e:
    #         raise Exception(str(e))

    def retrieve_generic_response(self, token, validation_function):

        try:
            thread_id, assistant_id = self.decode_and_check_token(token)

            messages = self.assistant_manager.retrieve_messages(thread_id)

            json_data = self.get_last_message(messages)
            json_data = json_data.replace('```json', '').replace('```', '')

            json_data_dict = json.loads(json_data)

            validation_function(json_data_dict)

            self.assistant_manager.delete_assistant(assistant_id)

            # add status success after all checks went right
            json_data_dict["status"] = "success"

            return json_data_dict

        except json.JSONDecodeError:
            raise Exception("Response still pending, please wait.")

        except Exception as e:
            raise Exception(str(e))

    def retrieve_question_response(self, token):
        return self.retrieve_generic_response(token, self.is_valid_question_json)

    def retrieve_questions_response(self, token):
        return self.retrieve_generic_response(token, self.is_valid_questions_json)

    def retrieve_explanation_response(self, token):
        return self.retrieve_generic_response(token, self.is_valid_explanation_json)

    def encode_token(self, thread_id, assistant_id):
        try:
            ids_dict = {'thread_id': thread_id, 'assistant_id': assistant_id}
            ids_json = json.dumps(ids_dict)
            ids_bytes = ids_json.encode('utf-8')
            encoded_ids = base64.b64encode(ids_bytes)

            return encoded_ids.decode('utf-8')
        except Exception as e:
            raise Exception(f"Token encoding error: {e}")

    def decode_and_check_token(self, token):

        try:
            ids_bytes = base64.b64decode(token)
            ids_json = ids_bytes.decode('utf-8')
            ids_dict = json.loads(ids_json)

            if not ids_dict.get('thread_id') or not ids_dict.get('thread_id').strip():
                raise Exception("Invalid token: 'thread_id' is missing or empty")
            if not ids_dict.get('assistant_id') or not ids_dict.get('assistant_id').strip():
                raise Exception("Invalid token: 'assistant_id' is missing or empty")

            self.validate_thread_id(ids_dict['thread_id'])

            return ids_dict['thread_id'], ids_dict['assistant_id']
        except Exception:
            raise Exception("Invalid token!")

    def validate_thread_id(self, thread_id):
        manager = OpenAIAssistantManager()
        manager.retrieve_messages(thread_id=thread_id)

    def get_last_message(self, messages):
        if messages and messages.data:
            last_assistant_message = None

            # Iterate through the messages in reverse order to find the last assistant message.
            for message in reversed(messages.data):
                if message.role == 'assistant':
                    last_assistant_message = message
                    break

            if last_assistant_message:
                # Access the content of the last assistant message.
                last_assistant_message_content = last_assistant_message.content[0].text.value

                # Now you can use `last_assistant_message_content` as the last message said by the assistant.
                return last_assistant_message_content
            else:
                raise Exception("Response still pending, please wait.")
        else:
            raise Exception("Please use a valid token!")

    def is_valid_questions_json(self, json_data_dict):

        if 'questions' not in json_data_dict or not json_data_dict['questions']:
            print("questions field not found")
            raise Exception("Something went wrong, please try again")

        # Iterate through each question
        for question in json_data_dict['questions']:
            self.validate_question(question)

    def is_valid_question_json(self, json_data_dict):
        # Check if the JSON has 'question' key, and it's not an empty object
        if 'question' not in json_data_dict or not json_data_dict['question']:
            print("question field not found!")
            raise Exception("Something went wrong, please try again")

        # check if question field is valid
        self.validate_question(json_data_dict['question'])

    def validate_question(self, question):
        # Check for required fields
        required_fields = ['question_type', 'question_level', 'question_subject', 'question_text', 'answer_options',
                           'correct_answer']
        if any(field not in question for field in required_fields):
            print("A field is missing in the question!")
            raise Exception("Something went wrong, please try again")

        # Check if certain fields are not empty or whitespace
        fields_to_check = ['question_type', 'question_level', 'question_subject', 'question_text', 'correct_answer']
        if any(not question[field].strip() for field in fields_to_check):
            print("A field contains whitespace or is empty in a question!")
            raise Exception("Something went wrong, please try again")

        # Check if there are at least 2 and no more than 5 answer options
        if not (2 <= len(question['answer_options']) <= 5):
            print("Wrong amount answer options given!")
            raise Exception("Something went wrong, please try again")

        # Check if each answer option is not empty or whitespace
        if any(not option.strip() for option in question['answer_options']):
            print("Empty or whitespace answer option found!")
            raise Exception("Something went wrong, please try again")

        # Check if correct_answer is one of the answer options
        if question['correct_answer'] not in question['answer_options']:
            print("Correct answer is not found in answer options!")
            raise Exception("Something went wrong, please try again")

    def is_valid_explanation_json(self, json_data_dict):
        # Check if the JSON has 'explanation' key and its not an empty object
        if 'explanation' not in json_data_dict or not json_data_dict['explanation']:
            print("explanation field not found!")
            raise Exception("Something went wrong, please try again")

        explanation = json_data_dict['explanation']

        # Check for required fields
        required_fields = ['info', 'tips']
        if any(field not in explanation for field in required_fields):
            print("info or tips field is missing")
            raise Exception("Something went wrong, please try again")

        # Check if 'info' and 'tips' are not empty or whitespace
        if not explanation['info'].strip() or not explanation['tips'].strip():
            print("Info or tips are empty or whitespace")
            raise Exception("Something went wrong, please try again")
