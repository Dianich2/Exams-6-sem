var http = require('http');

let http_handler = (req, res) =>{
    process.stdout.write(`Request: ${req.method} ${req.url}\n`);

    res.end('OK');
}

let server = http.createServer(http_handler);
server.listen(3000);

process.stdin.on('data', (data)=>{
    const command = data.toString().trim();

    if(command === 'stop'){
        process.stdout.write('Server stopped\n');
        server.close(()=>{
            process.exit(0);
        });
    }

});