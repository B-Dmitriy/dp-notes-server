import {Schema, model} from 'mongoose';

const Task = new Schema({
    _id: String,
    _todoId: String,
    title: String,
    completed: Boolean,
    created: Date,
    deadline: Date,
    description: [String]
});

export default model('Task', Task);