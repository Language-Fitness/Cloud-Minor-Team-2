import graphene

class MultipleChoiceQuestion(graphene.ObjectType):
    type = graphene.String(description="Type question")
    question_subject = graphene.String(description="Question subject")
    question_text = graphene.String(description="Question text")
    options = graphene.List(graphene.String, description="Answer options")
    answer = graphene.String(description="Correct answer")

class ResponseMultipleChoiceQuestion(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    questions = graphene.List(MultipleChoiceQuestion, description="Multiple choice array")

class OpenAnswerQuestion(graphene.ObjectType):
    type = graphene.String(description="Type question")
    question_subject = graphene.String(description="Question subject")
    question_text = graphene.String(description="Question text")
    answer = graphene.String(description="Correct answer")

class ResponseOpenAnswerQuestion(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    questions = graphene.List(OpenAnswerQuestion, description="Open answer array")

class Explanation(graphene.ObjectType):
    info = graphene.String(description="Information for explanation")
    tips = graphene.String(description="Tips for next time")

class ResponseExplanation(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    explanation = graphene.Field(Explanation, description="Explanation body")

class ResponseAnswer(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    answer = graphene.String(description="The answer provided in the response")

class LevelEnum(graphene.Enum):
    A1 = "A1"
    A2 = "A2"
    B1 = "B1"
    B2 = "B2"
    C1 = "C1"
    C2 = "C2"

class SubjectEnum(graphene.Enum):
    grammatica = "Grammatica"
    spelling = "Spelling"
    woordenschat = "Woordenschat"
    zinsontleding = "Zinsontleding"