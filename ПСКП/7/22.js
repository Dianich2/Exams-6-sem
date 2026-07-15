var http = require('http');
var fs = require('fs');
var path = require('path');

let http_handler = (req, res)=>{

    if(req.method !== 'GET'){
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }

    if(req.url !== '/download'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }

    const filepath = path.join(__dirname, 'test.txt');

    res.statusCode = 200;
    res.setHeader('Content-Type', 'text/plain');
    res.setHeader('Content-Disposition', 'attachment; filename="downloaded_file.txt"');

    const stream = fs.createReadStream(filepath);

    stream.pipe(res);

    stream.on('error', () =>{
        res.statusCode = 500;
        return res.end('Internal Server Error');
    });
};

let server = http.createServer(http_handler);
server.listen(3000);