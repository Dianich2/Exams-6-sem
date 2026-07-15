const websocket = require('ws');

const server = new websocket.Server({port:3000});

server.on('connection', (socket)=>{
    console.log('Client connected');

    socket.on('message', (message)=>{
        const data = JSON.parse(message);

        console.log(data);

        const resp = {
            status: 'success',
            id: data.id,
            name: data.name
        };

        socket.send(JSON.stringify(resp));
    })
});