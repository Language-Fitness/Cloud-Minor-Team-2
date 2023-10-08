import { ApolloServer } from '@apollo/server';
import { startStandaloneServer } from '@apollo/server/standalone';

const typeDefs = `#graphql

#type User {
#    user_id: String
#    firstname: String
#    lastname: String
#    email: String
#    password: String
#    school_id: String
#    whitelist_module: [String]
#    role_id: String
#    rating: Int
#    settings: [String]
#    created_at: String
#    updated_at: String
#    soft_deleted: Boolean
#}
#
#type Role {
#    role_id: String
#    name: String
#    permissions: [String]
#    created_at: String
#    updated_at: String
#    soft_deleted: Boolean
#}

type School {
    id: String
    name: String
    location: String
    created_at: String
    updated_at: String
    soft_deleted: Boolean
}

type Module {
    id: String
    name: String
    description: String
    difficulty: Int
    category: String
    made_by: String
    private: Boolean
    key: String
    created_at: String
    updated_at: String
    soft_deleted: Boolean
}

type Class {
    id: String
    module_id: String
    name: String
    description: String
    difficulty: Int
    created_at: String
    updated_at: String
    soft_deleted: Boolean
}

type Exercise {
    id: String
    class_id: String
    name: String
    question: String
    answer: String
    question_type_id: String
    difficulty: Int
    created_at: String
    updated_at: String
    soft_deleted: Boolean
}

type Exercise_Type {
    id: String
    settings: [String]
    created_at: String
    updated_at: String
    soft_deleted: Boolean
}

type Result {
    id: String
    exercise_id: String
    user_id: String
    class_id: String
    module_id: String
    input: String
    result: String
    created_at: String
    updated_at: String
    soft_deleted: Boolean
}

input ResultInput {
    exercise_id: String
    user_id: String
    class_id: String
    module_id: String
    input: String
    result: String
}

type Query {
#    users: [User]
#    roles: [Role]
    oneSchool(id: String!): School
    schools: [School]
    oneModule(id: String!): Module
    modules: [Module]
    oneClass(id: String!): Class
    classes: [Class]
    oneExercise(id: String!): Exercise
    exercises: [Exercise]
    oneExerciseType(id: String!): Exercise_Type
    exerciseTypes: [Exercise_Type]
    oneResult(id: String!): Result
    results: [Result]
}

type Mutation {
    createResult(input: ResultInput!): Result
}
`;

const schools = [
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        name: "Example School 1",
        location: "City A",
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: false
    },
    {
        id: "2",
        name: "Example School 2",
        location: "City B",
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: true
    }
];

const modules = [
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        name: "Module A",
        description: "This is Module A's description",
        difficulty: 3,
        category: "Science",
        made_by: "User123",
        private: false,
        key: "module123",
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: false
    },
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        name: "Module B",
        description: "This is Module B's description",
        difficulty: 2,
        category: "Math",
        made_by: "User456",
        private: true,
        key: "module456",
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: true
    }
];

const classes = [
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        module_id: "module123",
        name: "Class 101",
        description: "This is Class 101's description",
        difficulty: 2,
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: false
    },
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        module_id: "module456",
        name: "Class 202",
        description: "This is Class 202's description",
        difficulty: 3,
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: true
    }
];

const exercises = [
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        class_id: "class123",
        name: "Exercise 1",
        question: "What is 2 + 2?",
        answer: "4",
        question_type_id: "type1",
        difficulty: 2,
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: false
    },
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        class_id: "class456",
        name: "Exercise 2",
        question: "What is the capital of France?",
        answer: "Paris",
        question_type_id: "type2",
        difficulty: 3,
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: true
    }
];

const exerciseTypes = [
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        settings: ["Setting 1", "Setting 2", "Setting 3"],
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: false
    },
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        settings: ["Setting 4", "Setting 5"],
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: true
    }
];

const results = [
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        exercise_id: "exercise123",
        user_id: "user123",
        class_id: "class123",
        module_id: "module123",
        input: "2 + 2",
        result: "4",
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: false
    },
    {
        id: "750a64f0-660c-11ee-8c99-0242ac120002",
        exercise_id: "exercise456",
        user_id: "user456",
        class_id: "class456",
        module_id: "module456",
        input: "What is the capital of France?",
        result: "Paris",
        created_at: "2023-10-08",
        updated_at: "2023-10-08",
        soft_deleted: true
    }
];



const resolvers = {
    Query: {
        // School queries
        oneSchool: (parent, args) => {
            const { id } = args;
            const school = schools.find((school) => school.id === id);
            if (!school) {
                throw new Error(`School with ID ${id} not found`);
            }
            return school;},
        schools: () => schools,

        // Module queries
        oneModule: (parent, args) => {
            const { id } = args;
            const module = modules.find((module) => module.id === id);
            if (!module) {
                throw new Error(`Module with ID ${id} not found`);
            }
            return module;},
        modules: () => modules,

        // Class queries
        oneClass: (parent, args) => {
            const { id } = args;
            const oneClass = classes.find((oneClass) => oneClass.id === id);
            if (!oneClass) {
                throw new Error(`Class with ID ${id} not found`);
            }
            return oneClass;},
        classes: () => classes,

        // Exercise queries
        oneExercise: (parent, args) => {
            const { id } = args;
            const exercise = exercises.find((exercise) => exercise.id === id);
            if (!exercise) {
                throw new Error(`Exercise with ID ${id} not found`);
            }
            return exercise;},
        exercises: () => exercises,

        // Exercise Type queries
        oneExerciseType: (parent, args) => {
            const { id } = args;
            const exerciseType = exerciseTypes.find((exerciseType) => exerciseType.id === id);
            if (!exerciseType) {
                throw new Error(`ExerciseType with ID ${id} not found`);
            }
            return exerciseTypes;},
        exerciseTypes: () => exerciseTypes,

        // Result queries
        oneResult: (parent, args) => {
            const { id } = args;
            const result = results.find((result) => result.id === id);
            if (!result) {
                throw new Error(`Result with ID ${id} not found`);
            }
            return result;},
        results: () => results
    },
    Mutation: {
        // Result mutation
        createResult: (parent, args) => {
            const { input } = args;
            const newResult = {
                id: "750a64f0-660c-11ee-8c99-0242ac120002",
                exercise_id: input.exercise_id,
                user_id: input.user_id,
                class_id: input.class_id,
                module_id: input.module_id,
                input: input.input,
                result: input.result,
                created_at: new Date().toISOString(),
                updated_at: new Date().toISOString(),
                soft_deleted: false,
            };
            results.push(newResult)
            return newResult;
        },
    }
};

const server = new ApolloServer({
    typeDefs,
    resolvers,
});

const { url } = await startStandaloneServer(server, {
    listen: { port: 4000 },
});

console.log(`🚀  Server ready at: ${url}`);