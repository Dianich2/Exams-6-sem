const websocket = require('ws');

const server = new websocket.Server({port:3000});

server.on('connection', (socket)=>{
    console.log('Client connected');

    socket.send('Hello from server');

    socket.on('message', (message)=>{
        console.log(`From client: ${message.toString()}`)

        socket.send(`Server: ${message.toString()}`);
    })

    socket.on('pong', ()=>{
        console.log('Pong from client');
    })

    socket.on('close', ()=>{
        clearInterval(interval);
        console.log('Client disconnected');
    })

    const interval = setInterval(()=>{
        if(socket.readyState === websocket.OPEN){
            console.log('Server send ping');
            socket.ping()
        }
    }, 3000)
})

