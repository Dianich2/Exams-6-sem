var express = require('express');

const app = express();

app.use(express.urlencoded({extended: true}));

app.get('/test', (req, res)=>{
    res.json(req.body);
})

app.listen(3000);