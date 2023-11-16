# GraphQL Queries and Mutations

## Queries

### Get Exercise by Module ID

```graphql
query {
  GetExerciseByModuleId(moduleId: "your_module_id_here") {
    id
    class_Id
    name
    question
    answers
    pos_correct_answer
    question_type_id
    difficulty
    created_at
    updated_at
    soft_deleted
  }
}
```

## Mutations
### Add Exercise
```graphql
mutation {
  AddExercise(exercise: {
    class_Id: "your_class_id_here"
    name: "Exercise Name"
    question: "Exercise Question"
    answers: "Comma,Separated,Answers"
    pos_correct_answer: 1
    question_type_id: "your_question_type_id_here"
    difficulty: 0.5
  }) {
    id
    class_Id
    name
    question
    answers
    pos_correct_answer
    question_type_id
    difficulty
    created_at
    updated_at
    soft_deleted
  }
}
```
### Update Exercise
```graphql
mutation {
  UpdateExercise(exercise: {
    class_Id: "your_class_id_here"
    name: "Updated Exercise Name"
    question: "Updated Exercise Question"
    answers: "Updated,Comma,Separated,Answers"
    pos_correct_answer: 2
    question_type_id: "your_question_type_id_here"
    difficulty: 0.7
  }) {
    id
    class_Id
    name
    question
    answers
    pos_correct_answer
    question_type_id
    difficulty
    created_at
    updated_at
    soft_deleted
  }
}
```
### Delete Exercise
```graphql
mutation {
  DeleteExercise(id: "your_exercise_id_here") {
    id
    class_Id
    name
    question
    answers
    pos_correct_answer
    question_type_id
    difficulty
    created_at
    updated_at
    soft_deleted
  }
}
```