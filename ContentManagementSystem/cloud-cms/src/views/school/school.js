export const headers = [
    {title: 'Name', key: 'name', align: 'start', sortable: false},
    {title: 'Location', key: 'location', align: 'start', sortable: false},
    {title: 'Made by', key: 'made_by', align: 'start', sortable: false},
    {title: 'Actions', key: 'actions', align: 'start', sortable: false},
]

export const listSchoolsQuery = `
    query ListSchools($filter: ListSchoolFilter, $paginate: Paginator) {
      listSchools(filter: $filter, paginate: $paginate) {
        has_openai_access
        id
        location
        made_by
        name
      }
    }
`;