const {Schema, model} = require('mongoose');

const Todo = new Schema({
    _id: Number,
    title: String,
    completed: Boolean
});

module.exports = model('Todo', Todo);