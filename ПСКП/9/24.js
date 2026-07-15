var http = require('http');
var queryString = require('querystring');

let http_handler = (req, res)=>{

    if(req.method !== 'POST'){
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }

    if(req.url !== '/tests'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }

    let body = '';

    req.on('data', (chunk)=>{
        body += chunk;
    });

    req.on('end', ()=>{
        const params = queryString.parse(body);
        res.end(`name=${params.name}, age=${params.age}`);
    });
}

let server = http.createServer(http_handler);
server.listen(3000);
