const websocket = require('ws');

const socket = new websocket('ws://localhost:3000');

socket.on('open', ()=>{
    console.log('Connected to server');

    socket.send('Hello from client');
})

socket.on('message', (message)=>{
    console.log(`From server: ${message.toString()}`);
});

socket.on('close', ()=>{
    console.log('Connection closed');
})

socket.on('error', (err)=>{
    console.log(`Error: ${err.message}`)
})


setTimeout(()=>{
    socket.close();
}, 10000)