var express = require('express');

const app = express();

app.use(express.json());

app.get('/test', (req, res)=>{
    const {id, name} = req.body;

    console.log(`id = ${id}, name = ${name}`);
    res.json({id: id, name: name, status: 'success'});
});

app.listen(3000);