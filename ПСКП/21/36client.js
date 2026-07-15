const websocket = require('ws');

const socket = new websocket('ws://localhost:3000');

socket.on('open', ()=>{

    const message = {
        id: 1,
        name: 'Dianich'
    };

    socket.send(JSON.stringify(message));
})

socket.on('message', (message)=>{
    const data = JSON.parse(message);

    console.log(data);
});