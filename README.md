## Basic Flow: usage of language fitness
Underlying queries will make a module with a class, exercises and results to demonstrate the basic 
usage of this application. The last query is a combined query to retrieve all data from a single module
and its child component in one single query to demonstrate the gateway.

**Prerequisites:**
_Usage of the gateway requires a bearer access token with needed permissions. This can only be obtained through the login
proces. (Token should be set in the Authorization header)_

### Create Module
```graphql
mutation CreateModule {
    createModule(input: {
        name: "Demo Module",
        school_id: "c9a49544-2661-4d0a-981d-07259ec75c2d"
        description: "This is a demo module.",
        difficulty: B2,
        category: Woordenschat,
        private: true,
        key: "demo-key"
    }) {
        id
        school_id
        name
        description
        difficulty
        category
        made_by
        private
        key
        created_at
        updated_at
        soft_deleted
    }
}
```

### Create Class
```graphql
mutation CreateClass {
    createClass(input: {
        module_Id: "id-of-created-module"
        name: "Demo class",
        description: "This is a demo class."
        difficulty: B2,
    }) {
        id
        name
        description
        difficulty
        made_by
        created_at
        updated_at
        soft_deleted
    }
}
```

### Create Exercises 1
```graphql
mutation {
    CreateExercise(exercise: {
        class_Id: "id-of-created-class",
        module_id: "id-of-created-module",
        name: "D of T - Oefening 1",
        question: "Hij (werken) ________ elke dag aan zijn project.",
        answers: [
            { value: "werkt", correct: true },
            { value: "werkd", correct: false },
            { value: "werken", correct: false }
        ],
        difficulty: B2,
    }) {
        id
        class_Id
        module_id
        name
        question
        answers {
            value
            correct
        }
        difficulty
        made_by
    }
}
```

### Create Exercise 2
```graphql
mutation {
    CreateExercise(exercise: {
        class_Id: "id-of-created-class",
        module_id: "id-of-created-module",
        name: "D of T - Oefening 2",
        question: "Wij (organiseren) ________ een feest voor zijn verjaardag.",
        answers: [
            { value: "organiseren", correct: true },
            { value: "organiseert", correct: false },
            { value: "organiseerd", correct: false }
        ],
        difficulty: B2,
    }) {
        id
        class_Id
        module_id
        name
        question
        answers {
            value
            correct
        }
        difficulty
        made_by
    }
}
```

### Create Exercise 3
```graphql
mutation {
    CreateExercise(exercise: {
        class_Id: "id-of-created-class",
        module_id: "id-of-created-module",
        name: "D of T - Oefening 3",
        question: "Zij (beslissen) ________ om morgen te gaan winkelen.",
        answers: [
            { value: "beslist", correct: false },
            { value: "beslissen", correct: true },
            { value: "beslisd", correct: false }
        ],
        difficulty: B2,
    }) {
        id
        class_Id
        module_id
        name
        question
        answers {
            value
            correct
        }
        difficulty
        made_by
    }
}
```

### Create Exercise 4
```graphql
mutation {
    CreateExercise(exercise: {
        class_Id: "id-of-created-class",
        module_id: "id-of-created-module",
        name: "D of T - Oefening 4",
        question: "De studenten (leren) ________ voor het examen.",
        answers: [
            { value: "leerd", correct: false },
            { value: "leren", correct: true },
            { value: "leert", correct: false }
        ],
        difficulty: B2,
    }) {
        id
        class_Id
        module_id
        name
        question
        answers {
            value
            correct
        }
        difficulty
        made_by
    }
}
```

### Create Result 1
```graphql
mutation {
    CreateResult(input: {
        exercise_id: "id-of-created-exercise"
        user_id: "id-of-user"
        class_id: "id-of-created-class",
        module_id: "id-of-created-module",
        input: "Answer-to-exercise-1"
        result: true
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

### Create Result 2
```graphql
   mutation {
     CreateResult(input: {
       exercise_id: "id-of-created-exercise"
       user_id: "id-of-user"
       class_id: "id-of-created-class"
       module_id: "id-of-created-module"
       input: "yourInput"
       result: true
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

### Create Result 3
```graphql
   mutation {
     CreateResult(input: {
       exercise_id: "id-of-created-exercise"
       user_id: "id-of-user"
       class_id: "id-of-created-class"
       module_id: "id-of-created-module"
       input: "yourInput"
       result: true
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

### Create Result 4
```graphql
   mutation {
     CreateResult(input: {
       exercise_id: "id-of-created-exercise"
       user_id: "id-of-user"
       class_id: "id-of-created-class"
       module_id: "id-of-created-module"
       input: "yourInput"
       result: true
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

### Get All results
```graphql
query GetAllChildrenForModule {
    getModule(id: "1230c264-b961-4b47-977e-871c706cbb8a") {
        id
        school_id
        name
        description
        difficulty
        category
        made_by
        private
        key
        created_at
        updated_at
        soft_deleted
    }
    listClasses(
        filter: {
          module_id: "1230c264-b961-4b47-977e-871c706cbb8a"
        },
        paginate: {
          Step: 0,
          amount: 4
        }) {
        id
        module_Id
        name
        description
        difficulty
        made_by
    }
    ListExercise(
       filter: {
         class_Id: "643fb675-946f-48d5-8ea4-aea4772aafbf"
       }
       paginator: {
         amount: 10
         Step: 1
       }
     ) {
       id
       class_Id
       module_id
       name
       question
       answers {
         value
         correct
       }
       difficulty
       made_by
     }
     ListResults(
       filter: {
         classId: "643fb675-946f-48d5-8ea4-aea4772aafbf" 
       }
       paginator: {
         amount: 10
         Step: 0
       }
     ) {
       id
       exercise_id
       user_id
       class_id
       module_id
       input
       result
     }
}
```