## Create Class
```graphql
mutation CreateClass {
    createClass(input: {
        module_Id: "c0c60670-9293-11ee-b9d1-0242ac120002"
        name: "Sample Class",
        description: "sample description"
        difficulty: 1,
    }) {
        id
        name
        description
        difficulty
        made_by
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
        made_by
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
        module_Id
        name
        description
        difficulty
        made_by
    }
}
```

## Update a class
```graphql
mutation UpdateClass {
    updateClass(
        id: "your-class-id-to-update",
        input: {
            module_Id: "c0c60670-9293-11ee-b9d1-0242ac120002"
            name: "Updated Module Name",
            description: "Updated description",
            difficulty: 2,
        }
    ) {
        id
        module_Id
        name
        description
        difficulty
        made_by
        created_at
        updated_at
        soft_deleted
    }
}
```

## Delete a class
```graphql
mutation DeleteClass {
    deleteClass(id: "your-class-id-to-delete")
}
```