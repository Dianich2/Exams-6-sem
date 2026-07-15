var express = require('express');

const app = express();

app.get('/tests/:id', (req, res)=>{
    const id = req.params.id;

    res.send(`Test id = ${id}`);
});

app.listen(3000);