import { ApolloServer } from '@apollo/server';
import { startStandaloneServer } from '@apollo/server/standalone';
import { ApolloGateway } from '@apollo/gateway';
import { watch } from 'fs';
import { readFile } from 'fs/promises';

const server = new ApolloServer({
    gateway: new ApolloGateway({
        async supergraphSdl({ update, healthCheck }) {

            // create a file watcher
            const watcher = watch('./supergraph.graphql');

            // subscribe to file changes
            watcher.on('change', async () => {

                // update the supergraph schema
                try {
                    const updatedSupergraph = await readFile('./supergraph.graphql', 'utf-8');

                    // optional health check update to ensure our services are responsive
                    await healthCheck(updatedSupergraph);

                    // update the supergraph schema
                    update(updatedSupergraph);
                } catch (e) {

                    // handle errors that occur during health check or while updating the supergraph schema
                    console.error(e);
                }
            });

            return {
                supergraphSdl: await readFile('./supergraph.graphql', 'utf-8'),

                // cleanup is called when the gateway is stopped
                async cleanup() {
                    watcher.close();
                },
            };
        },
    }),
});

const { url } = await startStandaloneServer(server);
console.log(`🚀  Server ready at ${url}`);