const axios = require('axios');


const baseInstance = axios.create({
    baseURL: 'http://localhost:6333',
    headers: {
        'Content-Type': 'application/json',
        'api-key': process.env.API_KEY
    }
});

const run = async () => {
    const responseCluster = await baseInstance.get('/cluster')
    const currentPeerId = responseCluster.data.result.peer_id;

    const responseCollections = await baseInstance.get('/collections')
    const collections = responseCollections.data.result.collections.map(c => c.name);

    for (let collection of collections) {
        const responseCollection = await baseInstance.get(`/collections/${collection}`);
        // const status = responseCollection.data.result.status;
        // if (status !== "green") {
        //     console.log("Collection", collection, "is not ready");
        //     process.exit(1);
        // }

        const replication_factor = responseCollection.data.result.config.params.replication_factor;
        const shard_number = responseCollection.data.result.config.params.shard_number;

        const responseCollectionCluster = await baseInstance.get(`/collections/${collection}/cluster`);

        const shards = [...responseCollectionCluster.data.result.local_shards, ...responseCollectionCluster.data.result.remote_shards]
        for (shard of responseCollectionCluster.data.result.local_shards) {
            if (shard.state !== "Active") {
                console.log("Shard", shard.shard_id, "is not active locally for collection", collection);
                process.exit(1);
            }
        }
        // Loop through every shard number
        for (let i = 0; i < shard_number; i++) {
            const shardCount = shards.filter(s => s.shard_id === i && s.state === "Active").length;
            if (shardCount < replication_factor) {
                console.log("Collection", collection, " | Shard", i, "|", shardCount, "/", replication_factor, "replicas");
                process.exit(1);
            }
        }
    }
}

run();