var http = require('http');
var url = require('url');

let http_handler = (req, res)=>{
    if(req.method === 'GET'){
        const parsedUrl = url.parse(req.url, true);
        if(parsedUrl.pathname !== '/tests'){
            res.statusCode = 404;
            return res.end('404 Not Found');
        }

        const id = parsedUrl.query.id;

         if(id == null){
            res.statusCode = 400;
            return res.end('Bad Request');
        }

        res.statusCode = 200;
        res.end(`Test id = ${id}`);
    }
    else{
        res.statusCode = 405;
        res.end('Method Not Allowed');
    }
};

let server = http.createServer(http_handler);
server.listen(3000);

