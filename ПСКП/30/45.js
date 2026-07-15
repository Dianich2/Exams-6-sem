const http = require('http');
const fs = require('fs');
const url = require('url');
const { sequelize, Faculty, 
    Pulpit, Teacher, Subject, 
    Auditorium, AuditoriumType } = require('./db');

(async () => {
    try {
        await sequelize.authenticate();
        console.log("sequelize initialized");

        const server = http.createServer(async (req, res) => {

            const parsed = url.parse(req.url, true);
            const method = req.method;
            const pathname = parsed.pathname;

    if (method === 'GET' && pathname === '/api/faculties') {
        try {
            const result = await Faculty.findAll({include:{model: Pulpit, include:[{model:Subject}, {model:Teacher}]}});
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(result));
        } catch (err) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: err.message }));
        }
        return;
    }

    if (method === 'GET' && pathname === '/api/pulpits') {
        try {
            const result = await Pulpit.findAll({include:[{model: Subject}, {model: Teacher}]});
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(result));
        } catch (err) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: err.message }));
        }
        return;
    }

    if (method === 'GET' && pathname === '/api/subjects') {
        try {
            const result = await Subject.findAll();
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(result));
        } catch (err) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: err.message }));
        }
        return;
    }

    if (method === 'GET' && pathname === '/api/auditoriumstypes') {
        try {
            const result = await AuditoriumType.findAll();
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(result));
        } catch (err) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: err.message }));
        }
        return;
    }

    if (method === 'GET' && pathname === '/api/auditoriums') {
        try {
            const result = await Auditorium.findAll();
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(result));
        } catch (err) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: err.message }));
        }
        return;
    }

    if (method === 'POST' && pathname === '/api/faculties') {
      let body = '';
      req.on('data', chunk => body += chunk);
      req.on('end', async () => {
          try {
              const data = JSON.parse(body);
              const result = await Faculty.create(data);
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify(result));
          } catch (e) {
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify({ error: e.message }));
          }
      });
      return;
    }

    if (method === 'POST' && pathname === '/api/pulpits') {
      let body = '';
      req.on('data', chunk => body += chunk);
      req.on('end', async () => {
          try {
              const data = JSON.parse(body);
              const result = await Pulpit.create(data);
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify(result));
          } catch (e) {
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify({ error: e.message }));
          }
      });
      return;
    }

    if (method === 'POST' && pathname === '/api/subjects') {
      let body = '';
      req.on('data', chunk => body += chunk);
      req.on('end', async () => {
          try {
              const data = JSON.parse(body);

              const result = await Subject.create(data);
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify(result));
          } catch (e) {
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify({ error: e.message }));
          }
      });
      return;
    }

    
    if (method === 'POST' && pathname === '/api/auditoriumstypes') {
      let body = '';
      req.on('data', chunk => body += chunk);
      req.on('end', async () => {
          try {
              const data = JSON.parse(body);

              const result = await AuditoriumType.create(data);
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify(result));
          } catch (e) {
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify({ error: e.message }));
          }
      });
      return;
    }

    if (method === 'POST' && pathname === '/api/auditoriums') {
      let body = '';
      req.on('data', chunk => body += chunk);
      req.on('end', async () => {
          try {
              const data = JSON.parse(body);

              const result = await Auditorium.create(data);
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify(result));
          } catch (e) {
              res.writeHead(200, { 'Content-Type': 'application/json' });
              res.end(JSON.stringify({ error: e.message }));
          }
      });
      return;
    }

    if (method === 'PUT' && pathname === '/api/faculties') {
        let body = '';
        req.on('data', chunk => body += chunk);
        req.on('end', async () => {
            try {
                const data = JSON.parse(body);

                const result = await Faculty.findByPk(data.faculty);

                if (!result) {
                    res.writeHead(404, { 'Content-Type': 'application/json' });
                    return res.end(JSON.stringify({ error: "Faculty not found" }));
                }
                await Faculty.update(data, {
                    where: { faculty: data.faculty }
                });
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify(data));
            } catch (e) {
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: e.message }));
            }
        });
        return;
    }

    if (method === 'PUT' && pathname === '/api/pulpits') {
        let body = '';
        req.on('data', chunk => body += chunk);
        req.on('end', async () => {
            try {
                const data = JSON.parse(body);

                const result = await Pulpit.findByPk(data.pulpit);

                if (!result) {
                    res.writeHead(404, { 'Content-Type': 'application/json' });
                    return res.end(JSON.stringify({ error: "Pulpit not found" }));
                }
                await Pulpit.update(data, {
                    where: { pulpit: data.pulpit }
                });
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify(data));
            } catch (e) {
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: e.message }));
            }
        });
        return;
    }

    if (method === 'PUT' && pathname === '/api/subjects') {
        let body = '';
        req.on('data', chunk => body += chunk);
        req.on('end', async () => {
            try {
                const data = JSON.parse(body);

                const result = await Subject.findByPk(data.subject);

                if (!result) {
                    res.writeHead(404, { 'Content-Type': 'application/json' });
                    return res.end(JSON.stringify({ error: "Subject not found" }));
                }
                await Subject.update(data, {
                    where: { subject: data.subject }
                });

                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify(data));
            } catch (e) {
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: e.message }));
            }
        });
        return;
    }

    if (method === 'PUT' && pathname === '/api/auditoriumstypes') {
        let body = '';
        req.on('data', chunk => body += chunk);
        req.on('end', async () => {
            try {
                const data = JSON.parse(body);

                const result = await AuditoriumType.findByPk(data.auditorium_type);

                if (!result) {
                    res.writeHead(404, { 'Content-Type': 'application/json' });
                    return res.end(JSON.stringify({ error: "Auditorium_type not found" }));
                }
                await AuditoriumType.update(data, {
                    where: { auditorium_type: data.auditorium_type }
                });

                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify(data));
            } catch (e) {
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: e.message }));
            }
        });
        return;
    }

    if (method === 'PUT' && pathname === '/api/auditoriums') {
        let body = '';
        req.on('data', chunk => body += chunk);
        req.on('end', async () => {
            try {
                const data = JSON.parse(body);

                const result = await Auditorium.findByPk(data.auditorium);

                if (!result) {
                    res.writeHead(404, { 'Content-Type': 'application/json' });
                    return res.end(JSON.stringify({ error: "Auditorium not found" }));
                }
                await Auditorium.update(data, {
                    where: { auditorium: data.auditorium }
                });

                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify(data));
            } catch (e) {
                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: e.message }));
            }
        });
        return;
    }

    if (method === 'DELETE' && pathname.startsWith('/api/faculties/')) {
        const code = pathname.split('/')[3];

        try {
            const result = await Faculty.findByPk(code);

            if (!result) {
                res.writeHead(404, { 'Content-Type': 'application/json' });
                return res.end(JSON.stringify({ error: "Faculty not found" }));
            }
            
            const temp = result.get();
            await result.destroy();

            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(temp));
        } catch (e) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: e.message }));
        }
        return;
    }

    if (method === 'DELETE' && pathname.startsWith('/api/pulpits/')) {
        const code = pathname.split('/')[3];

        try {
            const result = await Pulpit.findByPk(code);

            if (!result) {
                res.writeHead(404, { 'Content-Type': 'application/json' });
                return res.end(JSON.stringify({ error: "Pulpit not found" }));
            }
            
            const temp = result.get();
            await result.destroy();

            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(temp));
        } catch (e) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: e.message }));
        }
        return;
    }

    if (method === 'DELETE' && pathname.startsWith('/api/subjects/')) {
        const code = pathname.split('/')[3];

        try {
            const result = await Subject.findByPk(code);

            if (!result) {
                res.writeHead(404, { 'Content-Type': 'application/json' });
                return res.end(JSON.stringify({ error: "Subject not found" }));
            }
            
            const temp = result.get();
            await result.destroy();

            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(temp));
        } catch (e) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: e.message }));
        }
        return;
    }

    if (method === 'DELETE' && pathname.startsWith('/api/auditoriumstypes/')) {
        const code = pathname.split('/')[3];

        try {
            const result = await AuditoriumType.findByPk(code);

            if (!result) {
                res.writeHead(404, { 'Content-Type': 'application/json' });
                return res.end(JSON.stringify({ error: "Auditorium_type not found" }));
            }
            
            const temp = result.get();
            await result.destroy();

            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(temp));
        } catch (e) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: e.message }));
        }
        return;
    }

    if (method === 'DELETE' && pathname.startsWith('/api/auditoriums/')) {
        const code = pathname.split('/')[3];

        try {
            const result = await Auditorium.findByPk(code);

            if (!result) {
                res.writeHead(404, { 'Content-Type': 'application/json' });
                return res.end(JSON.stringify({ error: "Auditorium not found" }));
            }
            
            const temp = result.get();
            await result.destroy();

            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(temp));
        } catch (e) {
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ error: e.message }));
        }
        return;
    } 
    res.writeHead(404, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ error: "Route not found" }));
});
    server.listen(4000, () => console.log("Server started on port 4000"));
}
catch (err) {
    console.error("INIT ERROR:", err);
}})();
