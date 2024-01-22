export const headers = [
    {title: 'Name', key: 'name', align: 'start', sortable: false},
    {title: 'Question', key: 'question', align: 'start', sortable: false},
    {title: 'Answers', key: 'answers', align: 'start', sortable: false},
    {title: 'Difficulty', key: 'difficulty', align: 'start', sortable: false},
    {title: 'Actions', key: 'actions', align: 'start', sortable: false},
]

export const difficulties = ['A1', 'A2', 'B1', 'B2', 'C1', 'C2']

export const listExercisesQuery = `
    query ListExercise($filter: ExerciseFilter!, $paginator: Paginator!) {
      ListExercise(filter: $filter, paginator: $paginator) {
        answers {
          correct
          value
        }
        class_Id
        difficulty
        id
        made_by
        module_id
        name
        question
      }
    }
`;

export const deleteExerciseQuery = `
    mutation DeleteObject($filter: SagaFilter) {
      deleteObject(filter: $filter) {
        id
        object_id
        object_type
        text
        status
      }
    }
`;

export const createExerciseQuery = `
    mutation Mutation($exercise: ExerciseInput!) {
      CreateExercise(exercise: $exercise) {
        id
        module_id
        class_Id
        name
        difficulty
        question
        answers {
          value
          correct
        }
        made_by
        created_at
        updated_at
        soft_deleted
      }
    }
`;

export const generateMcQuery = `
  mutation generateMC($amountQuestions: Int!, $questionLevel: LevelEnum!, $questionSubject: SubjectEnum!) {
    generateMultipleChoiceQuestions (
      amountQuestions: $amountQuestions,
      questionLevel: $questionLevel,
      questionSubject: $questionSubject
    ) {
      response {
        status
        message
        token
      }
    }
  }
`;

export const retrieveMcQuery = `
    query RetrieveMultipleChoiceQuestions($token: String!) {
      retrieveMultipleChoiceQuestions(token: $token) {
        status
        message
        questions {
          questionType
          questionLevel
          questionSubject
          questionText
          answerOptions
          correctAnswer
        }
      }
    }
`