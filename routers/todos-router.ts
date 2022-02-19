import {Request, Response} from 'express';

const express = require('express');
const Todo = require('../models/Todo');

const router = express.Router();

router.get('/', async (req: Request, res: Response) => {

    const result = await Todo.find();

    res.send(result);
});

router.post('/', async (req: Request, res: Response) => {

    const {title} = req.body;

    const result = await new Todo({
        _id: new Date(),
        title,
        completed: false
    });

    await result.save();

    res.status(201).send(result);
});

router.put('/:id', async (req: Request, res: Response) => {

    const _id = +req.params.id;

    const {title, completed} = req.body;

    await Todo.updateOne({_id},{title, completed});

    res.status(200).send({_id,title,completed});
});

router.delete('/:id', async (req: Request, res: Response) => {

    const _id = +req.params.id;

    await Todo.deleteOne({_id});

    res.status(204).send({message: `Todo with id: ${_id} was been delete.`});
});

module.exports = router;