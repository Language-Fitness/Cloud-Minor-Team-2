### Create Module Testing

```graphql
mutation CreateModule {
    createModule(input: {
        name: "Sample Module 2",
        description: "This is a sample module.",
        difficulty: 1,
        category: "Sample Category",
        private: true,
        key: "sample-key"
    }) {
        id
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
    listModules {
        id
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
            difficulty: 2,
            category: "Updated Category",
            private: true,
            key: "updated-key"
        }
    ) {
        id
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

