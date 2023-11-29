from graphene import ObjectType, String, Field
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter

class Query(ObjectType):
    retrieve_response = Field(String, thread_id=String(required=True))

    async def resolve_retrieve_response(self, info, thread_id):

        adapter = AssistantAPIAdapter()
        response = adapter.Retrieve_Response(thread_id)
        return response

    async def resolve_retrieve_multiple_choice_questions(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.Retrieve_Response(token)
        return response
    async def resolve_retrieve_open_answer_questions(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.Retrieve_Response(token)
        return response

    async def resolve_retrieve_explanation(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.Retrieve_Response(token)
        return response
    async def resolve_retrieve_answer(self, info, token):
        adapter = AssistantAPIAdapter()
        response = adapter.Retrieve_Response(token)
        return response