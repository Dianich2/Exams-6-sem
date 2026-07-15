var http = require('http');
var oracledb = require('oracledb');
var url = require('url');

oracledb.outFormat = oracledb.OUT_FORMAT_OBJECT;

const dbConfig = {
    username: 'cel_admin',
    password: '222',
    connectString: 'localhost:1521/CEL_PDB'
}

let http_handler = async(req, res)=>{
    if(req.method !== 'DELETE'){
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }

    const parts = url.parse(req.url, true).path.split('/');

    if(parts[1] != 'celebrities'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }

    const id = parts[2];
    if(id == null){
        res.statusCode = 400;
        return res.end('Bad Request');
    }

    let conn;
    try{
        conn = await oracledb.getConnection(dbConfig);

        const result = await conn.execute(
            `delete from celebrities
             where id = :id`,
             {
                id: id
             },
             {
                autoCommit: true
             }
        );

        res.setHeader('Content-Type', 'application/json');

        if(result.rowsAffected === 0){
            res.statusCode = 404;
            return res.end(JSON.stringify(
                {err: 'Celebrity not found'}
            ));
        }

        res.statusCode = 200;
        res.end(JSON.stringify(
            {message: 'Celebrity deleted'}
        ));
    }
    catch(err){
        res.statusCode = 500;
        res.end(`Error: ${err.message}`);
    }
    finally{
        if(conn){
            await conn.close();
        }
    }
}

let server = http.createServer(http_handler);
server.listen(3000);