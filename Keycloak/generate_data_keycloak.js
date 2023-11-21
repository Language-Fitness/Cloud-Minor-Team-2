export const newStudentData = {
    firstName: 'Merlijn',
    lastName: 'Busch',
    email: 'Merlijn@student.com',
    enabled: true,
    credentials: [{
        type: 'password',
        value: 'secret',
        temporary: false
    }],
    username: 'Merlijn@student.com'
}

export const newStudentAttributes = {
    school_id: 'de4447f4-8855-11ee-b9d1-0242ac120002',
    whitelist_module: JSON.stringify(['module1', 'module2', 'module3']),
    role_id: 'e6707d12-8855-11ee-b9d1-0242ac120002',
    rating: 4.5,
    settings: JSON.stringify({
        darkmode: true,
    }),
    created_at: '2023-01-01T00:00:00Z',
    updated_at: '2023-01-02T12:34:56Z',
    soft_deleted: false
}

export const newTeacherData = {
    firstName: 'Bram',
    lastName: 'Terlouw',
    email: 'bram@teacher.com',
    enabled: true,
    credentials: [{
        type: 'password',
        value: 'secret',
        temporary: false
    }],
    username: 'bram@teacher.com'
}

export const newTeacherAttributes = {
    school_id: 'de4447f4-8855-11ee-b9d1-0242ac120002',
    whitelist_module: JSON.stringify(['module1', 'module2', 'module3']),
    role_id: '7d083854-8857-11ee-b9d1-0242ac120002',
    rating: 4.5,
    settings: JSON.stringify({
        darkmode: true,
    }),
    created_at: '2023-01-01T00:00:00Z',
    updated_at: '2023-01-02T12:34:56Z',
    soft_deleted: false
}

export const newAdminData = {
    firstName: 'chad',
    lastName: 'admin',
    email: 'admin@admin.com',
    enabled: true,
    credentials: [{
        type: 'password',
        value: 'secret',
        temporary: false
    }],
    username: 'admin@admin.com'
}

export const newAdminAttributes = {
    school_id: 'de4447f4-8855-11ee-b9d1-0242ac120002',
    whitelist_module: JSON.stringify(['module1', 'module2', 'module3']),
    role_id: 'a47bae98-8857-11ee-b9d1-0242ac120002',
    rating: 4.5,
    settings: JSON.stringify({
        darkmode: true,
    }),
    created_at: '2023-01-01T00:00:00Z',
    updated_at: '2023-01-02T12:34:56Z',
    soft_deleted: false
}

export const newClientData = {
    clientId: 'demo-client',
    enabled: true,
    standardFlowEnabled: false,
    implicitFlowEnabled: false,
    directAccessGrantsEnabled: false,
    serviceAccountsEnabled: true,
    authorizationServicesEnabled: true,
};