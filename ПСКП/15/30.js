var http = require('http');
var fs = require('fs');

const file = fs.createWriteStream('downloaded_file.txt');

http.get('http://localhost:3000/download', (res)=>{
    res.pipe(file);

    file.on('finish', ()=>{
        file.close();
        console.log('File downloaded');
    });
});
