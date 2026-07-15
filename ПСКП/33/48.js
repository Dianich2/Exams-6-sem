var express = require('express');

const app = express();

app.get('/tests', (req, res)=>{
    const id = req.query.id;

    res.end(`Test id = ${id}`);
});

app.listen(3000);