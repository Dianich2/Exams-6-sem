var express = require('express');

const app = express();

app.get('/download', (req, res)=>{
    res.download('test.txt');
});

app.listen(3000);
