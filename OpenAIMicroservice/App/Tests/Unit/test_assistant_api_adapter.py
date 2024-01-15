import base64
import json
import unittest
from unittest.mock import patch, MagicMock
from openai.pagination import SyncCursorPage
from openai.types.beta.threads import ThreadMessage, MessageContentText
from openai.types.beta.threads.message_content_text import Text
from services.openai.assistant_api_adapter import AssistantAPIAdapter
from utils.exceptions.assistant_api_exception import AssistantAPIException
import logging


class TestAssistantAPIAdapter(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        logging.disable(logging.CRITICAL)

    @classmethod
    def tearDownClass(cls):
        logging.disable(logging.NOTSET)

    @patch.dict('os.environ', {'OPENAI_API_KEY': 'mocked_api_key'})
    def setUp(self):
        self.assistantAPIAdapter = AssistantAPIAdapter("mocked_bearer_token")

    # ALL TESTS FOR: GENERATE_MULTIPLE_CHOICE_QUESTIONS

    # SUCCESS TESTS

    # 1. Successfully running the method
    @patch('services.openai.assistant_api_adapter.AssistantAPIAdapter.encode_token')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    def test_generate_multiple_choice_questions_success(
            self, mock_load_assistant, mock_create_assistant, mock_create_thread,
            mock_create_message, mock_run_thread, mock_encode_token
    ):
        # Arrange
        assistant_json = {
            "name": "MultipleChoiceQuestionAssistant",
            "instructions": "Make multiple choice questions please? :)",
            "tools": [],
            "model": "gpt-4-1106-preview"
        }
        mock_load_assistant.return_value = assistant_json
        mock_create_assistant.return_value = MagicMock(id="asst_jAMIS9Va4KP4F3yLuXvhana9")
        mock_create_thread.return_value = MagicMock(id="thread_aYIPRRW5BKCamyQhjm3XpiQm")
        mock_run_thread.return_value = MagicMock(thread_id="thread_aYIPRRW5BKCamyQhjm3XpiQm",
                                                 assistant_id="asst_jAMIS9Va4KP4F3yLuXvhana9")
        expected_token_dict = {
            'thread_id': 'thread_aYIPRRW5BKCamyQhjm3XpiQm',
            'assistant_id': 'asst_jAMIS9Va4KP4F3yLuXvhana9',
            'endpoint_id': 'questions'
        }
        expected_encoded_token = base64.b64encode(json.dumps(expected_token_dict).encode()).decode()

        mock_encode_token.return_value = expected_encoded_token

        # Act
        result = self.assistantAPIAdapter.generate_multiple_choice_questions("Woordenschat", "B2", 5)

        # Assert
        mock_load_assistant.assert_called_once()
        mock_create_assistant.assert_called_once_with(assistant_json)
        mock_create_thread.assert_called_once()
        mock_create_message.assert_called_once()
        mock_run_thread.assert_called_once()
        mock_encode_token.assert_called_once_with(
            'thread_aYIPRRW5BKCamyQhjm3XpiQm',
            'asst_jAMIS9Va4KP4F3yLuXvhana9',
            'questions',
            None
        )
        self.assertEqual(result, expected_encoded_token)

    # FAIL TESTS

    # 1.1 Failure in Loading Assistant
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    def test_generate_multiple_choice__load_assistant_failure_exception(self, mock_load_assistant):
        mock_load_assistant.side_effect = Exception("Load assistant failure")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.2 Failure in Creating Assistant (Exception)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    def test_generate_multiple_choice__create_assistant_failure_exception(self, mock_create_assistant,
                                                                          mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.side_effect = Exception(
            "Message could not be created. check fields json fields in assistant json.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.3 Failure in Creating Assistant (AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    def test_generate_multiple_choice__create_assistant_failure_assistant_api_exception(self, mock_create_assistant,
                                                                                        mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.4 Failure in Creating Thread (Exception)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    def test_create_thread_failure_exception(self, mock_create_thread, mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.side_effect = Exception("An unexpected error occurred while creating thread.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.5 Failure in Creating Thread (AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    def test_generate_multiple_choice__create_thread_failure_assistant_api_exception(self, mock_create_thread,
                                                                                     mock_create_assistant,
                                                                                     mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.6 Failure in Creating Message(AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    def test_generate_multiple_choice__create_message_failure_assistant_api_exception(self, mock_create_message,
                                                                                      mock_create_thread,
                                                                                      mock_create_assistant,
                                                                                      mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.7 Failure in Creating Message(Exception)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    def test_generate_multiple_choice__create_message_failure_exception(self, mock_create_message, mock_create_thread,
                                                                        mock_create_assistant,
                                                                        mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.side_effect = Exception("Message could not be created. check thread_id and message.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.8 Failure in Running Thread(AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    def test_generate_multiple_choice__run_thread_failure_assistant_api_exception(self, mock_run_thread,
                                                                                  mock_create_message,
                                                                                  mock_create_thread,
                                                                                  mock_create_assistant,
                                                                                  mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.return_value = MagicMock()
        mock_run_thread.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.9 Failure in Running Thread
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    def test_generate_multiple_choice__run_thread_failure_exception(self, mock_run_thread, mock_create_message,
                                                                    mock_create_thread,
                                                                    mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.return_value = MagicMock()
        mock_run_thread.side_effect = Exception("Thread could not be run. check thread_id and assistant_id.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.10 Failure in Token Encoding
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    @patch('services.openai.assistant_api_adapter.AssistantAPIAdapter.encode_token')
    def test_generate_multiple_choice__encode_token_failure(self, mock_encode_token, mock_run_thread,
                                                            mock_create_message, mock_create_thread,
                                                            mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.return_value = MagicMock()
        mock_run_thread.return_value = MagicMock()

        mock_encode_token.side_effect = AssistantAPIException("Token encoding error")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # ALL TESTS FOR: READ_MULTIPLE_CHOICE_QUESTIONS_FROM_FILE

    # SUCCESS TESTS

    # 2. Successfully running the method
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_file')
    @patch('services.openai.assistant_api_adapter.AssistantAPIAdapter.encode_token')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message_with_attachment')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    def test_read_multiple_choice_questions_from_file_success(
            self, mock_load_assistant, mock_create_assistant, mock_create_thread,
            mock_create_message_with_attachment, mock_run_thread, mock_encode_token, mock_create_file
    ):
        # Arrange
        assistant_json = {
            "name": "FileToMultipleChoiceQuestionAssistant",
            "instructions": "Retrieve the dutch multiple choice questions from the given file? :)",
            "tools": [{"type": "retrieval"}, {"type": "code_interpreter"}],
            "model": "gpt-4-1106-preview"
        }
        mock_create_file.return_value = MagicMock(id="file-4h5h0tPUCjZd2TaAcZgO7LkY")
        mock_load_assistant.return_value = assistant_json
        mock_create_assistant.return_value = MagicMock(id="asst_jAMIS9Va4KP4F3yLuXvhana9")
        mock_create_thread.return_value = MagicMock(id="thread_aYIPRRW5BKCamyQhjm3XpiQm")
        mock_run_thread.return_value = MagicMock(thread_id="thread_aYIPRRW5BKCamyQhjm3XpiQm",
                                                 assistant_id="asst_jAMIS9Va4KP4F3yLuXvhana9")
        expected_token_dict = {
            'thread_id': 'thread_aYIPRRW5BKCamyQhjm3XpiQm',
            'assistant_id': 'asst_jAMIS9Va4KP4F3yLuXvhana9',
            'endpoint_id': 'file-questions',
            'file_id': 'file-4h5h0tPUCjZd2TaAcZgO7LkY'
        }
        expected_encoded_token = base64.b64encode(json.dumps(expected_token_dict).encode()).decode()

        mock_encode_token.return_value = expected_encoded_token

        # Act
        result = self.assistantAPIAdapter.read_multiple_choice_questions_from_file(
            "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",
            "test.docx")

        # Assert
        mock_create_file.assert_called_once()
        mock_load_assistant.assert_called_once()
        mock_create_assistant.assert_called_once_with(assistant_json)
        mock_create_thread.assert_called_once()
        mock_create_message_with_attachment.assert_called_once()
        mock_run_thread.assert_called_once()
        mock_encode_token.assert_called_once_with(
            'thread_aYIPRRW5BKCamyQhjm3XpiQm',
            'asst_jAMIS9Va4KP4F3yLuXvhana9',
            'file-questions',
            'file-4h5h0tPUCjZd2TaAcZgO7LkY'
        )

        self.assertEqual(result, expected_encoded_token)

    # FAIL TESTS

    # 2.1 Failure in Loading Assistant

    # ALL TESTS FOR: GENERATE_EXPLANATION

    # SUCCESS TESTS

    # 3. Successfully running the method

    @patch('services.openai.assistant_api_adapter.AssistantAPIAdapter.encode_token')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    def test_generate_explanation_success(
            self, mock_load_assistant, mock_create_assistant, mock_create_thread,
            mock_create_message, mock_run_thread, mock_encode_token
    ):
        # Arrange
        assistant_json = {
            "name": "ExplanationAssistant",
            "instructions": "give explanation please? :)",
            "tools": [],
            "model": "gpt-4-1106-preview"
        }
        mock_load_assistant.return_value = assistant_json
        mock_create_assistant.return_value = MagicMock(id="asst_jAMIS9Va4KP4F3yLuXvhana9")
        mock_create_thread.return_value = MagicMock(id="thread_aYIPRRW5BKCamyQhjm3XpiQm")
        mock_run_thread.return_value = MagicMock(thread_id="thread_aYIPRRW5BKCamyQhjm3XpiQm",
                                                 assistant_id="asst_jAMIS9Va4KP4F3yLuXvhana9")
        expected_token_dict = {
            'thread_id': 'thread_aYIPRRW5BKCamyQhjm3XpiQm',
            'assistant_id': 'asst_jAMIS9Va4KP4F3yLuXvhana9',
            'endpoint_id': 'explanation'
        }
        expected_encoded_token = base64.b64encode(json.dumps(expected_token_dict).encode()).decode()
        mock_encode_token.return_value = expected_encoded_token

        # Act
        result = self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                               "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                               "je bent op zoek naar een specifiek zeilschip",
                                                               "je let goed op om te zorgen dat alles goed gaat")

        # Assert
        mock_load_assistant.assert_called_once()
        mock_create_assistant.assert_called_once_with(assistant_json)
        mock_create_thread.assert_called_once()
        mock_create_message.assert_called_once()
        mock_run_thread.assert_called_once()
        mock_encode_token.assert_called_once_with(
            'thread_aYIPRRW5BKCamyQhjm3XpiQm',
            'asst_jAMIS9Va4KP4F3yLuXvhana9',
            'explanation',
            None
        )
        self.assertEqual(result, expected_encoded_token)

    # FAIL TESTS

    # 3.1 Failure in Loading Assistant
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    def test_generate_explanation__load_assistant_failure_exception(self, mock_load_assistant):
        mock_load_assistant.side_effect = Exception("Load assistant failure")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.2 Failure in Creating Assistant (Exception)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    def test_generate_explanation__create_assistant_failure_exception(self, mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.side_effect = Exception(
            "Message could not be created. check fields json fields in assistant json.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.3 Failure in Creating Assistant (AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    def test_generate_explanation__create_assistant_failure_assistant_api_exception(self, mock_create_assistant,
                                                                                    mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.4 Failure in Creating Thread (Exception)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    def test_generate_explanation__create_thread_failure_exception(self, mock_create_thread, mock_create_assistant,
                                                                   mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.side_effect = Exception("An unexpected error occurred while creating thread.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.5 Failure in Creating Thread (AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    def test_generate_explanation__create_thread_failure_assistant_api_exception(self, mock_create_thread,
                                                                                 mock_create_assistant,
                                                                                 mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.6 Failure in Creating Message(AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    def test_generate_explanation__create_message_failure_assistant_api_exception(self, mock_create_message,
                                                                                  mock_create_thread,
                                                                                  mock_create_assistant,
                                                                                  mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.7 Failure in Creating Message(Exception)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    def test_generate_explanation__create_message_failure_exception(self, mock_create_message, mock_create_thread,
                                                                    mock_create_assistant,
                                                                    mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.side_effect = Exception("Message could not be created. check thread_id and message.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.8 Failure in Running Thread(AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    def test_generate_explanation__run_thread_failure_assistant_api_exception(self, mock_run_thread,
                                                                              mock_create_message, mock_create_thread,
                                                                              mock_create_assistant,
                                                                              mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.return_value = MagicMock()
        mock_run_thread.side_effect = AssistantAPIException(
            "Request could not be fulfilled. Please provide an valid OpenAI API key in CMS.")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.9 Failure in Running Thread
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    def test_generate_explanation__run_thread_failure_exception(self, mock_run_thread, mock_create_message,
                                                                mock_create_thread,
                                                                mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.return_value = MagicMock()
        mock_run_thread.side_effect = Exception("Thread could not be run. check thread_id and assistant_id.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # 3.10 Failure in Token Encoding
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_thread')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_message')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.run_thread')
    @patch('services.openai.assistant_api_adapter.AssistantAPIAdapter.encode_token')
    def test_generate_explanation__encode_token_failure(self, mock_encode_token, mock_run_thread, mock_create_message,
                                                        mock_create_thread,
                                                        mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.return_value = MagicMock()
        mock_run_thread.return_value = MagicMock()

        mock_encode_token.side_effect = AssistantAPIException("Token encoding error")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_explanation("Woordenschat",
                                                          "wat doe je als je ‘een oogje in het zeil houdt’?",
                                                          "je bent op zoek naar een specifiek zeilschip",
                                                          "je let goed op om te zorgen dat alles goed gaat")

    # ALL TESTS FOR: RETRIEVE_RESPONSE

    # SUCCESS TESTS

    # 4. Successfully running the method

    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.retrieve_messages')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.delete_assistant')
    def test_retrieve_response_success(
            self, mock_delete_assistant, mock_retrieve_messages):
        # Arrange

        mock_data = SyncCursorPage[ThreadMessage](
            data=[
                ThreadMessage(
                    id='msg_dj7lOVE8hr2yfpbSNKEANORy',
                    assistant_id='asst_SgpG7DNnrZH5PXjYQFCIKfvz',
                    content=[
                        MessageContentText(
                            text=Text(
                                annotations=[],
                                value='```json\n{\n  "questions": [\n    {\n      "question_type": "multiple-choice",\n      "question_level": "B2",\n      "question_subject": "werkwoordvervoegingen",\n      "question_text": "Hoe vervoeg je het werkwoord \'lopen\' in de verleden tijd, eerste persoon enkelvoud?",\n      "answer_options": [\n        "Ik liep",\n        "Ik loopte",\n        "Ik lop",\n        "Ik liepen",\n        "Ik loopde"\n      ],\n      "correct_answer": "Ik liep"\n    },\n    {\n      "question_type": "multiple-choice",\n      "question_level": "B2",\n      "question_subject": "werkwoordvervoegingen",\n      "question_text": "Wat is de correcte vervoeging van het werkwoord \'vliegen\' in de onvoltooid tegenwoordige tijd, derde persoon meervoud?",\n      "answer_options": [\n        "Ze vliegt",\n        "Zij vliegen",\n        "Ze vlieg",\n        "Zij vliegs",\n        "Zij vliegden"\n      ],\n      "correct_answer": "Zij vliegen"\n    }\n  ]\n}\n```',
                            ),
                            type='text'
                        )
                    ],
                    created_at=1705271176,
                    file_ids=[],
                    metadata={},
                    object='thread.message',
                    role='assistant',
                    run_id='run_O0OV2jBtQEePTC7574JwIjnH',
                    thread_id='thread_YVkCtkv1IAO4jSvkA2GbvMzs'
                ),
                ThreadMessage(
                    id='msg_pAOrNGz9pvmSElU3oEFoPnNU',
                    assistant_id=None,
                    content=[
                        MessageContentText(
                            text=Text(
                                annotations=[],
                                value='onderwerp: EnumMeta.werkwoordvervoegingen, nederlands niveau: EnumMeta.B2, aantal vragen: 2',
                            ),
                            type='text'
                        )
                    ],
                    created_at=1705271175,
                    file_ids=[],
                    metadata={},
                    object='thread.message',
                    role='user',
                    run_id=None,
                    thread_id='thread_YVkCtkv1IAO4jSvkA2GbvMzs'
                )
            ],
            object='list',
            first_id='msg_dj7lOVE8hr2yfpbSNKEANORy',
            last_id='msg_pAOrNGz9pvmSElU3oEFoPnNU',
            has_more=False
        )

        # Mock validation function
        def mock_is_valid_generated_questions_json(json_data_dict):
            pass

        mock_retrieve_messages.return_value = mock_data
        mock_delete_assistant.return_value = MagicMock(return_value=True)

        # Act
        result = self.assistantAPIAdapter.retrieve_response(
            "eyJ0aHJlYWRfaWQiOiAidGhyZWFkX1lWa0N0a3YxSUFPNGpTdmtBMkdidk16cyIsICJhc3Npc3RhbnRfaWQiOiAiYXNzdF9TZ3BHN0RObnJaSDVQWGpZUUZDSUtmdnoiLCAiZW5kcG9pbnRfaWQiOiAicXVlc3Rpb25zIn0=0",
            "questions", mock_is_valid_generated_questions_json)

        # Assert
        mock_retrieve_messages.assert_called_with("thread_YVkCtkv1IAO4jSvkA2GbvMzs")
        mock_delete_assistant.assert_called_once_with("asst_SgpG7DNnrZH5PXjYQFCIKfvz")

        # Creating expected output json
        json_data = '''
        {
          "questions": [
            {
              "question_type": "multiple-choice",
              "question_level": "B2",
              "question_subject": "werkwoordvervoegingen",
              "question_text": "Hoe vervoeg je het werkwoord 'lopen' in de verleden tijd, eerste persoon enkelvoud?",
              "answer_options": [
                "Ik liep",
                "Ik loopte",
                "Ik lop",
                "Ik liepen",
                "Ik loopde"
              ],
              "correct_answer": "Ik liep"
            },
            {
              "question_type": "multiple-choice",
              "question_level": "B2",
              "question_subject": "werkwoordvervoegingen",
              "question_text": "Wat is de correcte vervoeging van het werkwoord 'vliegen' in de onvoltooid tegenwoordige tijd, derde persoon meervoud?",
              "answer_options": [
                "Ze vliegt",
                "Zij vliegen",
                "Ze vlieg",
                "Zij vliegs",
                "Zij vliegden"
              ],
              "correct_answer": "Zij vliegen"
            }
          ]
        }
        '''

        expected_json_data_dict = json.loads(json_data)

        self.assertEqual(result, expected_json_data_dict)


if __name__ == '__main__':
    unittest.main()
