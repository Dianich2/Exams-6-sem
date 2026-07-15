var http = require('http');
var fs = require('fs');

const boundary = '----myBoundary';

const options = {
    hostname: 'localhost',
    port: 3000,
    path: '/upload',
    method: 'POST',
    headers:{
        'Content-Type':`multipart/form-data; boundary=${boundary}`
    }
}

const req = http.request(options, (res)=>{
    let data = '';

    res.on('data', (chunk)=>{
        data += chunk;
    })

    res.on('end', ()=>{
        console.log(data);
    })
});

req.write(
    `--${boundary}\r\n` +
    `Content-Disposition: form-data; name="file"; filename="test.txt"\r\n` +
    `Content-Type: text/plain\r\n\r\n`
);

const fileStream = fs.createReadStream('test.txt');

fileStream.pipe(req, {end:false});

fileStream.on('end', ()=>{
    req.end(`\r\n--${boundary}--\r\n`);
});
