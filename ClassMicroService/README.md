## Auth Headers (REQUIRED)
```
{
    "Authorization": "Bearer <token>"
}
```

## Create Class
```graphql
mutation CreateClass {
    createClass(input: {
        module_Id: "c0c60670-9293-11ee-b9d1-0242ac120002"
        name: "Sample Class",
        description: "sample description"
        difficulty: B2,
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
    listClasses(
        filter: {
            name: {
                input: "Class"
                type: ends
            }
            module_id: "c0c60670-9293-11ee-b9d1-0242ac120002"
            difficulty: B2
        },
        paginate: {
            Step: 0,
            amount: 4
        }) {
        id
        module_Id
        name
        description
        difficulty
        made_by
    }
}
```

## Get all classes (ADMIN)
```graphql
query ListClasses {
    listClasses(
        filter: {
            name: {
                input: "Class"
                type: ends
            }
            module_id: "c0c60670-9293-11ee-b9d1-0242ac120002"
            difficulty: B2
            softDelete: false
            made_by: "id-of-creator"
        },
        paginate: {
            Step: 0,
            amount: 4
        }) {
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