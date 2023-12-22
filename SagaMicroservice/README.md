```graphql
mutation {
    deleteObject(filter: {
        soft_deleted: false,
        object_id: "example_id",
        object_type: School
    }) {
        id
        text
        status
        object_id
        object_type
    }
}
```