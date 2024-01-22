export const headers = [
    {title: 'Name', align: 'start', sortable: false, key: 'name',},
    {title: 'Description', key: 'description', align: 'start', sortable: false},
    {title: 'Difficulty', key: 'difficulty', align: 'start', sortable: false},
    {title: 'Category', key: 'category', align: 'start', sortable: false},
    {title: 'Actions', key: 'actions', align: 'start', sortable: false},
    {title: 'Classes', key: 'classes', align: 'center', sortable: false},
]

export const categories = ['Grammatica', 'Spelling', 'Woordenschat', 'Werkwoordspelling', 'Uitdrukkingen', 'Interpunctie', 'Werkwoordvervoegingen', 'Fast_Track']

export const difficulties = ['A1', 'A2', 'B1', 'B2', 'C1', 'C2']

export const listModulesQuery = `
    query ListModules($filter: ModuleFilter, $paginate: Paginator) {
      listModules(filter: $filter, paginate: $paginate) {
        id
        name
        school_id
        description
        category
        difficulty
        made_by
        private
      }
    }
`;

export const deleteModuleQuery = `
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

export const createModuleQuery = `
    mutation CreateModule($input: ModuleInputCreate!) {
      createModule(input: $input) {
        id
        name
        description
        school_id
        category
        difficulty
        made_by
        private
        key
      }
    }
`;