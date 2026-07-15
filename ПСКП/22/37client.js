const websocket = require('ws');

const socket = new websocket('ws://localhost:3000');

socket.on('open', ()=>{
    socket.send(JSON.stringify({
        method: 'sum',
        params: [2, 3],
        id: 1
    }))

    socket.send(JSON.stringify({
        method: 'hello',
        params: ['Dianich'],
        id: 2
    }));
})

socket.on('message', (message)=>{
    console.log(JSON.parse(message))
});