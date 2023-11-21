import {
    newAdminAttributes,
    newAdminData,
    newClientData,
    newStudentAttributes,
    newStudentData, newTeacherAttributes,
    newTeacherData
} from "./generate_data_keycloak.js";


const keycloakServer        = 'http://localhost:8888';
const masterRealm           = 'master';
const projectRealm          = 'cloud-project';
const adminClientId         = 'admin-cli';
const adminClientSecret     = 'LbB9TqSh9CDaD81pX5gMcre5laxaC6J0';
const adminTokenEndpoint    = `${keycloakServer}/realms/${masterRealm}/protocol/openid-connect/token`;
const clientTokenEndpoint   = `${keycloakServer}/realms/${projectRealm}/protocol/openid-connect/token`;

async function init() {
    // PREREQUISITES //
    // ---------------------------------------------------- //
    // After initial build of keycloak:
    // 1. Log into the admin console
    // 2. Adjust Capability config of admin-cli client -> {
    //      Client authentication   = On
    //      Authorization           = On
    //      Direct access grants    = Unchecked
    //      }
    // 3. Give admin-cli the 'admin' role mapping
    // 4. Copy Client Secret to global variable 'adminClientSecret' in this file


    // MASTER REALM //
    // ---------------------------------------------------- //
    // Get admin token to perform initial realm setup
    const adminToken = await loginAsAdminWithTheAdminCli()

    // Get admin CLI client and corresponding service account
    const adminClients              = await getAllClientsInRealm(adminToken, masterRealm);
    const adminCLIClient            = adminClients.find(client => client.clientId === 'admin-cli')
    const adminCLIServiceAccount    = await getServiceAccountUser(adminToken, masterRealm, 'admin-cli')

    // Get all roles in realm, filter for create-realm, and assign to service account of admin-cli
    const adminRoles                = await getAllClientRolesInRealm(adminToken, masterRealm, adminCLIClient.id);
    const createRealmRole           = adminRoles.find(role => role.name === 'create-realm');
    await addRoleMappingsToEntity(adminToken, masterRealm, 'users', adminCLIServiceAccount.id, adminCLIClient.id, [createRealmRole])


    // PROJECT REALM //
    // ---------------------------------------------------- //
    // Create a new realm and a new client which is to be used to add new users to the realm
    await createRealm(adminToken, projectRealm)
    await createClientInRealm(adminToken, projectRealm, newClientData);

    // Get the created client and realm management client (to fetch role which is to be assigned to the created client.)
    const clients               = await getAllClientsInRealm(adminToken, projectRealm);
    const realmManagementClient = clients.find(client => client.clientId === 'realm-management');
    const createdClient         = clients.find(client => client.clientId === newClientData.clientId)
    const serviceAccount        = await getServiceAccountUser(adminToken, projectRealm, newClientData.clientId)

    // Create the roles for teacher, student and admin in the client scope
    await createRole(adminToken, projectRealm, createdClient.id, 'student')
    await createRole(adminToken, projectRealm, createdClient.id, 'teacher')
    await createRole(adminToken, projectRealm, createdClient.id, 'admin')
    // this should be done with the roles defined in the roles directory
    // -> looped through files and added to keycloak roles

    // Create groups in the project realm
    await addNewGroup(adminToken, projectRealm, 'group_students')
    await addNewGroup(adminToken, projectRealm, 'group_teachers')
    await addNewGroup(adminToken, projectRealm, 'group_administrators')

    // Fetch all groups to use to assign the roles to.
    const groups = await getGroups(adminToken, projectRealm);
    const studentGroup = groups.find(group => group.name === 'group_students')
    const teacherGroup = groups.find(group => group.name === 'group_teachers')
    const adminGroup = groups.find(group => group.name === 'group_administrators')

    // Get all roles in realm for realm-client, filter for manage-users role, and assign to service account of created client
    const realmManagementRoles  = await getAllClientRolesInRealm(adminToken, projectRealm, realmManagementClient.id);
    const manageUsersRole       = realmManagementRoles.find(role => role.name === 'manage-users');
    await addRoleMappingsToEntity(adminToken, projectRealm, 'users', serviceAccount.id, realmManagementClient.id, [manageUsersRole])

    // Get all roles in realm for created-client, filter for designated roles
    const userRoles     = await getAllClientRolesInRealm(adminToken, projectRealm, createdClient.id);
    const studentRole   = userRoles.find(role => role.name === 'student');
    const teacherRole   = userRoles.find(role => role.name === 'teacher');
    const adminRole     = userRoles.find(role => role.name === 'admin');

    // Add  designated roles to student group
    await addRoleMappingsToEntity(
        adminToken,
        projectRealm,
        'groups',
        studentGroup.id,
        createdClient.id,
        [studentRole,teacherRole,adminRole]
    )

    // Add  designated roles to teacher group
    await addRoleMappingsToEntity(
        adminToken,
        projectRealm,
        'groups',
        teacherGroup.id,
        createdClient.id,
        [studentRole,teacherRole,adminRole]
    )

    // Add  designated roles to admin group
    await addRoleMappingsToEntity(
        adminToken,
        projectRealm,
        'groups',
        adminGroup.id,
        createdClient.id,
        [studentRole,teacherRole,adminRole]
    )

    // Log into the created client and get the bearer token. Add new users using this client
    const clientToken = await loginAsClient(createdClient.secret);
    await addNewUser(clientToken, projectRealm, newStudentData, newStudentAttributes);
    await addNewUser(clientToken, projectRealm, newTeacherData, newTeacherAttributes);
    await addNewUser(clientToken, projectRealm, newAdminData, newAdminAttributes);
}

