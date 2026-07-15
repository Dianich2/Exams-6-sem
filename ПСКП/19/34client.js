const websocket = require('ws');

const socket = new websocket('ws://localhost:3000');

socket.on('open', ()=>{
    console.log('Connected to server');

    socket.send('Hello from client');
})

socket.on('ping', ()=>{
    console.log('Ping from server');
})

socket.on('pong', ()=>{
    console.log('Pong from server');
})

socket.on('close', ()=>{
    console.log('Connection closed');
});

socket.on('error', (err)=>{
    console.log(`Error: ${err.message}`);
})

setTimeout(()=>{
    socket.close();
}, 10000)