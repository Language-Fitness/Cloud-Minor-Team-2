from graphene import ObjectType, String, Field
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter

class Query(ObjectType):
    retrieve_response = Field(String, thread_id=String(required=True))

    async def resolve_retrieve_response(self, info, thread_id):
        # Your logic to retrieve the response goes here.
        # This is just a placeholder. Replace it with a call to your actual data retrieval function.
        adapter = AssistantAPIAdapter()
        explanation = adapter.Retrieve_Response(thread_id)
        return explanation