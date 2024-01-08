
1. **Get Exercise by ID:**
   ```graphql
   query {
     GetExercise(ExerciseId: "yourExerciseId") {
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
       created_at
       updated_at
       soft_deleted
       made_by
     }
   }
   ```

2. **List Exercises with Filtering and Pagination:**
   ```graphql
   query {
     ListExercise(
       filter: {
         name: "yourExerciseName"
         softDelete: true
         difficulty: A1
         class_Id: "yourClassId"
         module_id: "yourModuleId"
         made_by: "yourMadeById"
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
   }
   ```

3. **Create Exercise:**
   ```graphql
   mutation {
     CreateExercise(exercise: {
       class_Id: "yourClassId"
       module_id: "yourModuleId"
       name: "New Exercise"
       question: "What is the question?"
       answers: [
         { value: "Option 1", correct: true },
         { value: "Option 2", correct: false },
         { value: "Option 3", correct: false }
       ]
       difficulty: A2
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

4. **Update Exercise:**
   ```graphql
   mutation {
     UpdateExercise(id: "yourExerciseId", exercise: {
       class_Id: "yourUpdatedClassId"
       module_id: "yourUpdatedModuleId"
       name: "Updated Exercise"
       question: "What is the updated question?"
       answers: [
         { value: "Updated Option 1", correct: true },
         { value: "Updated Option 2", correct: false },
         { value: "Updated Option 3", correct: false }
       ]
       difficulty: B1
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

5. **Delete Exercise:**
   ```graphql
   mutation {
     DeleteExercise(ExerciseID: "yourExerciseId")
   }
   ```