import base64
import json
import logging
import re
from io import BytesIO
from Utils.Exceptions.AssistantAPIException import AssistantAPIException
from Services.OpenAI.openai_assistant_manager import OpenAIAssistantManager

class AssistantAPIAdapter:
    def __init__(self):
        self.assistant_manager = OpenAIAssistantManager()

    def generate_multiple_choice_questions(self, question_subject, question_level, amount_questions):
        try:
            assistant_json = self.assistant_manager.load_assistant(
                "Services/OpenAI/Assistants/MultipleChoiceQuestionAssistant.json")
            assistant = self.assistant_manager.create_assistant(assistant_json)

            request = f"onderwerp: {question_subject}, nederlands niveau: {question_level}, aantal vragen: {amount_questions}"

            thread = self.assistant_manager.create_thread()
            self.assistant_manager.create_message(thread.id, request)

            run = self.assistant_manager.run_thread(thread.id, assistant.id)
            token = self.encode_token(run.thread_id, run.assistant_id, "questions", None)
            return token

        except AssistantAPIException as e:
            logging.error(f"An error occurred in generate_multiple_choice_questions. error: {e}")
            raise

        except Exception as e:
            logging.error(f"Unexpected error occurred in generate_multiple_choice_questions. error: {e}")
            raise

    def read_multiple_choice_questions_from_file(self, file_data, filename):
        try:
            binary_data = base64.b64decode(file_data)
            binary_object = BytesIO(binary_data)

            file = self.assistant_manager.create_file(filename, binary_object)

            assistant_json = self.assistant_manager.load_assistant(
                "Services/OpenAI/Assistants/FileToMultipleChoiceQuestionAssistant.json")
            assistant = self.assistant_manager.create_assistant(assistant_json)

            request = f"bestandsnaam: {filename}, Geef de multiple choice vragen uit het bestand"

            thread = self.assistant_manager.create_thread()
            self.assistant_manager.create_message_with_attachment(thread.id, request, file.id)

            run = self.assistant_manager.run_thread(thread.id, assistant.id)
            token = self.encode_token(run.thread_id, run.assistant_id, "file-questions", file.id)
            return token

        except AssistantAPIException as e:
            logging.error(f"An error occurred in read_multiple_choice_questions_from_file. error: {e}")
            raise

        except Exception as e:
            logging.error(f"Unexpected error occurred in read_multiple_choice_questions_from_file. error: {e}")
            raise

    def generate_explanation(self, question_subject, question_text, given_answer, correct_answer):
        try:
            assistant_json = self.assistant_manager.load_assistant(
                "Services/OpenAI/Assistants/ExplanationAssistant.json")
            assistant = self.assistant_manager.create_assistant(assistant_json)

            request = f"onderwerp van de vraag: {question_subject}, de vraag zelf : {question_text}, gegeven antwoord: {given_answer}, het correcte antwoord: {correct_answer}"

            thread = self.assistant_manager.create_thread()
            self.assistant_manager.create_message(thread.id, request)

            run = self.assistant_manager.run_thread(thread.id, assistant.id)
            token = self.encode_token(run.thread_id, run.assistant_id, "explanation", None)
            return token

        except AssistantAPIException as e:
            logging.error(f"An error occurred in generate_explanation. error: {e}")
            raise

        except Exception as e:
            logging.error(f"Unexpected error occurred in generate_explanation. error: {e}")
            raise


    def retrieve_response(self, token, endpoint_id_for_check, validation_function):
        try:
            thread_id, assistant_id, file_id = self.decode_and_check_token(token, endpoint_id_for_check)
            messages = self.assistant_manager.retrieve_messages(thread_id)
            last_message_data = self.get_last_message(messages)

            json_data_dict = self.extract_json_from_message(last_message_data)

            validation_function(json_data_dict)

            self.assistant_manager.delete_assistant(assistant_id)

            if file_id:
                self.assistant_manager.delete_file(file_id)

            return json_data_dict

        except AssistantAPIException as e:
            logging.warning(f"API warning in retrieve_response: {e}")
            raise

        except Exception as e:
            logging.error(f"Unexpected API error in retrieve_response: {e}")
            raise Exception(f"An unexpected error occurred: {e}")

    def retrieve_generated_multiple_choice_questions_response(self, token):
        return self.retrieve_response(token, "questions", self.is_valid_generated_questions_json)

    def retrieve_explanation_response(self, token):
        return self.retrieve_response(token, "explanation", self.is_valid_explanation_json)

    def retrieve_multiple_choice_questions_from_file_response(self, token):
        return self.retrieve_response(token, "file-questions", self.is_valid_custom_questions_json)

    def extract_json_from_message(self, message_data):
        try:
            json_pattern = r'\{.*\}'

            # Search for JSON in the string
            matched_json = re.search(json_pattern, message_data, re.DOTALL)

            # Extracting the JSON part
            extracted_json = matched_json.group(0) if matched_json else None

            if extracted_json is None:
                raise AssistantAPIException("Response still pending, please wait.")

            json_data_dict = json.loads(extracted_json)

            return json_data_dict
        except AssistantAPIException as e:
            raise AssistantAPIException(str(e))
        except Exception as e:
            logging.error(f"Error in extract_json_from_message. error:{e}")
            raise Exception(f"Error in extract_json_from_message. error:{e}")


    def encode_token(self, thread_id, assistant_id, endpoint_id, file_id):
        try:
            json_dict = {'thread_id': thread_id, 'assistant_id': assistant_id, 'endpoint_id': endpoint_id}

            if file_id is not None and file_id.strip():
                json_dict["file_id"] = file_id

            ids_json = json.dumps(json_dict)
            ids_bytes = ids_json.encode('utf-8')
            encoded_ids = base64.b64encode(ids_bytes)

            return encoded_ids.decode('utf-8')
        except Exception as e:
            logging.error(f"Token encoding error: {e}")
            raise AssistantAPIException(f"Token encoding error: {e}")

    def decode_and_check_token(self, token, endpoint_id_for_check):

        try:
            ids_bytes = base64.b64decode(token)
            ids_json = ids_bytes.decode('utf-8')
            ids_dict = json.loads(ids_json)

            if not ids_dict.get('thread_id') or not ids_dict.get('thread_id').strip():
                raise Exception("Invalid token: 'thread_id' is missing or empty")
            if not ids_dict.get('assistant_id') or not ids_dict.get('assistant_id').strip():
                raise Exception("Invalid token: 'assistant_id' is missing or empty")
            if not ids_dict.get('endpoint_id') or not ids_dict.get('endpoint_id').strip():
                raise Exception("Invalid token: 'endpoint_id' is missing or empty")

            if endpoint_id_for_check != ids_dict.get('endpoint_id'):
                raise Exception("Not a valid endpoint_id")

            file_id = None
            if 'file_id' in ids_dict and ids_dict['file_id'].strip():
                file_id = ids_dict['file_id']

            self.assistant_manager.retrieve_messages(ids_dict['thread_id'])

            return ids_dict['thread_id'], ids_dict['assistant_id'], file_id
        except Exception as e:
            # adding loging with prometheus
            logging.error(f"Error in decode_and_check_token: {e}")
            raise AssistantAPIException("Invalid token!")

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

                return last_assistant_message_content
            else:
                raise AssistantAPIException("Response still pending, please wait.")
        else:
            raise AssistantAPIException("Please use a valid token!")

    def is_valid_generated_questions_json(self, json_data_dict):

        if 'questions' not in json_data_dict or not json_data_dict['questions']:
            logging.warning("questions field not found")
            raise AssistantAPIException("questions field not found")

        # Iterate through each question
        for question in json_data_dict['questions']:
            self.validate_question(question)

    def is_valid_custom_questions_json(self, json_data_dict):
        # Check if the JSON has 'question' key, and it's not an empty object
        if 'questions' not in json_data_dict or not json_data_dict['questions']:
            raise AssistantAPIException("No Questions have been found in file!")

        # check questions are valid
        for question in json_data_dict['questions']:
            required_fields = ['question_type', 'question_level', 'question_subject', 'question_text', 'answer_options',
                               'correct_answer']
            if any(field not in question for field in required_fields):
                raise AssistantAPIException("A required field in the file is missing. check all fields in the file!")

            # Check if certain fields are not empty or whitespace
            fields_to_check = ['question_type', 'question_level', 'question_subject', 'question_text', 'correct_answer']
            if any(not question[field].strip() for field in fields_to_check):
                raise AssistantAPIException("A required field is empty. Check all fields in the file!")

            # Check if there are at least 2 and no more than 5 answer options
            if not (2 <= len(question['answer_options']) <= 5):
                raise AssistantAPIException(
                    "Incorrect number of answer options provided in a question(2 minimum and 5 maximum)!")

            # Check if each answer option is not empty or whitespace
            if any(not option.strip() for option in question['answer_options']):
                raise AssistantAPIException("A empty answer option is provided!")

            # Check if correct_answer is one of the answer options
            if question['correct_answer'] not in question['answer_options']:
                raise AssistantAPIException("Correct answer is not among the provided answer options!")

    def validate_question(self, question):
        # Check for required fields
        required_fields = ['question_type', 'question_level', 'question_subject', 'question_text', 'answer_options',
                           'correct_answer']
        if any(field not in question for field in required_fields):
            logging.warning("A required field is missing in the question")
            raise AssistantAPIException("A required field is missing in the question")

        # Check if certain fields are not empty or whitespace
        fields_to_check = ['question_type', 'question_level', 'question_subject', 'question_text', 'correct_answer']
        if any(not question[field].strip() for field in fields_to_check):
            logging.warning("A required field is empty or contains only whitespace in a question")
            raise AssistantAPIException("A required field is empty or contains only whitespace in a question")

        # Check if there are at least 2 and no more than 5 answer options
        if not (2 <= len(question['answer_options']) <= 5):
            logging.warning("Incorrect number of answer options provided")
            raise AssistantAPIException("Incorrect number of answer options provided")

        # Check if each answer option is not empty or whitespace
        if any(not option.strip() for option in question['answer_options']):
            logging.warning("Empty or whitespace answer option found!")
            raise AssistantAPIException("Empty or whitespace answer option found!")

        # Check if correct_answer is one of the answer options
        if question['correct_answer'] not in question['answer_options']:
            logging.warning("Correct answer is not among the provided answer options")
            raise AssistantAPIException("Correct answer is not among the provided answer options")

    def is_valid_explanation_json(self, json_data_dict):
        # Check if the JSON has 'explanation' key and its not an empty object
        if 'explanation' not in json_data_dict or not json_data_dict['explanation']:
            logging.warning("explanation field not found!")
            raise AssistantAPIException("explanation field not found!")

        explanation = json_data_dict['explanation']

        # Check for required fields
        required_fields = ['info', 'tips']
        if any(field not in explanation for field in required_fields):
            logging.warning("info or tips field is missing")
            raise AssistantAPIException("info or tips field is missing")

        # Check if 'info' and 'tips' are not empty or whitespace
        if not explanation['info'].strip() or not explanation['tips'].strip():
            logging.warning("Info or tips are empty or whitespace")
            raise AssistantAPIException("Info or tips are empty or whitespace")
