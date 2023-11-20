import {newAdminData, newClientData, newStudentData, newTeacherData} from "./generate_data_keycloak.js";

const keycloakServer    = 'http://localhost:8888';
const masterRealm             = 'master';
const projectRealm         = 'Test2';
const adminClientId     = 'admin-cli';
const adminClientSecret = 'WQpcgyqH7FTFkFh0iGkQK2LOg430VV7t';
const adminTokenEndpoint     = `${keycloakServer}/realms/${masterRealm}/protocol/openid-connect/token`;
const clientTokenEndpoint     = `${keycloakServer}/realms/${projectRealm}/protocol/openid-connect/token`;

async function init() {
    const adminToken = await loginAsAdminWithTheAdminCli()
    await createRealm(adminToken, projectRealm)
    await createClientInRealm(adminToken, projectRealm, newClientData);

    const serviceAccount        = await getServiceAccountUser(adminToken, projectRealm, newClientData.clientId)

    const clients               = await getAllClientsInRealm(adminToken, projectRealm);
    const createdClient         = clients.find(client => client.clientId === newClientData.clientId)
    const realmManagementClient = clients.find(client => client.clientId === 'realm-management');

    const roles                 = await getAllClientRolesInRealm(adminToken, projectRealm, realmManagementClient.id);
    const manageUsersRole       = roles.find(role => role.name === 'manage-users');
    const roleArr               = [manageUsersRole]

    await addRoleMappingsToUser(adminToken, projectRealm, serviceAccount.id, realmManagementClient.id, roleArr)

    const clientToken = await loginAsClient(createdClient.secret)
    await addNewUser(clientToken, projectRealm, newStudentData);
    await addNewUser(clientToken, projectRealm, newTeacherData);
    await addNewUser(clientToken, projectRealm, newAdminData);
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

async function addNewUser(accessToken, realmName, user) {
    const createUserEndpoint = `${keycloakServer}/admin/realms/${realmName}/users`;

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

