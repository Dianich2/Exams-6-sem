var http = require('http');

let http_handler = (req, res)=>{
    const parm = req.url.match(/^\/tests\/(\d+)$/);
    if(req.method === 'GET' && parm){
        let id = parm[1];

        res.statusCode = 200;
        res.end(`Test id = ${id}`);
    }
    else if(req.method === 'GET'){
        res.statusCode = 404;
        res.end('404 Not Found');
    }
    else{
        res.statusCode = 405;
        res.end('Method Not Allowed');
    }
};

let server = http.createServer(http_handler);
server.listen(3000);