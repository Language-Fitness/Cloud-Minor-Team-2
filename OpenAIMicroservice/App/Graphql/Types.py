import graphene

# Define custom GraphQL types here, for example:

class Question(graphene.ObjectType):
    text = graphene.String(description="The text of the question.")

class Answer(graphene.ObjectType):
    text = graphene.String(description="The text of the answer.")
    confidence = graphene.Float(description="The confidence score of the answer.")
