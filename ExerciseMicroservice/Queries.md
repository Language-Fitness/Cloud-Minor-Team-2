
1. **Get Exercise by ID:**
   ```graphql
   query {
     GetExercise(ExerciseId: "yourExerciseId") {
       id
       class_Id
       module_id
       name
       question
       answers
       pos_correct_answer
       difficulty
       created_at
       updated_at
       soft_deleted
       made_by
     }
   }
   ```

2. **List Exercises with Filters and Pagination:**
   ```graphql
   query {
     ListExercise(filter: { 
     name: "yourExerciseName", 
     difficulty: B1, class_Id: "yourClassId" }, 
     paginator: { 
        amount: 10, 
        Step: 1 }) 
   {
       id
       class_Id
       module_id
       name
       question
       answers
       pos_correct_answer
       difficulty
       created_at
       updated_at
       soft_deleted
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
       name: "NewExercise"
       question: "What is the question?"
       answers: "Option A,Option B,Option C"
       pos_correct_answer: 2
       difficulty: B2
     }) {
       id
       class_Id
       module_id
       name
       question
       answers
       pos_correct_answer
       difficulty
       created_at
       updated_at
       soft_deleted
       made_by
     }
   }
   ```

4. **Update Exercise:**
   ```graphql
   mutation {
     UpdateExercise(id: "yourExerciseId", exercise: {
       class_Id: "updatedClassId"
       module_id: "updatedModuleId"
       name: "UpdatedExercise"
       question: "What is the updated question?"
       answers: "Option A,Option B,Option C,Option D"
       pos_correct_answer: 3
       difficulty: C1
     }) {
       id
       class_Id
       module_id
       name
       question
       answers
       pos_correct_answer
       difficulty
       created_at
       updated_at
       soft_deleted
       made_by
     }
   }
   ```

5. **Delete Exercise:**
   ```graphql
   mutation {
     DeleteExercise(ExerciseID: "yourExerciseId") {
       id
       class_Id
       module_id
       name
       question
       answers
       pos_correct_answer
       difficulty
       created_at
       updated_at
       soft_deleted
       made_by
     }
   }
   ```
