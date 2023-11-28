import graphene

class Response(graphene.ObjectType):
    type = graphene.String()
    data = graphene.JSONString()