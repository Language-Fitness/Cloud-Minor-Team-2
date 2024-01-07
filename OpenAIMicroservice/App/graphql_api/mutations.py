import graphene
from services.openai.assistant_api_adapter import AssistantAPIAdapter
from utils.exceptions.security_exception import SecurityException
from .types import SubjectEnum, LevelEnum, TokenResponse
from utils.exceptions.validation_exception import ValidationException
from utils.exceptions.assistant_api_exception import AssistantAPIException
from .validators import validate_minimum_int, validate_string, validate_file
from .security import Security

class GenerateMultipleChoiceQuestions(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_level = LevelEnum(required=True)
        amount_questions = graphene.Int(required=True)

    response = graphene.Field(TokenResponse)

    def mutate(self, info, question_subject, question_level, amount_questions):
        try:
            security = Security()
            bearer_token = security.extract_token_from_header(info)
            security.validate_token(bearer_token)
            security.has_required_role(bearer_token, "openai_generate_questions")

            # validate amount questions
            validate_minimum_int("amount_questions", amount_questions)
            adapter = AssistantAPIAdapter(bearer_token)

            token = adapter.generate_multiple_choice_questions(question_subject, question_level, amount_questions)

            return GenerateMultipleChoiceQuestions(
                TokenResponse(status="success", message="Generating questions started!", token=token))
        except SecurityException as e:
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message=str(e)))
        except ValidationException as e:
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message=str(e)))
        except AssistantAPIException as e:
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message=str(e)))
        except Exception as e:
            # Generic error for unexpected exceptions
            print(str(e))
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error",
                                                                 message="An unexpected error occurred while generating questions. Please try again later."))

class ReadMultipleChoiceQuestionsFromFile(graphene.Mutation):
    class Arguments:
        file_data = graphene.String(required=True)
        filename = graphene.String(required=True)

    response = graphene.Field(TokenResponse)

    def mutate(self, info, file_data, filename):

        try:
            security = Security()
            bearer_token = security.extract_token_from_header(info)
            security.validate_token(bearer_token)
            security.has_required_role(bearer_token, "openai_generate_questions_from_file")

            # validate file
            validate_file(file_data)

            adapter = AssistantAPIAdapter(bearer_token)
            token = adapter.read_multiple_choice_questions_from_file(file_data, filename)
            return ReadMultipleChoiceQuestionsFromFile(
                TokenResponse(status="success", message="Reading questions from the file started!", token=token))

        except SecurityException as e:
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message=str(e)))
        except ValidationException as e:
            return ReadMultipleChoiceQuestionsFromFile(TokenResponse(status="error", message=str(e)))
        except AssistantAPIException as e:
            return ReadMultipleChoiceQuestionsFromFile(TokenResponse(status="error", message=str(e)))
        except Exception:
            # Generic error for unexpected exceptions
            return ReadMultipleChoiceQuestionsFromFile(TokenResponse(status="error",
                                                                     message="An unexpected error occurred while reading questions from the file. Please try again later."))


class GenerateExplanation(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_text = graphene.String(required=True)
        given_answer = graphene.String(required=True)
        correct_answer = graphene.String(required=True)

    response = graphene.Field(TokenResponse)

    def mutate(self, info, question_subject, question_text, given_answer, correct_answer):
        try:
            security = Security()
            bearer_token = security.extract_token_from_header(info)
            security.validate_token(bearer_token)
            security.has_required_role(bearer_token, "openai_generate_explanation")

            # validate given strings
            validate_string("question_text", question_text)
            validate_string("given_answer", given_answer)
            validate_string("correct_answer", correct_answer)

            adapter = AssistantAPIAdapter(bearer_token)
            token = adapter.generate_explanation(question_subject, question_text, given_answer, correct_answer)

            return GenerateExplanation(
                TokenResponse(status="success", message="Generating explanation started!.", token=token))

        except SecurityException as e:
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message=str(e)))
        except ValidationException as e:
            return GenerateExplanation(TokenResponse(status="error", message=str(e)))
        except AssistantAPIException as e:
            return GenerateExplanation(TokenResponse(status="error", message=str(e)))
        except Exception:
            # Generic error for unexpected exceptions
            return GenerateExplanation(TokenResponse(status="error",
                                                     message="An unexpected error occurred while generating explanation. Please try again later."))


class Mutation(graphene.ObjectType):
    generate_explanation = GenerateExplanation.Field()
    read_multiple_choice_questions_from_file = ReadMultipleChoiceQuestionsFromFile.Field()
    generate_multiple_choice_questions = GenerateMultipleChoiceQuestions.Field()
