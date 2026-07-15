var http = require('http');

let http_handler = (req, res) => {
    if(req.url === '/'){
        res.statusCode = 200;
        res.end('Start page');
    }
    else{
        res.statusCode = 404;
        res.end('404 Not Found');
    };
}


let server = http.createServer(http_handler);

server.listen(3000);
