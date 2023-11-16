## Create Class
```graphql
mutation CreateClass {
    createClass(input: {
        module_Id: "module-id"
        name: "Sample Class",
        description: "This is a sample class.",
        difficulty: 1,
    }) {
        id
        name
        description
        difficulty
        created_at
        updated_at
        soft_deleted
    }
}
```

## Get one Class
```graphql
query GetClass {
    getClass(id: "your-class-id") {
        id
        module_Id
        name
        description
        difficulty
        created_at
        updated_at
        soft_deleted
    }
}
```

## Get all classes
```graphql
query ListClasses {
    listClasses {
        id
        name
        description
        difficulty
        created_at
        updated_at
        soft_deleted
    }
}
```

## Update a class
```graphql
mutation UpdateClass {
    updateClass(
        id: "your-class-id-to-update",
        input: {
            module_Id: "updated class-id"
            name: "Updated Class Name",
            description: "Updated description",
            difficulty: 2,
        }
    ) {
        id
        module_Id
        name
        description
        difficulty
        created_at
        updated_at
        soft_deleted
    }
}
```

## Delete a class
```graphql
mutation DeleteClass {
    deleteClass(id: "your-class-id")
}
```