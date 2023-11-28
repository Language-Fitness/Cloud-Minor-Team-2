from fastapi import FastAPI
from starlette_graphene3 import GraphQLApp, make_graphiql_handler
from Graphql.Schema import schema as graphene_schema

app = FastAPI()

# Define the /graphql endpoint with GraphiQL enabled for development
app.add_websocket_route(
    "/",
    GraphQLApp(schema=graphene_schema, on_get=make_graphiql_handler())
)

app.add_route(
    "/",
    GraphQLApp(schema=graphene_schema, on_get=make_graphiql_handler())
)

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=4000)
