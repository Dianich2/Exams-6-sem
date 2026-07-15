var http = require('http');

const id = 2;

const options = {
    hostname: 'localhost',
    port: 3000,
    path: `/tests/${id}`,
    method: 'GET'
};

const req = http.request(options, (res)=>{
    let data = '';

    res.on('data', (chunk)=>{
        data += chunk;
    });

    res.on('end', ()=>{
        console.log(data);
    });
});

req.end();