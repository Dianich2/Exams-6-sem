const websocket = require('ws');

const server = new websocket.Server({port:3000});

const methods = {
    sum: (a, b) => a + b,
    hello: (name) => `Hello, ${name}`
}

server.on('connection', (socket)=>{
    console.log('Client connected');

    socket.on('message', (message)=>{
        try{
            const req = JSON.parse(message);

            const method = methods[req.method];

            if(!method){
                return socket.send(JSON.stringify({
                    error: 'Method Not Found',
                    id: req.id
                }))
            }

            const res = method(...req.params);

            socket.send(JSON.stringify({
                result: res,
                id: req.id
            }))
        }
        catch(err){
            socket.send(JSON.stringify({
                error: 'Bad request',
                id: null
            }));
        }
    })
})


