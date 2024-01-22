
1. **Generate Multiple Choice Questions:**
   ```graphql
   mutation{
     generateMultipleChoiceQuestions(questionSubject:woordenschat questionLevel:B2, amountQuestions: 5) {
       response{
        status
        message
        token
    }
   }
   ```
   questionSubject options [grammatica, spelling, woordenschat, uitdrukkingen, interpunctie, werkwoordvervoegingen ]
   questionLevel options [A1, A2, B1, B2, C1, C2]
   

2. **Retrieve Generated Multiple Choice Questions:**
   ```graphql
   query{
    retrieveMultipleChoiceQuestions(token: "Your token"){
     status
     message
     questions{
         questionType
         questionSubject
         questionLevel
         questionText
         answerOptions
         correctAnswer
         }
     }
   }
   ```

3. **Generate Explanation:**
   ```graphql
   mutation{
    generateExplanation(questionSubject: woordenschat questionText: "Question Text" givenAnswer: "Given answer" correctAnswer: "Correct answer"){
     response{
       status
       message
       token
    }
    }
   ```
   questionSubject options [grammatica, spelling, woordenschat, uitdrukkingen, interpunctie, werkwoordvervoegingen ]

4. **Retrieve Explanation:**
   ```graphql
   query{
   retrieveExplanation(token: "Your token"){
    status
    message
      explanation{
            info
            tips
       }
     }
   }
   ```

5. **Generate Questions From File:**
   ```graphql
   mutation{
      readMultipleChoiceQuestionsFromFile(filename: "Your file(.docx/.pdf)" fileData: "Your file Data in Base64"){
      response{
       status
       message
       token
      }
     }
   }
   ```

6. **Read Multiple Choice Questions From File:**
   ```graphql
   query{
      retrieveMultipleChoiceQuestionsFromFile(token: "Your token"){
       status
       message
       questions{
         questionType
         questionSubject
         questionLevel
         questionText
         answerOptions
         correctAnswer
       }
     }
   }
   ```