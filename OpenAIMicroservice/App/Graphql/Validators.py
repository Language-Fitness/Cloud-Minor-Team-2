from graphql import GraphQLError

def validate_minimum(value):
    if value < 1:
        raise GraphQLError("Amount questions must be at least 1")