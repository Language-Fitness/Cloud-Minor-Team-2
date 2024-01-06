from graphene import ObjectType, String, Field
from .types import ResponseExplanation, ResponseMultipleChoiceQuestion
from services.openai.assistant_api_adapter import AssistantAPIAdapter
from utils.exceptions.assistant_api_exception import AssistantAPIException
from utils.exceptions.validation_exception import ValidationException
from utils.exceptions.security_exception import SecurityException
from .security import Security
from .validators import validate_base64


class Query(ObjectType):
    retrieve_multiple_choice_questions = Field(ResponseMultipleChoiceQuestion, token=String(required=True))
    retrieve_explanation = Field(ResponseExplanation, token=String(required=True))
    retrieve_multiple_choice_questions_from_file = Field(ResponseMultipleChoiceQuestion, token=String(required=True))

    async def resolve_retrieve_multiple_choice_questions(self, info, token):

        try:
            security = Security()
            bearer_token = security.extract_token_from_header(info)
            # security.validate_token(bearer_token)
            # security.has_required_role(bearer_token, "openai_generate_questions")


            # validate if token is in base64
            validate_base64(token)

            adapter = AssistantAPIAdapter(bearer_token)
            response = adapter.retrieve_generated_multiple_choice_questions_response(token)

            return ResponseMultipleChoiceQuestion(status="success", message="Question(s) retrieved successfully",
                                                  questions=response['questions'])
        except SecurityException as e:
            return ResponseMultipleChoiceQuestion(status="error", message=str(e))
        except ValidationException as e:
            return ResponseMultipleChoiceQuestion(status="error", message=str(e))
        except AssistantAPIException as e:
            return ResponseMultipleChoiceQuestion(status="error", message=str(e))
        except Exception:
            return ResponseMultipleChoiceQuestion(status="error",
                                                  message="An unexpected error occurred. Please try again later.")

    async def resolve_retrieve_multiple_choice_questions_from_file(self, info, token):

        try:
            security = Security()
            bearer_token = security.extract_token_from_header(info)
            security.validate_token(bearer_token)
            security.has_required_role(bearer_token, "openai_generate_questions_from_file")

            # validate if token is in base64
            validate_base64(token)

            adapter = AssistantAPIAdapter(bearer_token)
            response = adapter.retrieve_multiple_choice_questions_from_file_response(token)

            return ResponseMultipleChoiceQuestion(status="success", message="Question(s) retrieved successfully",
                                                  questions=response['questions'])

        except SecurityException as e:
            return ResponseMultipleChoiceQuestion(status="error", message=str(e))
        except ValidationException as e:
            return ResponseMultipleChoiceQuestion(status="error", message=str(e))

        except AssistantAPIException as e:
            return ResponseMultipleChoiceQuestion(status="error", message=str(e))

        except Exception as e:
            return ResponseMultipleChoiceQuestion(status="error",
                                                  message="An unexpected error occurred. Please try again later.")

    async def resolve_retrieve_explanation(self, info, token):

        try:
            security = Security()
            bearer_token = security.extract_token_from_header(info)
            security.validate_token(bearer_token)
            security.has_required_role(bearer_token, "openai_generate_explanation")

            # validate if token is in base64
            validate_base64(token)

            adapter = AssistantAPIAdapter(bearer_token)
            response = adapter.retrieve_explanation_response(token)

            return ResponseExplanation(status="success", message="Explanation retrieved successfully",
                                       explanation=response["explanation"])

        except SecurityException as e:
            return ResponseExplanation(status="error", message=str(e))
        except ValidationException as e:
            return ResponseExplanation(status="error", message=str(e))

        except AssistantAPIException as e:
            return ResponseExplanation(status="error", message=str(e))

        except Exception as e:
            return ResponseExplanation(status="error", message="An unexpected error occurred. Please try again later.")
