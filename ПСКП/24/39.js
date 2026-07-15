var http = require('http');
var oracledb = require('oracledb');

oracledb.outFormat = oracledb.OUT_FORMAT_OBJECT;

const dbConfig = {
    username: 'cel_admin',
    password: '222',
    connectString: 'localhost:1521/CEL_PDB'
};

let http_handler = async (req, res)=>{
    if(req.method !== 'POST'){
        res.statusCode = 405;
        return res.end('Method Not Allowed');
    }

    if(req.url !== '/celebrities'){
        res.statusCode = 404;
        return res.end('404 Not Found');
    }

    let body = '';

    req.on('data', (chunk)=>{
        body += chunk;
    });

    req.on('end', async ()=>{
        let conn;

        try{
            const data = JSON.parse(body);
            conn = await oracledb.getConnection(dbConfig);

            conn.execute(
                `insert into celebrities(fullname, nationality, reqphotopath)
                 values(:fullname, :nationality, :reqphotopath)
                `,
                {
                    fullname: data.fullname,
                    nationality: data.nationality,
                    reqphotopath: data.reqphotopath
                },
                {
                    autoCommit:true
                }
            );

            res.statusCode = 201;
            res.setHeader('Content-Type', 'application/json');

            res.end(JSON.stringify({
                message: 'Celebrity created'
            }));

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
    })
    
}

let server = http.createServer(http_handler);
server.listen(3000);