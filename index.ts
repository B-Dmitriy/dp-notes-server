import {Request,Response} from 'express';

const express = require('express');
const mongoose = require('mongoose');
const cors = require('cors');

const PORT = process.env.PORT || 8088;
const todos = require('./routers/todos-router');

const app = express();

process.on('unhandledRejection', (err)=> {
    console.log(err);
});

app.use(cors());
app.use(express.json());

app.get('/', ((req: Request, res: Response) => {
    res.send('home');
}));

app.use('/todos', todos);
mongoose.connect(`mongodb+srv://admin:nkjGTJxiBG83fqr@todos-cluster.v8fim.mongodb.net/myFirstDatabase?retryWrites=true&w=majority`);

app.listen(PORT, () => {
    console.log(`Server started on ${PORT} port.`);
});
