import { ApolloServer } from '@apollo/server';
import { ApolloGateway, RemoteGraphQLDataSource } from '@apollo/gateway';
import { readFileSync } from 'fs';
import express from 'express';
import http from 'http';
import cors from 'cors';
import {expressMiddleware} from "@apollo/server/express4";
import * as bodyParser from "express";
import promBundle from 'express-prom-bundle';

const app = express();
const httpServer = http.createServer(app);
app.use(promBundle({
    includeMethod: true,
    includePath: true,
}));

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

await server.start();

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

app.use(
    '/',
    cors(),
    // 50mb is the limit that `startStandaloneServer` uses, but you may configure this to suit your needs
    bodyParser.json({ limit: '50mb' }),
    // expressMiddleware accepts the same arguments:
    // an Apollo Server instance and optional configuration options
    expressMiddleware(server, {
        context: async ({ req }) => ({ token: await getTokenForRequest(req) }),
    }),

);



await new Promise((resolve) => httpServer.listen(4000, () => resolve()));

console.log(`🚀  Server ready at localhost:${4000}`);