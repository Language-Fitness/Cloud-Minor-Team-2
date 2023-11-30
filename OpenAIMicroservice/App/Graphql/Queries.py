from graphene import ObjectType, String, Field
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter


class Query(ObjectType):
    retrieve_multiple_choice_questions = Field(String, thread_id=String(required=True))
    retrieve_open_answer_questions = Field(String, thread_id=String(required=True))
    retrieve_explanation = Field(String, thread_id=String(required=True))
    retrieve_answer = Field(String, thread_id=String(required=True))

    async def resolve_retrieve_multiple_choice_questions(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_multiple_choice_questions(token)
        return response

    async def resolve_retrieve_open_answer_questions(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_open_answer_questions(token)
        return response

    async def resolve_retrieve_explanation(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_explanation_questions(token)
        return response

    async def resolve_retrieve_answer(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.retrieve_answer(token)
        return response
