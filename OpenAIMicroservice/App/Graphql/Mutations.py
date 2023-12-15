import graphene
from graphql import GraphQLError
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter
from .Types import SubjectEnum, LevelEnum
from App.Utils.Exceptions import ValidationException, AssistantAPIException
from .Validators import validate_minimum_int, validate_string, validate_answer_options, validate_file


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


class GenerateMultipleChoiceQuestions(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_level = LevelEnum(required=True)
        amount_questions = graphene.Int(required=True)

    token = graphene.String()

    def mutate(self, info, question_subject, question_level, amount_questions):

        try:
            # validate amount questions
            validate_minimum_int("amount_questions", amount_questions)

            adapter = AssistantAPIAdapter()
            token = adapter.generate_multiple_choice_questions(question_subject, question_level, amount_questions)

            return GenerateMultipleChoiceQuestions(token=token)

        except ValidationException as e:
            raise GraphQLError(str(e))
        except AssistantAPIException as e:
            raise GraphQLError(str(e))
        except Exception:
            # Generic error for unexpected exceptions
            raise GraphQLError("An unexpected error occurred. Please try again later.")


class ReadMultipleChoiceQuestionsFromFile(graphene.Mutation):
    class Arguments:
        file_data = graphene.String(required=True)
        filename = graphene.String(required=True)

    token = graphene.String()

    def mutate(self, info, file_data, filename):

        try:
            # validate file
            validate_file(file_data)

            adapter = AssistantAPIAdapter()
            token = adapter.read_multiple_choice_questions_from_file(file_data, filename)
            return ReadMultipleChoiceQuestionsFromFile(token=token)

        except ValidationException as e:
            raise GraphQLError(str(e))
        except AssistantAPIException as e:
            raise GraphQLError(str(e))
        except Exception:
            # Generic error for unexpected exceptions
            raise GraphQLError("An unexpected error occurred. Please try again later.")


class GenerateExplanation(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_text = graphene.String(required=True)
        given_answer = graphene.String(required=True)
        correct_answer = graphene.String(required=True)

    token = graphene.String()

    def mutate(self, info, question_subject, question_text, given_answer, correct_answer):
        try:
            # validate given strings
            validate_string("question_text", question_text)
            validate_string("given_answer", given_answer)
            validate_string("correct_answer", correct_answer)

            adapter = AssistantAPIAdapter()
            token = adapter.generate_explanation(question_subject, question_text, given_answer, correct_answer)

            return GenerateExplanation(token=token)

        except ValidationException as e:
            raise GraphQLError(str(e))
        except AssistantAPIException as e:
            raise GraphQLError(str(e))
        except Exception:
            # Generic error for unexpected exceptions
            raise GraphQLError("An unexpected error occurred. Please try again later.")


class GenerateMultipleChoiceAnswer(graphene.Mutation):
    class Arguments:
        question_level = LevelEnum(required=True)
        question_subject = SubjectEnum(required=True)
        question_text = graphene.String(required=True)
        answer_options = graphene.List(graphene.String, required=True)

    token = graphene.String()

    def mutate(self, info, question_level, question_subject, question_text, answer_options):
        try:
            # validate given string
            validate_string("question_text", question_text)

            # validate given answer options
            validate_answer_options(answer_options)

            adapter = AssistantAPIAdapter()
            token = adapter.generate_multiple_choice_answer(question_level, question_subject, question_text,
                                                            answer_options)

            return GenerateExplanation(token=token)

        except ValidationException as e:
            raise GraphQLError(str(e))
        except AssistantAPIException as e:
            raise GraphQLError(str(e))
        except Exception:
            # Generic error for unexpected exceptions
            raise GraphQLError("An unexpected error occurred. Please try again later.")


class Mutation(graphene.ObjectType):
    generate_explanation = GenerateExplanation.Field()
    # generate_open_answer_questions = GenerateOpenAnswerQuestions.Field()
    read_multiple_choice_questions_from_file = ReadMultipleChoiceQuestionsFromFile.Field()
    generate_multiple_choice_questions = GenerateMultipleChoiceQuestions.Field()
    generate_multiple_choice_answer = GenerateMultipleChoiceAnswer.Field()
