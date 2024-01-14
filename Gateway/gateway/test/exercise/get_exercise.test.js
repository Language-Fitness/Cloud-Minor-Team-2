const axios = require('axios');
const { token, url, oldToken } = require('../const.js');

async function GetExerciseTest() {
    const exerciseId = "95f964a0-9749-4064-9162-cdd1b7b5d776";

    const postData = {
        query: `query {
          GetExercise(ExerciseId: "${exerciseId}") {
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
        }`
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        expect(responseData).toEqual({
            data: {
                GetExercise: {
                    id: "95f964a0-9749-4064-9162-cdd1b7b5d776",
                    class_Id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a4b",
                    module_id: "4bdaaf03-f5d0-43a9-a1d2-f5cc54ca7a8b",
                    name: "Updated Exercise",
                    question: "What is the updated question?",
                    answers: [
                        {
                            value: "Updated Option 1",
                            correct: true
                        },
                        {
                            value: "Updated Option 2",
                            correct: false
                        },
                        {
                            value: "Updated Option 3",
                            correct: false
                        }
                    ],
                    difficulty: "B1",
                    created_at: "2024-01-02T14:01:49Z",
                    updated_at: expect.any(String),
                    soft_deleted: false,
                    made_by: expect.any(String)
                }
            }
        });

        console.log('Test passed for GetExercise:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('GetExercise', () => {
    test('should return the expected exercise data', async () => {
        await GetExerciseTest();
    });
});


async function GetExerciseNotFoundTest() {
    const nonExistentExerciseId = "95f964a0-9749-4064-5162-cdd1b7b5d776";

    const postData = {
        query: `query {
          GetExercise(ExerciseId: "${nonExistentExerciseId}") {
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
        }`
    };

    const headers = {
        'Authorization': 'Bearer ' + token,
        'Content-Type': 'application/json',
    };

    try {
        const response = await axios.post(url, postData, { headers });
        const responseData = response.data;

        // Define the expected output for exercise not found
        const expectedOutput = {
            "errors": [
                {
                    "message": "exercise not found",
                }
            ],
            "data": {
                "GetExercise": null
            }
        };

        // Perform assertions based on the response data
        expect(responseData).toEqual(expectedOutput);

        console.log('Test passed for GetExercise not found:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('GetExercise Not Found', () => {
    test('should return error for non-existent exercise', async () => {
        await GetExerciseNotFoundTest();
    });
});

async function GetExerciseInvalidIdTest() {
    const postData = {
        query: `query {
            GetExercise(ExerciseId: "95f964a0") {
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
        }`,
        variables: {}
    };

    try {
        const response = await axios.post(url, postData, { headers: { 'Authorization': 'Bearer ' + token, 'Content-Type': 'application/json' }, validateStatus: () => true });
        const responseData = response.data;

        // Perform assertions based on the response data
        expect(responseData).toEqual({
            errors: [
                {
                    message: "Validation errors: ID :'95f964a0' is not a valid UUID",
                }
            ],
            data: {
                GetExercise: null
            }
        });

        console.log('Test passed with invalid ExerciseId:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('GetExercise with Invalid ExerciseId', () => {
    test('should return validation error for invalid ExerciseId', async () => {
        await GetExerciseInvalidIdTest();
    });
});
