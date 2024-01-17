const axios = require('axios');
const { token, url, oldToken } = require('../const.js');

async function UpdateExerciseTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d776";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
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
            data: {
                UpdateExercise: {
                    id: expect.any(String),
                    class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b",
                    module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b",
                    name: "Updated Exercise",
                    question: "What is the updated question?",
                    answers: [
                        { value: "Updated Option 1", correct: true },
                        { value: "Updated Option 2", correct: false },
                        { value: "Updated Option 3", correct: false }
                    ],
                    difficulty: "B1",
                    made_by: expect.any(String)
                }
            }
        });

        console.log('UpdateExerciseTest passed:', responseData);
    } catch (error) {
        console.error('UpdateExerciseTest failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise', () => {
    test('should update an existing exercise', async () => {
        await UpdateExerciseTest();
    });
});


async function UpdateExerciseInvalidClassIdTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d777";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "sdgfd"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
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
                UpdateExercise: null
            }
        });

        console.log('Test passed with invalid ClassId:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with Invalid ClassId', () => {
    test('should return validation error for invalid ClassId', async () => {
        await UpdateExerciseInvalidClassIdTest();
    });
});

async function UpdateExerciseInvalidModuleIdTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d778";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "sdgfd"
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
                UpdateExercise: null
            }
        });

        console.log('Test passed with invalid ModuleId:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with Invalid ModuleId', () => {
    test('should return validation error for invalid ModuleId', async () => {
        await UpdateExerciseInvalidModuleIdTest();
    });
});

async function UpdateExerciseInvalidNameTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d779";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
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
                UpdateExercise: null
            }
        });

        console.log('Test passed with invalid Name:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with Invalid Name', () => {
    test('should return validation error for invalid Name', async () => {
        await UpdateExerciseInvalidNameTest();
    });
});

async function UpdateExerciseInvalidQuestionTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d780";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "Updated Exercise"
            question: "?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????"
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
                UpdateExercise: null
            }
        });

        console.log('Test passed with invalid Question:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with Invalid Question', () => {
    test('should return validation error for invalid Question', async () => {
        await UpdateExerciseInvalidQuestionTest();
    });
});

async function UpdateExerciseNoIncorrectAnswersTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d781";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "Updated Exercise"
            question: "What is the updated question?"
            answers: [
              { value: "Updated Option 1", correct: true },
              { value: "Updated Option 3", correct: true }
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
                UpdateExercise: null
            }
        });

        console.log('Test passed with invalid Answers:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with no incorrect answers', () => {
    test('should return validation error for invalid Answers', async () => {
        await UpdateExerciseNoIncorrectAnswersTest();
    });
});

async function UpdateExerciseNoCorrectAnswersTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d782";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "Updated Exercise"
            question: "What is the updated question?"
            answers: [
              { value: "Updated Option 1", correct: false },
              { value: "Updated Option 3", correct: false }
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
                UpdateExercise: null
            }
        });

        console.log('Test passed with invalid Answers:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with no correct answers', () => {
    test('should return validation error for invalid Answers', async () => {
        await UpdateExerciseNoCorrectAnswersTest();
    });
});

async function UpdateExerciseNoAnswersTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d783";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "Updated Exercise"
            question: "What is the updated question?"
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
                UpdateExercise: null
            }
        });

        console.log('Test passed with invalid Answers:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with no answers', () => {
    test('should return validation error for invalid Answers', async () => {
        await UpdateExerciseNoAnswersTest();
    });
});

async function UpdateExerciseInvalidDifficultyTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d784";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "Updated Exercise"
            question: "What is the updated question?"
            answers: [
              { value: "Updated Option 1", correct: true },
              { value: "Updated Option 2", correct: false },
              { value: "Updated Option 3", correct: false }
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
        const response = await axios.post(url, postData, { headers, validateStatus: () => true });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Value \"A5\" does not exist in \"LanguageLevel\" enum. Did you mean the enum value \"A1\" or \"A2\"?"
                }
            ],
        });

        console.log('Test passed with invalid Difficulty:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with Invalid Difficulty', () => {
    test('should return validation error for invalid Difficulty', async () => {
        await UpdateExerciseInvalidDifficultyTest();
    });
});

async function UpdateExerciseNoTokenTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d785";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "Updated Exercise"
            question: "What is the updated question?"
            answers: [
              { value: "Updated Option 1", correct: true },
              { value: "Updated Option 2", correct: false },
              { value: "Updated Option 3", correct: false }
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
                    message: "invalid token introspect",
                }
            ],
            data: {
                UpdateExercise: null
            }
        });

        console.log('Test passed with no token:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with No Token', () => {
    test('should return invalid token introspect error', async () => {
        await UpdateExerciseNoTokenTest();
    });
});

async function UpdateExerciseClassIdAsIntTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d786";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: 1
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b"
            name: "Updated Exercise"
            question: "What is the updated question?"
            answers: [
              { value: "Updated Option 1", correct: true },
              { value: "Updated Option 2", correct: false },
              { value: "Updated Option 3", correct: false }
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
                    message: "Validation errors: ClassID :'1' is not a valid UUID"
                }
            ],
            data: {
                UpdateExercise: null
            }
        });

        console.log('Test passed with ClassId as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with ClassId as int', () => {
    test('should return validation error for ClassId as int', async () => {
        await UpdateExerciseClassIdAsIntTest();
    });
});

async function UpdateExerciseModuleIdAsIntTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d787";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b"
            module_id: 1
            name: "Updated Exercise"
            question: "What is the updated question?"
            answers: [
              { value: "Updated Option 1", correct: true },
              { value: "Updated Option 2", correct: false },
              { value: "Updated Option 3", correct: false }
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
                    message: "Validation errors: ModuleID :'1' is not a valid UUID"
                }
            ],
            data: {
                UpdateExercise: null
            }
        });

        console.log('Test passed with ModuleId as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with ModuleId as int', () => {
    test('should return validation error for ModuleId as int', async () => {
        await UpdateExerciseModuleIdAsIntTest();
    });
});

async function UpdateExerciseNameAsIntTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d788";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b",
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b",
            name: 1,
            question: "What is the updated question?",
            answers: [
              { value: "Updated Option 1", correct: true },
              { value: "Updated Option 2", correct: false },
              { value: "Updated Option 3", correct: false }
            ],
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
        const response = await axios.post(url, postData, { headers, validateStatus: () => true });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "String cannot represent a non string value: 1"
                }
            ],
        });

        console.log('Test passed with Name as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with Name as int', () => {
    test('should return validation error for Name as int', async () => {
        await UpdateExerciseNameAsIntTest();
    });
});

async function UpdateExerciseQuestionAsIntTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d788";

    const postData = {
        query: `mutation {
          UpdateExercise(id: "${exerciseId}", exercise: {
            class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b",
            module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b",
            name: "Updated Exercise",
            question: 1,
            answers: [
              { value: "Updated Option 1", correct: true },
              { value: "Updated Option 2", correct: false },
              { value: "Updated Option 3", correct: false }
            ],
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
        const response = await axios.post(url, postData, { headers, validateStatus: () => true });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "String cannot represent a non string value: 1"
                }
            ],
        });

        console.log('Test passed with Question as int:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('UpdateExercise with Question as int', () => {
    test('should return validation error for Question as int', async () => {
        await UpdateExerciseQuestionAsIntTest();
    });
});
