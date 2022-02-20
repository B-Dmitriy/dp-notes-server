import {v4} from 'uuid';
import Task from '../models/Task';

class TasksService {

    async getTasks(_todoId: string) {

        return Task.find({_todoId});
    }

    async createTask(_todoId: string, title: string, deadline: Date) {

        const task = await new Task({
            _id: v4(),
            _todoId: _todoId,
            title: title,
            completed: false,
            created: new Date(),
            deadline: deadline,
            description: []
        });

        await task.save();

        return task;
    }

    async updateTask(_todoId: string,
                     _id: string,
                     title: string,
                     completed: boolean,
                     deadline: Date,
                     description: string) {

        await Task
            .find({_todoId})
            .updateOne({_id}, {title, completed, deadline, description});

        return {_id, title, description};

    }

    async deleteTask(_todoId: string, _id: string) {

        await Task
            .find({_todoId})
            .deleteOne({_id});

        return {message: `Tasks with id: ${_id} was been delete.`};
    }

}

export default new TasksService();