
1. **List Results Query:**
   ```graphql
   {
     ListResults(
       filter: {
         softDelete: true
         exerciseId: "yourExerciseId"
         userId: "yourUserId"
         classId: "yourClassId"
         moduleId: "yourModuleId"
         input: "yourInput"
         result: true
       }
       paginator: {
         amount: 10
         step: 1
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

2. **Get Result by ID Query:**
   ```graphql
   {
     GetResultsByID(id: "yourResultId") {
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

3. **Create Result Mutation:**
   ```graphql
   mutation {
     CreateResult(input: {
       exercise_id: "yourExerciseId"
       user_id: "yourUserId"
       class_id: "yourClassId"
       module_id: "yourModuleId"
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

4. **Update Result Mutation:**
   ```graphql
   mutation {
     UpdateResult(id: "yourResultId", input: {
       exercise_id: "yourUpdatedExerciseId"
       user_id: "yourUpdatedUserId"
       class_id: "yourUpdatedClassId"
       module_id: "yourUpdatedModuleId"
       input: "yourUpdatedInput"
       result: false
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

5. **Delete Result Mutation:**
   ```graphql
   mutation {
     DeleteResult(id: "yourResultId")
   }
   ```
