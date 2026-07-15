const http = require('http');
const url = require('url');
const redis = require('redis');

const redisClient = redis.createClient({
    socket: {
        host: 'localhost',
        port: 6379
    }
});

async function start() {

    await redisClient.connect();

    console.log('Redis connected');

    const server = http.createServer(async (req, res) => {

        const parsedUrl = url.parse(req.url, true);

        res.setHeader(
            'Content-Type',
            'application/json'
        );

        try {

            // READ

            if (
                req.method === 'GET' &&
                parsedUrl.pathname === '/faculty'
            ) {

                const key =
                    parsedUrl.query.key;

                const value =
                    await redisClient.get(key);

                return res.end(
                    JSON.stringify({
                        key,
                        value
                    })
                );
            }

            // CREATE

            if (
                req.method === 'POST' &&
                parsedUrl.pathname === '/faculty'
            ) {

                let body = '';

                req.on('data',
                    chunk => body += chunk
                );

                req.on('end',
                    async () => {

                        const data =
                            JSON.parse(body);

                        await redisClient.set(
                            data.key,
                            data.value
                        );

                        res.end(
                            JSON.stringify(data)
                        );
                    });

                return;
            }

            // UPDATE

            if (
                req.method === 'PUT' &&
                parsedUrl.pathname === '/faculty'
            ) {

                let body = '';

                req.on('data',
                    chunk => body += chunk
                );

                req.on('end',
                    async () => {

                        const data =
                            JSON.parse(body);

                        await redisClient.set(
                            data.key,
                            data.value
                        );

                        res.end(
                            JSON.stringify(data)
                        );
                    });

                return;
            }

            // DELETE

            if (
                req.method === 'DELETE' &&
                parsedUrl.pathname === '/faculty'
            ) {

                const key =
                    parsedUrl.query.key;

                await redisClient.del(key);

                return res.end(
                    JSON.stringify({
                        deleted: key
                    })
                );
            }

            res.statusCode = 404;
            res.end('Not Found');

        }
        catch (err) {

            res.statusCode = 500;

            res.end(
                JSON.stringify({
                    error: err.message
                })
            );
        }
    });

    server.listen(3000);

    console.log(
        'Server started on port 3000'
    );
}

start();