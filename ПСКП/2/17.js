var http = require('http');

let http_handler = (req, res)=>{
    if(req.url === '/tests'){
        if(req.method === 'GET'){
            res.statusCode = 200;
            res.end('GET tests');
        }
        else if(req.method === 'POST'){
            res.statusCode = 201;
            res.end('POST test');
        }
        else if(req.method === 'PUT'){
            res.statusCode = 200;
            res.end('PUT tests');
        }
        else if(req.method === 'DELETE'){
            res.statusCode = 200;
            res.end('DELETE tests');
        }
        else{
            res.statusCode = 405;
            res.end('Method Not Allowed');
        }
    }
    else{
        res.statusCode = 404;
        res.end('404 Not Found');
    }
}

let server = http.createServer(http_handler);
server.listen(3000);