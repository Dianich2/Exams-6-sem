var express = require('express');
var cookieParser = require('cookie-parser');

const app = express();

app.use(cookieParser('test'));

app.get('/setBaseCookie', (req, res)=>{
    res.cookie(
        'username',
        'Dianich'
    );

    res.send('Set base cookie');
});

app.get('/setSignedCookie', (req, res)=>{
    res.cookie(
        'username',
        'Dianich',
        {
            signed: true
        }
    );

    res.send('Set signed cookie');
});

app.get('/getBaseCookie', (req, res)=>{
    res.json(req.cookies);
})

app.get('/getSignedCookie', (req, res)=>{
    res.json(req.signedCookies);
})

app.listen(3000);