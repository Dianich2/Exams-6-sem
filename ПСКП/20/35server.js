const websocket = require('ws');

const server = new websocket.Server({port:3000});

server.on('connection', (socket)=>{
    console.log('Client connected');

    socket.send('Hello from server');

    socket.on('message', (message)=>{
        console.log(`From client: ${message.toString()}`);

        server.clients.forEach((client) =>{
            if(client.readyState === WebSocket.OPEN){
                client.send(`Broadcast: ${message.toString()}`);
            }
        })
    })

    socket.on('close', ()=>{
        console.log('Client disconnected');
    });
})