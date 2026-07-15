const http = require('http');

const oracledb = require('oracledb');

oracledb.outFormat = oracledb.OUT_FORMAT_OBJECT;

const dbConfig = {
    user: 'cel_admin',
    password: '222',
    connectString: 'localhost:1521/CEL_PDB'
};

let http_handler = async (req, res)=>{
    if(req.method !== 'GET'){
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }

    if(req.url !== '/celebrities'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }

    let conn;

    try{
        conn = await oracledb.getConnection(dbConfig);

        const result = await conn.execute(
            `select id, fullname, nationality, reqphotopath from celebrities`
        );

        res.statusCode = 200;
        res.setHeader('Content-Type', 'application/json');
        
        res.end(JSON.stringify(result.rows));
    }
    catch(err){
        res.statusCode = 500;
        res.end(`Server error: ${err.message}`);
    }
    finally{
        if(conn){
            await conn.close();
        }
    }
}

let server = http.createServer(http_handler);
server.listen(3000);