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
    username: 'Merlijn@student.com',
    groups: ['permissions_student']
}

export const newStudentAttributes = {
    school_id: 'de4447f4-8855-11ee-b9d1-0242ac120002',
    whitelist_module: JSON.stringify([]),
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
    username: 'bram@teacher.com',
    groups: ['permissions_student', 'permissions_teacher']
}

export const newTeacherAttributes = {
    school_id: 'de4447f4-8855-11ee-b9d1-0242ac120002',
    whitelist_module: JSON.stringify([]),
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
    username: 'admin@admin.com',
    groups: ['permissions_student', 'permissions_teacher', 'permissions_administrator']
}

export const newAdminAttributes = {
    school_id: 'de4447f4-8855-11ee-b9d1-0242ac120002',
    whitelist_module: JSON.stringify([]),
    role_id: 'a47bae98-8857-11ee-b9d1-0242ac120002',
    created_at: '2023-01-01T00:00:00Z',
    updated_at: '2023-01-02T12:34:56Z',
    soft_deleted: false
}

export const newClientData = {
    clientId: 'user-management-client',
    enabled: true,
    standardFlowEnabled: false,
    implicitFlowEnabled: false,
    directAccessGrantsEnabled: false,
    serviceAccountsEnabled: true,
    authorizationServicesEnabled: true,
};

export const MicroserviceClients = [
    {
        clientId: 'login-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: true,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },
    {
        clientId: 'module-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: false,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },
    {
        clientId: 'class-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: false,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },
    {
        clientId: 'exercise-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: false,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },
    {
        clientId: 'result-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: false,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },
    {
        clientId: 'leaderboard-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: false,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },
    {
        clientId: 'school-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: false,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },    {
        clientId: 'openai-client',
        enabled: true,
        standardFlowEnabled: false,
        implicitFlowEnabled: false,
        directAccessGrantsEnabled: false,
        serviceAccountsEnabled: true,
        authorizationServicesEnabled: true,
    },
];
