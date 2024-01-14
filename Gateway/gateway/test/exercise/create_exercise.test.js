const axios = require('axios');
const { token, url, oldToken } = require('../const.js');


async function CreateExerciseTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers: headers });
        const responseData = response.data;
        console.log(response);

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            data: {
                CreateExercise: {
                    id: expect.any(String),
                    class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b",
                    module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b",
                    name: "New Exercise",
                    question: "What is the question?",
                    answers: [
                        { value: "Option 1", correct: true },
                        { value: "Option 2", correct: false },
                        { value: "Option 3", correct: false }
                    ],
                    difficulty: "A2",
                    made_by: expect.any(String)
                }
            }
        });

        console.log('Test passed:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise ', () => {
    test('should create a new exercise', async () => {
        await CreateExerciseTest();
    });
});

async function CreateExerciseInvalidClassIdTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "sdgfd"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Validation errors: ClassID :'sdgfd' is not a valid UUID"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with invalid ClassId:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with Invalid ClassId', () => {
    test('should return validation error for invalid ClassId', async () => {
        await CreateExerciseInvalidClassIdTest();
    });
});


async function CreateExerciseInvalidModuleIdTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "sdgfd"
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Validation errors: ModuleID :'sdgfd' is not a valid UUID"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with invalid ModuleId:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with Invalid ModuleId', () => {
    test('should return validation error for invalid ModuleId', async () => {
        await CreateExerciseInvalidModuleIdTest();
    });
});

async function CreateExerciseInvalidNameTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
                name: "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Validation errors: Name length should be less than 50"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with invalid Name:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with Invalid Name', () => {
    test('should return validation error for invalid Name', async () => {
        await CreateExerciseInvalidNameTest();
    });
});

async function CreateExerciseInvalidQuestionTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
                name: "New Exercise"
                question: "What is the question?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????"
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Validation errors: Question length should be less than 100"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with invalid Question:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with Invalid Question', () => {
    test('should return validation error for invalid Question', async () => {
        await CreateExerciseInvalidQuestionTest();
    });
});

async function CreateExerciseNoIncorrectAnswersTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
                name: "New Exercise"
                question: "What is the question?"
                answers: [
                  { value: "Option 1", correct: true },
                  { value: "Option 3", correct: true }
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "only one answer can be correct"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with invalid Answers:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with no incorrect answers', () => {
    test('should return validation error for invalid Answers', async () => {
        await CreateExerciseNoIncorrectAnswersTest();
    });
});

async function CreateExerciseNoCorrectAnswersTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
                name: "New Exercise"
                question: "What is the question?"
                answers: [
                  { value: "Option 1", correct: false },
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "at least one answer must be correct"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with invalid Answers:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with no correct answers', () => {
    test('should return validation error for invalid Answers', async () => {
        await CreateExerciseNoCorrectAnswersTest();
    });
});

async function CreateExerciseNoAnswersTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
                name: "New Exercise"
                question: "What is the question?"
                answers: []
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "exercise must have at least two answers"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with invalid Answers:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with no answers', () => {
    test('should return validation error for invalid Answers', async () => {
        await CreateExerciseNoAnswersTest();
    });
});

async function CreateExerciseInvalidDifficultyTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
                name: "New Exercise"
                question: "What is the question?"
                answers: [
                  { value: "Option 1", correct: true },
                  { value: "Option 2", correct: false },
                  { value: "Option 3", correct: false }
                ]
                difficulty: A5
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
            }`,
        variables: {}
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };
    try {
        const response = await axios.post(url, postData, { headers, validateStatus: () => true}, );
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message:  "Value \"A5\" does not exist in \"LanguageLevel\" enum. Did you mean the enum value \"A1\" or \"A2\"?",
                }
            ],
        });

        console.log('Test passed with invalid Difficulty:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with Invalid Difficulty', () => {
    test('should return validation error for invalid Difficulty', async () => {
        await CreateExerciseInvalidDifficultyTest();
    });
});

async function CreateExerciseNoTokenTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
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
            }`,
        variables: {}
    };
    try {
        const response = await axios.post(url, postData, { validateStatus: () => true });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "invalid token format: null",
                },
            ],
            data: {
                CreateExercise: null,
            },
        });

        console.log('Test passed with no token:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with No Token', () => {
    test('should return invalid token introspect error', async () => {
        await CreateExerciseNoTokenTest();
    });
});

async function CreateExerciseClassIdAsIntTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: 1
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
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
              }
            }`,
        variables: {}
    };
    try {
        const response = await axios.post(url, postData, { headers: { 'Authorization': 'Bearer ' + token, 'Content-Type': 'application/json' } });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Validation errors: ClassID :'1' is not a valid UUID"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with ClassId as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with ClassId as int', () => {
    test('should return validation error for ClassId as int', async () => {
        await CreateExerciseClassIdAsIntTest();
    });
});

async function CreateExerciseModuleIdAsIntTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
                module_id: 1
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
                module_id
                name
                question
                answers {
                  value
                  correct
                }
                difficulty
              }
            }`,
        variables: {}
    };
    try {
        const response = await axios.post(url, postData, { headers: { 'Authorization': 'Bearer ' + token, 'Content-Type': 'application/json' } });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Validation errors: ModuleID :'1' is not a valid UUID"
                }
            ],
            data: {
                CreateExercise: null
            }
        });

        console.log('Test passed with ModuleId as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with ModuleId as int', () => {
    test('should return validation error for ModuleId as int', async () => {
        await CreateExerciseModuleIdAsIntTest();
    });
});

async function CreateExerciseNameAsIntTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b",
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b",
                name: 1,
                question: "What is the question?",
                answers: [
                  { value: "Option 1", correct: true },
                  { value: "Option 2", correct: false },
                  { value: "Option 3", correct: false },
                ],
                difficulty: A2,
              }) {
                id,
                name,
                question,
                answers {
                  value,
                  correct,
                },
                difficulty,
              },
            }`,
        variables: {}
    };
    try {
        const response = await axios.post(url, postData, { headers: { 'Authorization': 'Bearer ' + token, 'Content-Type': 'application/json' }, validateStatus: () => true });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                { message: "String cannot represent a non string value: 1",}
            ],
        });

        console.log('Test passed with Name as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with Name as int', () => {
    test('should return validation error for Name as int', async () => {
        await CreateExerciseNameAsIntTest();
    });
});

async function CreateExerciseQuestionAsIntTest() {
    const postData = {
        query: `mutation {
              CreateExercise(exercise: {
                class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b",
                module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b",
                name: "New Exercise",
                question: 1,
                answers: [
                  { value: "Option 1", correct: true },
                  { value: "Option 2", correct: false },
                  { value: "Option 3", correct: false },
                ],
                difficulty: A2,
              }) {
                id,
                name,
                question,
                answers {
                  value,
                  correct,
                },
                difficulty,
              },
            }`,
        variables: {}
    };
    try {
        const response = await axios.post(url, postData, { headers: { 'Authorization': 'Bearer ' + token, 'Content-Type': 'application/json' }, validateStatus: () => true });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                { message: "String cannot represent a non string value: 1",}
            ],
        });

        console.log('Test passed with Question as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('CreateExercise with Question as int', () => {
    test('should return validation error for Question as int', async () => {
        await CreateExerciseQuestionAsIntTest();
    });
});