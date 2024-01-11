import base64
import json
import unittest
from unittest.mock import patch, MagicMock
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

    def setUp(self):
        self.assistantAPIAdapter = AssistantAPIAdapter()

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
    def test_load_assistant_failure_exception(self, mock_load_assistant):
        mock_load_assistant.side_effect = Exception("Load assistant failure")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.2 Failure in Creating Assistant (Exception)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    def test_create_assistant_failure_exception(self, mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.side_effect = Exception(
            "Message could not be created. check fields json fields in assistant json.")
        with self.assertRaises(Exception):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # 1.3 Failure in Creating Assistant (AssistantAPIException)
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.load_assistant')
    @patch('services.openai.openai_assistant_manager.OpenAIAssistantManager.create_assistant')
    def test_create_assistant_failure_assistant_api_exception(self, mock_create_assistant, mock_load_assistant):
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
    def test_create_thread_failure_assistant_api_exception(self, mock_create_thread, mock_create_assistant,
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
    def test_create_message_failure_assistant_api_exception(self, mock_create_message, mock_create_thread,
                                                            mock_create_assistant, mock_load_assistant):
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
    def test_create_message_failure_exception(self, mock_create_message, mock_create_thread, mock_create_assistant,
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
    def test_run_thread_failure_assistant_api_exception(self, mock_run_thread, mock_create_message, mock_create_thread,
                                                        mock_create_assistant, mock_load_assistant):
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
    def test_run_thread_failure_exception(self, mock_run_thread, mock_create_message, mock_create_thread,
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
    def test_encode_token_failure(self, mock_encode_token, mock_run_thread, mock_create_message, mock_create_thread,
                                  mock_create_assistant, mock_load_assistant):
        mock_load_assistant.return_value = MagicMock()
        mock_create_assistant.return_value = MagicMock()
        mock_create_thread.return_value = MagicMock()
        mock_create_message.return_value = MagicMock()
        mock_run_thread.return_value = MagicMock()

        mock_encode_token.side_effect = AssistantAPIException("Token encoding error")
        with self.assertRaises(AssistantAPIException):
            self.assistantAPIAdapter.generate_multiple_choice_questions("Subject", "Level", 5)

    # ALL TESTS FOR: GENERATE_EXPLANATION

    # SUCCESS TESTS

    # 2. Successfully running the method

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
        result = self.assistantAPIAdapter.generate_explanation("Woordenschat", "wat doe je als je ‘een oogje in het zeil houdt’?", "je bent op zoek naar een specifiek zeilschip", "je let goed op om te zorgen dat alles goed gaat")

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

    # 1.1 Failure in Loading Assistant


if __name__ == '__main__':
    unittest.main()
