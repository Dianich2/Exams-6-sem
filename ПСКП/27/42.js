var http = require('http');
var url = require('url');
const {MongoClient, ObjectId} = require('mongodb');

const uri = "mongodb://dianich:123321500Diana@ac-jugipkj-shard-00-00.m9iwtjy.mongodb.net:27017,ac-jugipkj-shard-00-01.m9iwtjy.mongodb.net:27017,ac-jugipkj-shard-00-02.m9iwtjy.mongodb.net:27017/?replicaSet=atlas-v1hcyt-shard-0&ssl=true&authSource=admin";

const client = new MongoClient(uri);

let db;

async function start() {
    try{
        await client.connect();
        console.log('Connected to MongoDB');

        db = client.db('BSTU');
        const server = http.createServer(async (req, res)=>{
            const parsedUrl = url.parse(req.url, true);
            const method = req.method;
            const pathname = parsedUrl.pathname;

            if(method === 'GET' && pathname === '/api/faculties'){
                try{
                    const docs = await db.collection("faculty").find().toArray();
                    res.writeHead(200, {"Content-Type": "application/json"});
                    return res.end(JSON.stringify(docs));
                }
                catch(err){
                    console.error(err);
                    res.writeHead(500);
                    return res.end('DB error');
                }
            }

            if(method === 'POST' && pathname === '/api/faculties'){
                let body = '';
                req.on('data', (chunk)=>{
                    body += chunk;
                })

                req.on('end', async() =>{
                    try{
                        const data = JSON.parse(body);
                        await db.collection("faculty").insertOne(data);

                        res.writeHead(200, {'Content-Type': 'application/json'});
                        return res.end(JSON.stringify(data));
                    }
                    catch(err){
                        console.error(`POST: ${err}`);
                        res.writeHead(500);
                        return res.end('DB error');
                    }
                });

                return;
            }

            if(method === 'PUT' && pathname === '/api/faculties'){
                let body = '';
                req.on('data', (chunk)=>{
                    body += chunk;
                })

                req.on('end', async()=>{
                    try{
                        const data = JSON.parse(body);
                        const mongoId = new ObjectId(data._id);
                        delete data._id;

                        await db.collection("faculty").findOneAndUpdate({_id: mongoId}, {$set: data});

                        res.writeHead(200, {"Content-Type": "application/json"});
                        return res.end(JSON.stringify({"_id": mongoId, ...data}));
                    }
                    catch(err){
                        console.error(`PUT error: ${err}`);
                        res.writeHead(500);
                        return res.end('DB error');
                    }
                })
                return;
            }

            if(method === 'DELETE' && pathname.startsWith('/api/faculties')){
                let code = pathname.split('/').pop();
                try{
                    const f = await db.collection("faculty").find({"faculty": code}).toArray();
                    await db.collection("faculty").deleteOne({faculty: code});

                    res.writeHead(200, {"Content-Type": "application/json"});
                    return res.end(JSON.stringify(f));
                }
                catch(err){
                    console.error(`DELETE error: ${err}`);
                    res.writeHead(500);
                    return res.end('DB error');
                }
                return;
            }
 
        });

        server.listen(4000);
    }
    catch(err){
        console.error(`Mongo connection error: ${err}`);
    }
}

start();



