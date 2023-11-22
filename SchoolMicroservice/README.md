## Create School
```graphql
mutation CreateSchool{
    createSchool(input: {
        name: "Sample School",
        location: "Sample location",
    }) {
        id
        name,
        location,
        created_at
        updated_at
        soft_deleted
    }
}
```

## Get one School
```graphql
query GetSchool {
    getSchool(id: "your-school-id") {
        id
        name
        location,
        created_at
        updated_at
        soft_deleted
    }
}
```

## Get all schools
```graphql
query ListSchools {
  listSchools {
    id,
    name,
    location,
    created_at,
    updated_at,
    soft_deleted
  }
}
```

## Update a school
```graphql
mutation UpdateSchool {
    updateSchool(
        id: "your-school-id-to-update",
        input: {
            name: "Updated School Name",
            location: "Updated School Location",
        }
    ) {
        id
        name,
        location,
        created_at,
        updated_at,
        soft_deleted
    }
}
```

## Delete a school
```graphql
mutation DeleteSchool {
    deleteSchool(id: "your-school-id")
}
```