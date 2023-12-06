import graphene
from App.Services.OpenAI.AssistantAPIAdapter import AssistantAPIAdapter
from .Types import SubjectEnum, LevelEnum


# class GenerateOpenAnswerQuestions(graphene.Mutation):
#     class Arguments:
#         subject = SubjectEnum(required=True)
#         level = LevelEnum(required=True)
#         amount_questions = graphene.Int(required=True)
#
#     token = graphene.String()
#
#     def mutate(self, info, subject, level, amount_questions):
#         adapter = AssistantAPIAdapter()
#         token = adapter.generate_open_answer_questions(subject, level, amount_questions)
#
#         return GenerateOpenAnswerQuestions(token=token)


class GenerateMultipleChoiceQuestions(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_level = LevelEnum(required=True)
        amount_questions = graphene.Int(required=True)

    token = graphene.String()

    def mutate(self, info, question_subject, question_level, amount_questions):
        adapter = AssistantAPIAdapter()
        token = adapter.generate_multiple_choice_questions(question_subject, question_level, amount_questions)

        return GenerateMultipleChoiceQuestions(token=token)


class GenerateExplanation(graphene.Mutation):
    class Arguments:
        question_subject = SubjectEnum(required=True)
        question_text = graphene.String(required=True)
        given_answer = graphene.String(required=True)
        correct_answer = graphene.String(required=True)

    token = graphene.String()

    def mutate(self, info, question_subject, question_text, given_answer, correct_answer):
        adapter = AssistantAPIAdapter()
        token = adapter.generate_explanation(question_subject, question_text, given_answer, correct_answer)

        return GenerateExplanation(token=token)


class GenerateMultipleChoiceAnswer(graphene.Mutation):
    class Arguments:
        question_level = LevelEnum(required=True)
        question_subject = SubjectEnum(required=True)
        question_text = graphene.String(required=True)
        answer_options = graphene.List(graphene.String, required=True)

    token = graphene.String()

    def mutate(self, info, question_level, question_subject, question_text, answer_options):
        adapter = AssistantAPIAdapter()
        token = adapter.generate_multiple_choice_answer(question_level, question_subject, question_text, answer_options)

        return GenerateExplanation(token=token)


class Mutation(graphene.ObjectType):
    generate_explanation = GenerateExplanation.Field()
    # generate_open_answer_questions = GenerateOpenAnswerQuestions.Field()
    generate_multiple_choice_questions = GenerateMultipleChoiceQuestions.Field()
    generate_answer = GenerateMultipleChoiceAnswer.Field()
