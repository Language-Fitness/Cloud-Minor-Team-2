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
const projectRealm          = 'Test2';
const adminClientId         = 'admin-cli';
const adminClientSecret     = 'w9Wxse58YrhGX9VQcb4jFDm2sv43Pyjg';
const adminTokenEndpoint    = `${keycloakServer}/realms/${masterRealm}/protocol/openid-connect/token`;
const clientTokenEndpoint   = `${keycloakServer}/realms/${projectRealm}/protocol/openid-connect/token`;

async function init() {
    // PREREQUISITES //
    // ---------------------------------------------------- //
    // After initial build of keycloak:
    // 1. Log into the admin console
    // 2. Adjust Capability config -> {
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
    await addRoleMappingsToUser(adminToken, masterRealm, adminCLIServiceAccount.id, adminCLIClient.id, [createRealmRole])

    // Create a new realm and a new client which is to be used to add new users to the realm
    await createRealm(adminToken, projectRealm)
    await createClientInRealm(adminToken, projectRealm, newClientData);


    // PROJECT REALM //
    // ---------------------------------------------------- //
    // Get the created client and realm management client (to fetch role which is to be assigned to the created client.)
    const clients               = await getAllClientsInRealm(adminToken, projectRealm);
    const realmManagementClient = clients.find(client => client.clientId === 'realm-management');
    const createdClient         = clients.find(client => client.clientId === newClientData.clientId)
    const serviceAccount        = await getServiceAccountUser(adminToken, projectRealm, newClientData.clientId)

    // Get all roles in realm, filter for manage-users role, and assign to service account of created client.
    const roles                 = await getAllClientRolesInRealm(adminToken, projectRealm, realmManagementClient.id);
    const manageUsersRole       = roles.find(role => role.name === 'manage-users');
    await addRoleMappingsToUser(adminToken, projectRealm, serviceAccount.id, realmManagementClient.id, [manageUsersRole])

    // Log into the created client and get the bearer token. Add new users using this client
    const clientToken = await loginAsClient(createdClient.secret)
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


async function addRoleMappingsToUser(accessToken, realmName, userId, clientId, roleMappings) {
    const userRoleMappingApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/users/${userId}/role-mappings/clients/${clientId}`;

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

async function addNewUser(accessToken, realmName, user, attributes) {
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

