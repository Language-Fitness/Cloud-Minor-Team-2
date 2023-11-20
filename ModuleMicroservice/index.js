const keycloakServer = 'http://localhost:8888';
const realm = 'master';
const clientId = 'admin-cli';
const clientSecret = 'NrOhJkHCMnBIIOLfVJoH3QAhQToZnzkE';

// Token endpoint URL
const tokenEndpoint = `${keycloakServer}/realms/${realm}/protocol/openid-connect/token`;
const realmName = 'example-realm';
const newClientData = {
    clientId: 'student-client',
    enabled: true,
    standardFlowEnabled: false,
    implicitFlowEnabled: false,
    directAccessGrantsEnabled: false,
    serviceAccountsEnabled: true,
    authorizationServicesEnabled: true,
    // Add any other client configuration properties as needed
};

const roleMappings = [
    {
        id: 'dec25575-e443-4e34-8d00-f708c5d58d6b',
        name: 'manage-users',
    },
];

// Make the fetch request
function loginAsAdminWithTheAdminCli() {
    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `grant_type=client_credentials&client_id=${clientId}&client_secret=${clientSecret}`,
    };

    // Return a Promise for the access token
    return fetch(tokenEndpoint, requestData)
        .then(response => response.json())
        .then(data => {
            return data.access_token;
        })
        .catch(error => {
            // Handle errors
            console.error('Error:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

async function init() {
    const token = await loginAsAdminWithTheAdminCli()
    // await createRealm(token, realmName)
    // await createClientInRealm(token, realmName, newClientData);
    const serviceAccount = await getServiceAccountUser(token, realmName, newClientData.clientId)

    console.log(serviceAccount);

    // return;
    // await getAllRolesInRealm(token, realmName)
    const clients = await getAllClientsInRealm(token, realmName)

    const realmManagementClient = clients.find(client => client.clientId === 'realm-management');
    await getAllClientRolesInRealm(token, realmName, realmManagementClient.id)
    await addRoleMappingsToUser(token, realmName, serviceAccount.id, roleMappings)
    // await getAuthenticationConfigDescription(token, realmName)
    // await getRealmManagementAccessToken(token, realmName)

    // GET /admin/realms/{realm}/clients/{id}/client-secret
}

init().then(r => console.log('finished'));

function createClientInRealm(accessToken, realmName, newClientData) {
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
    return fetch(adminApiEndpoint, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to create client: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(createdClientData => {
            // Handle the response from the Keycloak Admin REST API
            console.log('Client Created:', createdClientData);
            return createdClientData;
        })
        .catch(error => {
            // Handle errors
            console.error('Error creating client:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

function createRealm(accessToken, newRealmName) {
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
    return fetch(adminApiEndpoint, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to create realm: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(createdRealmData => {
            // Handle the response from the Keycloak Admin REST API
            console.log('Realm Created:', createdRealmData);
            return createdRealmData;
        })
        .catch(error => {
            // Handle errors
            console.error('Error:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

function getRealmManagementAccessToken(masterAccessToken, targetRealmName) {
    const tokenEndpoint = `${keycloakServer}/realms/${targetRealmName}/protocol/openid-connect/token`;

    // Request parameters
    const requestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Authorization': `Bearer ${masterAccessToken}`,
        },
        body: `grant_type=client_credentials&client_id=realm-management`,
    };

    // Return a Promise for the access token
    return fetch(tokenEndpoint, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to obtain realm-management access token: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(data => {
            // Access token obtained successfully
            const accessToken = data.access_token;

            // Return the access token
            return accessToken;
        })
        .catch(error => {
            // Handle errors
            console.error('Error obtaining realm-management access token:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

function getAuthenticationConfigDescription(adminAccessToken, realmName) {
    const adminApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/authentication/per-client-config-description`;

    // Request parameters
    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${adminAccessToken}`,
        },
    };

    // Return a Promise for the response data
    return fetch(adminApiEndpoint, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to fetch authentication config description: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(responseData => {
            // Handle the response data as needed
            console.log('Authentication Config Description:', responseData);
            return responseData;
        })
        .catch(error => {
            // Handle errors
            console.error('Error fetching authentication config description:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

function getServiceAccountUser(accessToken, realmName, clientId) {
    const userApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/users`;

    // Request parameters
    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    // Return a Promise for the service account user data
    return fetch(`${userApiEndpoint}?username=service-account-${clientId}`, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to get service account user: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(serviceAccountUserData => {
            // Handle the response from the Keycloak Admin REST API
            console.log('Service Account User Data:', serviceAccountUserData);
            return serviceAccountUserData[0]; // Assuming there is only one user with the given username
        })
        .catch(error => {
            // Handle errors
            console.error('Error getting service account user:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

function addRoleMappingsToUser(accessToken, realmName, userId, roleMappings) {

    const userRoleMappingApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/users/${userId}/role-mappings/clients/${newClientData.clientId}`;
    // /admin/realms/{realm}/users/{id}/role-mappings/clients/{client}

    // Request parameters for adding role mappings
    const roleMappingRequestData = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
        body: JSON.stringify(roleMappings),
    };

    // Return a Promise for the response data
    return fetch(userRoleMappingApiEndpoint, roleMappingRequestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to add role mappings to user: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(responseData => {
            // Handle the response from the Keycloak Admin REST API
            console.log('Role Mappings Added:', responseData);
            return responseData;
        })
        .catch(error => {
            // Handle errors
            console.error('Error adding role mappings to user:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

function getAllRolesInRealm(accessToken, realmName) {
    const rolesApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/roles`;

    // Request parameters for getting all roles
    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    // Return a Promise for the roles data
    return fetch(rolesApiEndpoint, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to get roles in realm: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(rolesData => {
            // Handle the response from the Keycloak Admin REST API
            console.log('Roles in Realm:', rolesData);
            return rolesData;
        })
        .catch(error => {
            // Handle errors
            console.error('Error getting roles in realm:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

function getAllClientRolesInRealm(accessToken, realmName, clientId) {
    const clientRolesApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients/${clientId}/roles`;

    // Request parameters for getting all client roles
    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    // Return a Promise for the client roles data
    return fetch(clientRolesApiEndpoint, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to get client roles in realm: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(clientRolesData => {
            // Handle the response from the Keycloak Admin REST API
            console.log('Client Roles in Realm:', clientRolesData);
            return clientRolesData;
        })
        .catch(error => {
            // Handle errors
            console.error('Error getting client roles in realm:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

// Function to get all clients in a realm
function getAllClientsInRealm(accessToken, realmName) {
    const clientsApiEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients`;

    // Request parameters for getting all clients
    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    // Return a Promise for the clients data
    return fetch(clientsApiEndpoint, requestData)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Failed to get clients in realm: ${response.status} - ${response.statusText}`);
            }
            return response.json();
        })
        .then(clientsData => {
            // Handle the response from the Keycloak Admin REST API
            console.log('Clients in Realm:', clientsData);
            return clientsData;
        })
        .catch(error => {
            // Handle errors
            console.error('Error getting clients in realm:', error);
            throw error; // Re-throw the error to propagate it to the caller
        });
}