init().then(r => console.log('finished'));

async function loginAsAdminWithTheAdminCli() {
    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `grant_type=client_credentials&client_id=${adminClientId}&client_secret=${adminClientSecret}`,
    };

    const response = await fetch(adminTokenEndpoint, requestData);
    const data = await response.json();
    return data.access_token;
}

async function loginAsClient(secret) {
    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `grant_type=client_credentials&client_id=${newClientData.clientId}&client_secret=${secret}`,
    };

    const response = await fetch(clientTokenEndpoint, requestData);
    const data = await response.json().catch(error => {
        console.log(error)
    });
    return data.access_token;
}

async function createClientInRealm(accessToken, realmName, newClientData) {
    const adminApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients`;

    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(newClientData),
    };

    await fetch(adminApiEndpoint, requestData)
}

async function createRealm(accessToken, newRealmName) {
    const adminApiEndpoint = `${keycloakServer}/admin/realms`;

    const newRealmData = {
        realm: newRealmName,
        enabled: true,
    };

    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(newRealmData),
    };

    await fetch(adminApiEndpoint, requestData)
}

async function createRole(accessToken, realmname, clientId, roleName) {
    const createRoleEndpoint = `${keycloakServer}/admin/realms/${realmname}/clients/${clientId}/roles`;

    const newRoleData = {
        name: roleName,
    }

    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(newRoleData)
    };

    await fetch(createRoleEndpoint, requestData)
}

async function getServiceAccountUser(accessToken, realmName, clientId) {
    const userApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/users`;

    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    const response = await fetch(`${userApiEndpoint}?username=service-account-${clientId}`, requestData);
    const serviceAccountUserData = await response.json();
    return serviceAccountUserData[0];
}


async function addRoleMappingsToEntity(accessToken, realmName, entity, entityId, clientId, roleMappings) {
    const userRoleMappingApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/${entity}/${entityId}/role-mappings/clients/${clientId}`;

    const roleMappingRequestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(roleMappings),
    };

    return await fetch(userRoleMappingApiEndpoint, roleMappingRequestData);
}

async function getAllClientRolesInRealm(accessToken, realmName, clientId) {
    const clientRolesApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients/${clientId}/roles`;

    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    const response = await fetch(clientRolesApiEndpoint, requestData);
    return await response.json();
}

async function getAllClientsInRealm(accessToken, realmName) {
    const clientsApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients`;

    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    const response = await fetch(clientsApiEndpoint, requestData);
    return await response.json();
}

async function addNewGroup(accessToken, realmName, groupName) {
    const createGroupEndpoint = `${keycloakServer}/admin/realms/${realmName}/groups`;

    const createGroupRequestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify({
            name: groupName
        }),
    };

    return await fetch(createGroupEndpoint, createGroupRequestData)
}

async function addNewUser(accessToken, realmName, user, attributes, clientId, role) {
    const createUserEndpoint = `${keycloakServer}/admin/realms/${realmName}/users`;

    user.attributes = attributes;

    const createUserRequestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(user),
    };

    return await fetch(createUserEndpoint, createUserRequestData)
}

async function getGroups(accessToken, realmName) {
    const getGroupsEndpoint = `${keycloakServer}/admin/realms/${realmName}/groups`;

    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    const response = await fetch(getGroupsEndpoint, requestData);
    return await response.json();
}

