## Queries

### Get Result by Exercise ID
```graphql
query {
  GetResultByExercise(exercise_id: "your-exercise-id") {
    id
    exercise_id
    user_id
    class_id
    module_id
    input
    result
    created_at
    updated_at
    soft_deleted
  }
}
```

### Get Results by Class ID
```graphql
query {
  GetResultsByClass(class_id: "your-class-id") {
    id
    exercise_id
    user_id
    class_id
    module_id
    input
    result
    created_at
    updated_at
    soft_deleted
  }
}
```

## Mutations:
### Create Result
```graphql
mutation {
  CreateResult(input: {
    exercise_id: "your-exercise-id"
    user_id: "your-user-id"
    class_id: "your-class-id"
    module_id: "your-module-id"
    input: "your-input"
    result: "your-result"
  }) {
    id
    exercise_id
    user_id
    class_id
    module_id
    input
    result
    created_at
    updated_at
    soft_deleted
  }
}
```

## Update Result
```graphql
mutation {
  UpdateResult(id: "result-id-to-update", input: {
    exercise_id: "updated-exercise-id"
    user_id: "updated-user-id"
    class_id: "updated-class-id"
    module_id: "updated-module-id"
    input: "updated-input"
    result: "updated-result"
  }) {
    id
    exercise_id
    user_id
    class_id
    module_id
    input
    result
    created_at
    updated_at
    soft_deleted
  }
}
```

## Delete Result
```graphql
mutation {
  DeleteResult(id: "result-id-to-delete") {
    id
    exercise_id
    user_id
    class_id
    module_id
    input
    result
    created_at
    updated_at
    soft_deleted
  }
}
```