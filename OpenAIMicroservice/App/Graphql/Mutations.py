import graphene
from Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter
from .Types import SubjectEnum, LevelEnum, TokenResponse
from Utils.Exceptions.ValidationException import ValidationException
from Utils.Exceptions.AssistantAPIException import AssistantAPIException
from .Validators import validate_minimum_int, validate_string, validate_file



class GenerateMultipleChoiceQuestions(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_level = LevelEnum(required=True)
        amount_questions = graphene.Int(required=True)

    response = graphene.Field(TokenResponse)

    def mutate(self, info, question_subject, question_level, amount_questions):

        try:
            # validate amount questions
            validate_minimum_int("amount_questions", amount_questions)

            adapter = AssistantAPIAdapter()
            token = adapter.generate_multiple_choice_questions(question_subject, question_level, amount_questions)

            return GenerateMultipleChoiceQuestions(TokenResponse(status="success", message="Generating questions started!", token=token))

        except ValidationException as e:
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message=str(e)))
        except AssistantAPIException as e:
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message=str(e)))
        except Exception:
            # Generic error for unexpected exceptions
            return GenerateMultipleChoiceQuestions(TokenResponse(status="error", message="An unexpected error occurred while generating questions. Please try again later."))


class ReadMultipleChoiceQuestionsFromFile(graphene.Mutation):
    class Arguments:
        file_data = graphene.String(required=True)
        filename = graphene.String(required=True)

    response = graphene.Field(TokenResponse)

    def mutate(self, info, file_data, filename):

        try:
            # validate file
            validate_file(file_data)

            adapter = AssistantAPIAdapter()
            token = adapter.read_multiple_choice_questions_from_file(file_data, filename)
            return ReadMultipleChoiceQuestionsFromFile(TokenResponse(status="success", message="Reading questions from the file started!", token=token))

        except ValidationException as e:
            return ReadMultipleChoiceQuestionsFromFile(TokenResponse(status="error", message=str(e)))
        except AssistantAPIException as e:
            return ReadMultipleChoiceQuestionsFromFile(TokenResponse(status="error", message=str(e)))
        except Exception:
            # Generic error for unexpected exceptions
            return ReadMultipleChoiceQuestionsFromFile(TokenResponse(status="error", message="An unexpected error occurred while reading questions from the file. Please try again later."))


class GenerateExplanation(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_text = graphene.String(required=True)
        given_answer = graphene.String(required=True)
        correct_answer = graphene.String(required=True)

    response = graphene.Field(TokenResponse)

    def mutate(self, info, question_subject, question_text, given_answer, correct_answer):
        try:
            # validate given strings
            validate_string("question_text", question_text)
            validate_string("given_answer", given_answer)
            validate_string("correct_answer", correct_answer)

            adapter = AssistantAPIAdapter()
            token = adapter.generate_explanation(question_subject, question_text, given_answer, correct_answer)

            return GenerateExplanation(TokenResponse(status="success", message="Generating explanation started!.", token=token))

        except ValidationException as e:
            return GenerateExplanation(TokenResponse(status="error", message=str(e)))
        except AssistantAPIException as e:
            return GenerateExplanation(TokenResponse(status="error", message=str(e)))
        except Exception:
            # Generic error for unexpected exceptions
            return GenerateExplanation(TokenResponse(status="error", message="An unexpected error occurred while generating explanation. Please try again later."))


# class GenerateOpenAnswerQuestions(graphene.Mutation):
#     class Arguments:
#         subject = SubjectEnum(required=True)
#         level = LevelEnum(required=True)
#         amount_questions = graphene.Int(required=True)
#
#     token = graphene.String()
#
#     def mutate(self, info, subject, level, amount_questions):
#         adapter = AssistantAPIAdapter()
#         token = adapter.generate_open_answer_questions(subject, level, amount_questions)
#
#         return GenerateOpenAnswerQuestions(token=token)


class Mutation(graphene.ObjectType):
    generate_explanation = GenerateExplanation.Field()
    # generate_open_answer_questions = GenerateOpenAnswerQuestions.Field()
    read_multiple_choice_questions_from_file = ReadMultipleChoiceQuestionsFromFile.Field()
    generate_multiple_choice_questions = GenerateMultipleChoiceQuestions.Field()
