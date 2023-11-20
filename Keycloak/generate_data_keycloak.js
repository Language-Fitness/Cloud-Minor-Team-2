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

export const newClientData = {
    clientId: 'demo-client',
    enabled: true,
    standardFlowEnabled: false,
    implicitFlowEnabled: false,
    directAccessGrantsEnabled: false,
    serviceAccountsEnabled: true,
    authorizationServicesEnabled: true,
};