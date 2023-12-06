import { ApolloServer } from '@apollo/server';
import { startStandaloneServer } from '@apollo/server/standalone';
import { ApolloGateway, RemoteGraphQLDataSource } from '@apollo/gateway';
import { readFileSync } from 'fs';

const supergraphSdl = readFileSync('./supergraph-cloud.graphql').toString();

class AuthenticatedDataSource extends RemoteGraphQLDataSource {
    willSendRequest({ request, context }) {
        try {
            request.http.headers.set('Authorization', 'Bearer ' + context.token);
        } catch (error) {
            throw new Error('An error occurred while processing your request.');
        }
    }
}

const server = new ApolloServer({
    gateway: new ApolloGateway({
        supergraphSdl,
        buildService({ name, url }) {
            return new AuthenticatedDataSource({ url });
        },
    }),
    formatError: (error) => {
        return {
            message: error.message,
        };
    },
});

async function getTokenForRequest(req) {
    const authorizationHeader = req.headers.authorization;

    if (authorizationHeader) {
        const [bearer, token] = authorizationHeader.split(' ');

        if (bearer === 'Bearer' && token) {
            return token;
        }
    }

    return null;
}

const { url } = await startStandaloneServer(server, {
    context: async ({ req, res }) => ({
        token: await getTokenForRequest(req),
    }),
});
console.log(`🚀  Server ready at ${url}`);