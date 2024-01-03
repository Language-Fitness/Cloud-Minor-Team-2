import graphene


class MultipleChoiceQuestion(graphene.ObjectType):
    question_type = graphene.String(description="Type question")
    question_level = graphene.String(description="Question subject")
    question_subject = graphene.String(description="Question level")
    question_text = graphene.String(description="Question text")
    answer_options = graphene.List(graphene.String, description="Answer options")
    correct_answer = graphene.String(description="Correct answer")


class ResponseMultipleChoiceQuestion(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    message = graphene.String(description="Message")
    questions = graphene.List(MultipleChoiceQuestion, description="Multiple choice array")


class Explanation(graphene.ObjectType):
    info = graphene.String(description="Information for explanation")
    tips = graphene.String(description="Tips for next time")


class ResponseExplanation(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    message = graphene.String(description="Message")
    explanation = graphene.Field(Explanation, description="Explanation field")


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
    uitdrukkingen = "Uitdrukkingen"
    interpunctie = "Interpunctie"
    werkwoordvervoegingen = "Werkwoordvervoegingen"


class TokenResponse(graphene.ObjectType):
    status = graphene.String(description="Status of the response")
    message = graphene.String(description="Message")
    token = graphene.String(description="Token")


class TokenData(graphene.ObjectType):
    assistant_id = graphene.String()
    thread_id = graphene.String()
