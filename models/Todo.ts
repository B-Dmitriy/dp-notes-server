import {Schema, model} from 'mongoose';

const Todo = new Schema({
    _id: String,
    title: String,
    description: String
});

export default model('Todo', Todo);