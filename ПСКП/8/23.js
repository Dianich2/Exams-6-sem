var http = require('http');
var fs = require('fs');

let http_handler = (req, res)=>{
    if(req.method !== 'POST'){
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }

    if(req.url !== '/upload'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }

    let body = Buffer.alloc(0);

    req.on('data', (chunk)=>{
        body = Buffer.concat([body, chunk]);
    });

    req.on('end', ()=>{
        const contentType = req.headers['content-type'];

        if(!contentType || !contentType.includes('multipart/form-data')){
            res.statusCode = 400;
            return res.end('Bad Request');
        }

        const boundary = '--' + contentType.split('boundary=')[1];

        const parts = body.toString('binary').split(boundary);

        for(const part of parts){
            if(part.includes('filename=')){
                const filenameMatch = part.match(/filename="(.+)"/);
                const filename = filenameMatch[1];

                const fileStart = part.indexOf('\r\n\r\n') + 4;
                const fileEnd = part.lastIndexOf('\r\n');

                const fileContent = part.substring(fileStart, fileEnd);

                fs.writeFileSync(
                    './' + filename,
                    Buffer.from(fileContent, 'binary')
                );

                res.statusCode = 200;
                return res.end('File uploaded');
            }
        }

        res.statusCode = 400;
        res.end('Bad Request');
    });
}

let server = http.createServer(http_handler);
server.listen(3000);