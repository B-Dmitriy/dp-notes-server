import Todo from '../models/Todo';
import { v4 } from 'uuid';

class TodosService {

    async getTodos() {
        return Todo.find();
    }

    async createTodo(title: string) {

        const todo = await new Todo({
            _id: v4(),
            title,
            description: '',
            tasks: []
        });

        await todo.save();

        return todo;
    }

    async updateTodo(_id: string, title: string, description: string) {

        await Todo.updateOne({_id}, {title, description});

        return {_id, title, description};
    }

    async deleteTodo(_id: string) {

        await Todo.deleteOne({_id});

        return {message: `Todo with id: ${_id} was been delete.`};
    }

}

export default new TodosService();