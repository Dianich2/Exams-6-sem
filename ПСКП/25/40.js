var http = require('http');
var oracledb = require('oracledb');
var url = require('url');

oracledb.outFormat = oracledb.OUT_FORMAT_OBJECT;

const dbConfig = {
    username: 'cel_admin',
    password: '222',
    connectString: 'localhost:1521/CEL_PDB'
};

let http_handler = async(req, res)=>{
    if(req.method !== 'PUT'){
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }

   const parsedUrl = url.parse(req.url, true);

    const parts = parsedUrl.path.split('/');
    if(parts[1] !== 'celebrities'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }

    const id = parts[2];

    if(id == null){
        res.statusCode = 400;
        return res.end('Bad Request');
    }

    let body = '';

    req.on('data', (chunk)=>{
        body += chunk;
    });

    req.on('end', async()=>{
        let conn;

        try{
            const data = JSON.parse(body);
            conn = await oracledb.getConnection(dbConfig);

            const result = conn.execute(
                `update celebrities set
                 fullname = :fullname,
                 nationality = :nationality,
                 reqphotopath = :reqphotopath
                 where id = :id`,
                 {
                    id: id,
                    fullname: data.fullname,
                    nationality: data.nationality,
                    reqphotopath: data.reqphotopath
                 },
                 {
                    autoCommit: true
                 }
            );

            res.setHeader('Content-Type', 'application/json');
            
            if((await result).rowsAffected === 0){
                res.statusCode = 404;
                return res.end(JSON.stringify(
                    {err: 'Celebrity not found'}
                ));
            }

            res.statusCode = 200;
            return res.end(JSON.stringify(
                {message: 'Celebrity updated'}
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
    });
}

let server = http.createServer(http_handler);
server.listen(3000);