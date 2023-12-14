from graphene import ObjectType, String, Field
from graphql import GraphQLError
from .Types import ResponseExplanation, ResponseAnswer, ResponseMultipleChoiceQuestion
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter
from .Validators import validate_base64


class Query(ObjectType):
    retrieve_multiple_choice_questions = Field(ResponseMultipleChoiceQuestion, token=String(required=True))
    # retrieve_open_answer_questions = Field(ResponseOpenAnswerQuestion, token=String(required=True))
    retrieve_explanation = Field(ResponseExplanation, token=String(required=True))
    retrieve_answer = Field(ResponseAnswer, token=String(required=True))

    async def resolve_retrieve_multiple_choice_questions(self, info, token):

        try:
            # validate if token is in base64
            validate_base64(token)

            adapter = AssistantAPIAdapter()
            response = adapter.retrieve_questions_response(token)

            return response

        except Exception as e:
            raise GraphQLError(str(e))

    # async def resolve_retrieve_open_answer_questions(self, info, token):
    #
    #     try:
    #         adapter = AssistantAPIAdapter()
    #         response = adapter.retrieve_response(token)
    #
    #         return response
    #
    #     except Exception as e:
    #         raise GraphQLError(str(e))

    async def resolve_retrieve_explanation(self, info, token):

        try:
            # validate if token is in base64
            validate_base64(token)

            adapter = AssistantAPIAdapter()
            response = adapter.retrieve_explanation_response(token)

            return response

        except Exception as e:
            raise GraphQLError(str(e))

    async def resolve_retrieve_answer(self, info, token):

        try:
            # validate if token is in base64
            validate_base64(token)

            adapter = AssistantAPIAdapter()
            response = adapter.retrieve_question_response(token)

            return response

        except Exception as e:
            raise GraphQLError(str(e))
