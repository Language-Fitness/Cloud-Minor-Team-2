import graphene
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter


class GenerateOpenAnswerQuestions(graphene.Mutation):
    class Arguments:
        subject = graphene.String(required=True)
        amount_questions = graphene.Int(required=True)

    token = graphene.String()

    def mutate(self, info, subject, amount_questions):
        adapter = AssistantAPIAdapter()
        token = adapter.generate_open_answer_questions(subject, amount_questions)

        return GenerateOpenAnswerQuestions(token=token)


class GenerateMultipleChoiceQuestions(graphene.Mutation):
    class Arguments:
        subject = graphene.String(required=True)
        amount_questions = graphene.Int(required=True)

    token = graphene.String()

    def mutate(self, info, subject, amount_questions):
        adapter = AssistantAPIAdapter()
        token = adapter.generate_multiple_choice_questions(subject, amount_questions)

        return GenerateMultipleChoiceQuestions(token=token)


class GenerateExplanation(graphene.Mutation):
    class Arguments:
        question = graphene.String(required=True)
        given_answer = graphene.String(required=True)

    token = graphene.String()

    def mutate(self, info, question, given_answer):
        adapter = AssistantAPIAdapter()
        token = adapter.generate_explanation(question, given_answer)

        return GenerateExplanation(token=token)


class GenerateAnswer(graphene.Mutation):
    class Arguments:
        question = graphene.String(required=True)
        question_info = graphene.String(required=True)

    token = graphene.String()

    def mutate(self, info, question, question_info):
        adapter = AssistantAPIAdapter()
        token = adapter.generate_answer(question, question_info)

        return GenerateExplanation(token=token)


class Mutation(graphene.ObjectType):
    generate_explanation = GenerateExplanation.Field()
    generate_open_answer_questions = GenerateOpenAnswerQuestions.Field()
    generate_multiple_choice_questions = GenerateMultipleChoiceQuestions.Field()
    generate_answer = GenerateAnswer.Field()
