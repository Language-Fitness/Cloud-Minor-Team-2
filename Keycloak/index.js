import {
    newAdminAttributes,
    newAdminData,
    newClientData,
    newStudentAttributes,
    newStudentData, newTeacherAttributes,
    newTeacherData, MicroserviceClients
} from "./generate_data_keycloak.js";

import * as fs from "fs";
import * as path from "path";


const keycloakServer        = 'http://localhost:8888';
const masterRealm           = 'master';
const projectRealm          = 'cloud-project';
const adminClientId         = 'admin-cli';
const adminClientSecret     = 'deCEIv8Slq5GUeSWUZWShRvSyd5VhGsA';
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

    for (const client of MicroserviceClients) {
        await (async (client) => {
            await createClientInRealm(adminToken, projectRealm, client);
        })(client);
    }

    // Get the created client and realm management client (to fetch role which is to be assigned to the created client.)
    const clients               = await getAllClientsInRealm(adminToken, projectRealm);
    const realmManagementClient = clients.find(client => client.clientId === 'realm-management');
    const createdClient         = clients.find(client => client.clientId === newClientData.clientId)
    const serviceAccount        = await getServiceAccountUser(adminToken, projectRealm, newClientData.clientId)

    // Get all roles in realm for realm-client, filter for manage-users role, and assign to service account of created client
    const realmManagementRoles  = await getAllClientRolesInRealm(adminToken, projectRealm, realmManagementClient.id);
    const manageUsersRole       = realmManagementRoles.find(role => role.name === 'manage-users');
    await addRoleMappingsToEntity(adminToken, projectRealm, 'users', serviceAccount.id, realmManagementClient.id, [manageUsersRole])

    // Create groups in the project realm
    await addNewGroup(adminToken, projectRealm, 'permissions_student')
    await addNewGroup(adminToken, projectRealm, 'permissions_teacher')
    await addNewGroup(adminToken, projectRealm, 'permissions_administrator')

    // Fetch all groups to use to assign the roles to.
    const groups = await getGroups(adminToken, projectRealm);
    const studentGroup = groups.find(group => group.name === 'permissions_student')
    const teacherGroup = groups.find(group => group.name === 'permissions_teacher')
    const adminGroup = groups.find(group => group.name === 'permissions_administrator')

    // Loop through all files with roles, create them and assign to designated group
    try {
        const files = fs.readdirSync('./roles');
        for (const file of files) {
            const filePath = path.join('./roles', file);
            const contents = fs.readFileSync(filePath, 'utf8');
            const jsonArr = JSON.parse(contents)

            for (const item of jsonArr) {
                await createRole(adminToken, projectRealm, createdClient.id, item['name'])
                const createdRole = await getRoleByName(adminToken, projectRealm, createdClient.id, item['name'])

                let id;
                const group_section = file.split('_')[0];

                if (group_section === 'student') {
                    id = studentGroup.id;
                } else if (group_section === 'teacher') {
                    id = teacherGroup.id;
                } else {
                    id = adminGroup.id
                }

                await addRoleMappingsToEntity(
                    adminToken,
                    projectRealm,
                    'groups',
                    id,
                    createdClient.id,
                    [createdRole]
                )
            }
        }
    } catch (err) {
        console.error(err);
    }

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

async function getRoleByName(accessToken, realmName, clientId, roleName) {
    const getRoleEndpoint = `${keycloakServer}/admin/realms/${realmName}/clients/${clientId}/roles/${roleName}`;

    const requestData = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`,
        },
    };

    const response = await fetch(getRoleEndpoint, requestData);
    return await response.json();
}

