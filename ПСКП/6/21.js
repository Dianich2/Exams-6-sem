var http = require('http');

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

    req.on('data', (chunk) =>{
        body += chunk;
    });

    req.on('end', ()=>{
        const reqData = JSON.parse(body);

        const respData = {
            id: reqData.id,
            name: reqData.name,
            status: 'success'
        };

        res.statusCode = 200;
        res.setHeader('Content-Type', 'application/json');
        res.end(JSON.stringify(respData));
    })
};

let server = http.createServer(http_handler);
server.listen(3000);