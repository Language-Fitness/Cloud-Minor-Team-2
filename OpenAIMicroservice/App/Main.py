from fastapi import FastAPI
from starlette_graphene3 import GraphQLApp, make_graphiql_handler
from graphql_api.schema import schema as graphene_schema
import uvicorn

app = FastAPI()

# Define the /graphql_api endpoint with GraphiQL enabled for development
app.add_websocket_route(
    "/query",
    GraphQLApp(schema=graphene_schema, on_get=make_graphiql_handler())
)

app.add_route(
    "/query",
    GraphQLApp(schema=graphene_schema, on_get=make_graphiql_handler())
)

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=4000)
