from graphene import ObjectType, String, Field
from graphql import GraphQLError
from .Types import ResponseExplanation, ResponseAnswer, ResponseOpenAnswerQuestion,ResponseMultipleChoiceQuestion
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter

class Query(ObjectType):
    retrieve_multiple_choice_questions = Field(ResponseMultipleChoiceQuestion, token=String(required=True))
    retrieve_open_answer_questions = Field(ResponseOpenAnswerQuestion, token=String(required=True))
    retrieve_explanation = Field(ResponseExplanation, token=String(required=True))
    retrieve_answer = Field(ResponseAnswer, token=String(required=True))

    async def resolve_retrieve_multiple_choice_questions(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_multiple_choice_questions(token)

        #if 'error' in response:
         #   raise GraphQLError(response['error'])

        return response

    async def resolve_retrieve_open_answer_questions(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_open_answer_questions(token)

        if 'error' in response:
            raise GraphQLError(response['error'])

        return response

    async def resolve_retrieve_explanation(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_explanation_questions(token)

        if 'error' in response:
            raise GraphQLError(response['error'])

        return response

    async def resolve_retrieve_answer(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_answer(token)

        if 'error' in response:
            raise GraphQLError(response['error'])

        return response