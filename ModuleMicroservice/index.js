const keycloakServer = 'http://localhost:8888';
const realm = 'master';
const adminClientId = 'admin-cli';
const adminClientSecret = 'NrOhJkHCMnBIIOLfVJoH3QAhQToZnzkE';

// Token endpoint URL
const tokenEndpoint = `${keycloakServer}/realms/${realm}/protocol/openid-connect/token`;
const realmName = 'Test2';
const newClientData = {
    clientId: 'demo-client',
    enabled: true,
    standardFlowEnabled: false,
    implicitFlowEnabled: false,
    directAccessGrantsEnabled: false,
    serviceAccountsEnabled: true,
    authorizationServicesEnabled: true,
    // Add any other client configuration properties as needed
};

// Make the fetch request
async function loginAsAdminWithTheAdminCli() {
    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `grant_type=client_credentials&client_id=${adminClientId}&client_secret=${adminClientSecret}`,
    };

    const response = await fetch(tokenEndpoint, requestData);
    const data = await response.json();
    return data.access_token;
}


async function init() {
    const token = await loginAsAdminWithTheAdminCli()
    await createRealm(token, realmName)
    await createClientInRealm(token, realmName, newClientData);

    const serviceAccount        = await getServiceAccountUser(token, realmName, newClientData.clientId)
    const clients               = await getAllClientsInRealm(token, realmName)
    const realmManagementClient = clients.find(client => client.clientId === 'realm-management');
    const roles                 = await getAllClientRolesInRealm(token, realmName, realmManagementClient.id);
    const manageUsersRole       = roles.find(role => role.name === 'manage-realm');
    const roleArr               = [manageUsersRole]

    await addRoleMappingsToUser(token, realmName, serviceAccount.id, realmManagementClient.id, roleArr)
}

init().then(r => console.log('finished'));

async function createClientInRealm(accessToken, realmName, newClientData) {
    const adminApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients`;

    // Request parameters
    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(newClientData),
    };

    // Return a Promise for the created client data
    await fetch(adminApiEndpoint, requestData)
}

async function createRealm(accessToken, newRealmName) {
    const adminApiEndpoint = `${keycloakServer}/admin/realms`;

    // New realm data (adjust as needed)
    const newRealmData = {
        realm: newRealmName,
        enabled: true,
    };

    // Request parameters
    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(newRealmData),
    };

    // Return a Promise for the created realm data
    await fetch(adminApiEndpoint, requestData)
}

async function getServiceAccountUser(accessToken, realmName, clientId) {
    const userApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/users`;

    // Request parameters
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

    // Request parameters for adding role mappings
    const roleMappingRequestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(roleMappings),
    };

    const response = await fetch(userRoleMappingApiEndpoint, roleMappingRequestData);
    return await response.json();
}

async function getAllClientRolesInRealm(accessToken, realmName, clientId) {
    const clientRolesApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients/${clientId}/roles`;

    // Request parameters for getting all client roles
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

    // Request parameters for getting all clients
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
