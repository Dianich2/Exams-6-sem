const websocket = require('ws');

let server = new websocket.Server({port:3000});


server.on('connection', (socket)=>{
    console.log('client connected');

    socket.send('Hello from server');

    socket.on('message', (message)=>{
        console.log(`From client: ${message.toString()}`);

        socket.send(`Server: ${message}`);
    })

    server.on('close', ()=>{
        console.log('client disconnected');
    })
})



