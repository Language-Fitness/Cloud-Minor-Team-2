import graphene
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter

class GenerateOpenAnswerQuestions(graphene.Mutation):
    class Arguments:
        subject = graphene.String(required=True)
        amount_questions = graphene.Int(required=True)


    thread_id = graphene.String()


    def mutate(self, info, subject, amount_questions):
        adapter = AssistantAPIAdapter()
        test = adapter.Generate_Multiple_Choice_Questions(subject, amount_questions)


        return GenerateOpenAnswerQuestions(thread_id=test)

class GenerateMultipleChoiceQuestions(graphene.Mutation):

    class Arguments:
        subject = graphene.String(required=True)
        amount_questions = graphene.Int(required=True)


    thread_id = graphene.String()

    def mutate(self, info, subject, amount_questions):

        adapter = AssistantAPIAdapter()
        test = adapter.Generate_Multiple_Choice_Questions(subject, amount_questions)

        return GenerateMultipleChoiceQuestions(thread_id=test)

class GenerateExplanation(graphene.Mutation):
    class Arguments:
        question = graphene.String(required=True)
        given_answer = graphene.String(required=True)


    thread_id = graphene.String()


    def mutate(self, info, question, given_answer):
        adapter = AssistantAPIAdapter()
        test = adapter.Generate_Explanation(question, given_answer)

        return GenerateExplanation(thread_id=test)

class GenerateAnswer(graphene.Mutation):
    class Arguments:
        question = graphene.String(required=True)
        question_info = graphene.String(required=True)


    thread_id = graphene.String()

    def mutate(self, info, question, question_info):
        adapter = AssistantAPIAdapter()
        test = adapter.Generate_Answer(question, question_info)

        return GenerateExplanation(thread_id=test)


class Mutation(graphene.ObjectType):
    generate_explanation = GenerateExplanation.Field()
    generate_open_answer_questions = GenerateOpenAnswerQuestions.Field()
    generate_multiple_choice_questions = GenerateMultipleChoiceQuestions.Field()
    generate_answer = GenerateAnswer.Field()