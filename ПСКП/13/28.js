var http = require('http');

const body = JSON.stringify(
    {
        id: 2,
        name: 'Dianich'
    }
);

const options = {
    hostname: 'localhost',
    port: 3000,
    path: '/tests',
    method: 'POST',
    headers:{
        'Content-Type': 'application/json',
        'Content-Length': Buffer.byteLength(body)
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
})

req.write(body);
req.end();