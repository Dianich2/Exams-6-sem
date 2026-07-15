const express = require('express');

const app = express();
const PORT = 3000;

app.use(express.json());

app.use((req, res, next) => {
    console.log(`${req.method} ${req.url}`);
    next();
});

let users = [
    { id: 1, name: 'Diana' },
    { id: 2, name: 'Anna' }
];

app.get('/', (req, res) => {
    res.send('Express server is working');
});

app.get('/users', (req, res) => {
    res.json(users);
});

app.get('/users/:id', (req, res) => {
    const id = Number(req.params.id);

    const user = users.find(u => u.id === id);

    if (!user) {
        return res.status(404).json({
            message: 'User not found'
        });
    }

    res.json(user);
});

app.post('/users', (req, res) => {
    const name = req.body.name;

    if (!name) {
        return res.status(400).json({
            message: 'Name is required'
        });
    }

    const user = {
        id: users.length + 1,
        name: name
    };

    users.push(user);

    res.status(201).json(user);
});

app.put('/users/:id', (req, res) => {
    const id = Number(req.params.id);
    const name = req.body.name;

    const user = users.find(u => u.id === id);

    if (!user) {
        return res.status(404).json({
            message: 'User not found'
        });
    }

    user.name = name;

    res.json(user);
});

app.delete('/users/:id', (req, res) => {
    const id = Number(req.params.id);

    const index = users.findIndex(u => u.id === id);

    if (index === -1) {
        return res.status(404).json({
            message: 'User not found'
        });
    }

    users.splice(index, 1);

    res.json({
        message: 'User deleted'
    });
});

app.use((req, res) => {
    res.status(404).json({
        message: 'Route not found'
    });
});

app.listen(PORT, () => {
    console.log(`Server started on port ${PORT}`);
});