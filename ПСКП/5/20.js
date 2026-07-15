var http = require('http')

let http_handler = (req, res)=>{
    if(req.method === 'POST' && req.url === '/tests'){
        let body = '';
        req.on('data', (chunk) =>{
            body += chunk;
        });

        req.on('end', ()=>{
            const data = JSON.parse(body);

            res.end(`id=${data.id}, name=${data.name}`);
        });

    }
    else if(req.method === 'POST'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }
    else{
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }
};

let server = http.createServer(http_handler);
server.listen(3000);