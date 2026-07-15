var express = require('express');

const app = express();

app.get('/old301', (req, res)=>{
    res.redirect(
        301, 
        '/new'
    );
});

app.get('/old307', (req, res)=>{
    res.redirect(
        307,
        '/new'
    );
})

app.post('/old301P', (req, res)=>{
    res.redirect(
        301, 
        '/new'
    );
});

app.post('/old307P', (req, res)=>{
    res.redirect(
        307,
        '/new'
    );
})

app.get('/new', (req, res) => {

    res.send(
        'New page'
    );

});

app.post('/new', (req, res) => {

    res.send(
        'New page'
    );

});

app.listen(3000);