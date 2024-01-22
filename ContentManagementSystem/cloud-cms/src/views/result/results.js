export const headers = [
    {title: 'Exercise', align: 'start', sortable: false, key: 'exercise_id'},
    {title: 'User', align: 'start', sortable: false, key: 'user_id'},
    {title: 'Input', key: 'input', align: 'start', sortable: false},
    {title: 'Result', key: 'result', align: 'start', sortable: false},
]

export const listResultQuery = `
    query Query($filter: ResultFilter!, $paginator: Paginator!) {
      ListResults(filter: $filter, paginator: $paginator) {
        id
        module_id
        class_id
        exercise_id
        user_id
        input
        result
      }
    }
`;