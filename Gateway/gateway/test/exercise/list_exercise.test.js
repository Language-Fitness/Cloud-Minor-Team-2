const axios = require('axios');
const { token, url, oldToken } = require('../const.js');

async function ListExerciseTest() {

    const postData = {
        query: `query {
      ListExercise(
        filter: {
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
        expect(responseData.data.ListExercise.length).toBeGreaterThan(0);

        //console.log('Test passed:', responseData);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

// Run the test
describe('ListExercise', () => {
    test('should successfully list exercises', async () => {
        await ListExerciseTest();
    });
});

async function ListExerciseTestWithOldToken() {
    const postData = {
        query: `query {
      ListExercise(
        filter: {
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
    }`,
        variables: {}
    };

    const headersWithOldToken = {
        'Authorization': 'Bearer ' + oldToken,
        'Content-Type': 'application/json',
    };

    try {
        // Attempt with the old (invalid) token
        const responseWithOldToken = await axios.post(url, postData, { headers: headersWithOldToken });
        const responseDataWithOldToken = responseWithOldToken.data;

        // Perform assertions based on the response data for the old token
        expect(responseDataWithOldToken).toEqual({
            errors: [
                {
                    message: 'invalid token introspect',
                },
            ],
            data: {
                ListExercise: null,
            },
        });

        //console.log('Test passed with old (invalid) token:', responseDataWithOldToken);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

// only work if introspect is enabled
// describe('ListExercise with Old Token', () => {
//     test('should return invalid token introspect error', async () => {
//         await ListExerciseTestWithOldToken();
//     });
// });

async function ListExerciseTestWithNoToken() {
    const postData = {
        query: `query {
      ListExercise(
        filter: {
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
    }`,
        variables: {}
    };

    try {
        // Attempt with no token
        const responseWithNoToken = await axios.post(url, postData);
        const responseDataWithNoToken = responseWithNoToken.data;

        expect(responseDataWithNoToken).toEqual({
            errors: [
                {
                    message: "invalid token format: null",
                },
            ],
            data: {
                ListExercise: null,
            },
        });

        //console.log('Test passed with no token:', responseDataWithNoToken);
    } catch (error) {
        console.error('Test failed:', error.message);
        throw error;
    }
}

describe('ListExercise with No Token', () => {
    test('should return invalid token introspect error', async () => {
        await ListExerciseTestWithNoToken();
    });
});
