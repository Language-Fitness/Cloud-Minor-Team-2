### Create Module Testing

```graphql
mutation CreateModule {
    createModule(input: {
        name: "Sample Module 2",
        school_id: "9154f86e-9eb5-11ee-8c90-0242ac120002"
        description: "This is a sample module.",
        difficulty: B2,
        category: Woordenschat,
        private: true,
        key: "sample-key"
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

### Get one Module
```graphql
query GetModule {
    getModule(id: "your-module-id") {
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

### Get all Module
```graphql
query ListModules {
    listModules(
        filter: {
            school_id: "<school-id>"
            made_by: "<user-id>"
            softDelete: false
            name: {
                input: "input"
                type: eq
            }
            difficulty: B2
            category: Woordenschat
            private: false
        },
        paginate: {
            amount: 10,
            Step: 0
        }
    ) {
        id
        school_id
        name
        description
        difficulty
        category
        made_by
        private
    }
}
```

### Update a module
```graphql
mutation UpdateModule {
    updateModule(
        id: "your-module-id-to-update",
        input: {
            name: "Updated Module Name",
            description: "Updated description",
            difficulty: C1,
            category: Werkwoordvervoegingen,
            private: true,
            key: "updated-key"
        }
    ) {
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

### Delete a module
```graphql
mutation DeleteModule {
    deleteModule(id: "your-module-id-to-delete")
}
```

