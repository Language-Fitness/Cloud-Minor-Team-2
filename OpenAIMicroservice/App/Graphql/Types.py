import graphene

class MultipleChoiceQuestion(graphene.ObjectType):
    type = graphene.String()
    question_info = graphene.String()
    question_text = graphene.String()
    options = graphene.List(graphene.String)
    answer = graphene.String()

class ResponseMultipleChoiceQuestion(graphene.ObjectType):
    status = graphene.String()
    data = graphene.List(MultipleChoiceQuestion)

class OpenAnswerQuestion(graphene.ObjectType):
    type = graphene.String()
    question_info = graphene.String()
    question_text = graphene.String()
    answer = graphene.String()

class ResponseOpenAnswerQuestion(graphene.ObjectType):
    status = graphene.String()
    questions = graphene.List(OpenAnswerQuestion)

class Explanation(graphene.ObjectType):
    info = graphene.String()
    tips = graphene.String()

class ResponseExplanation(graphene.ObjectType):
    status = graphene.String()
    explanation = graphene.Field(Explanation)

class ResponseAnswer(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    answer = graphene.String(description="The answer provided in the response")
