export const headers = [
    {title: 'Name', key: 'name', align: 'start', sortable: false},
    {title: 'Module_Id', align: 'start', sortable: false, key: 'module_Id',},
    {title: 'Description', key: 'description', align: 'start', sortable: false},
    {title: 'Difficulty', key: 'difficulty', align: 'start', sortable: false},
    {title: 'Actions', key: 'actions', align: 'start', sortable: false},
    {title: 'Exercises', key: 'exercises', align: 'center', sortable: false},
]

export const difficulties = ['A1', 'A2', 'B1', 'B2', 'C1', 'C2']

export const listClassesQuery = `
    query ListClasses($filter: ListClassFilter, $paginate: Paginator) {
      listClasses(filter: $filter, paginate: $paginate) {
        id
        name
        description
        difficulty
        module_Id
        made_by
      }
    }
`;

export const deleteClassQuery = `
    mutation DeleteObject($filter: SagaFilter) {
      deleteObject(filter: $filter) {
        id
        object_id
        object_type
        text
        status
      }
    }
`;

export const createClassQuery = `
    mutation CreateClass($input: ClassInput!) {
      createClass(input: $input) {
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
`;