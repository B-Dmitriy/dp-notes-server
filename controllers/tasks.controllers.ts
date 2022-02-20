import {Request, Response} from 'express';
import TasksService from '../service/tasks.service';

class TasksController {

    async getTasks(req: Request, res: Response) {

        const _todoId = req.params.todoId;

        const result = await TasksService.getTasks(_todoId);

        res
            .status(200)
            .send(result);
    }

    async createTask(req: Request, res: Response) {

        const _todoId = req.params.todoId;

        const {title, deadline} = req.body;

        const result = await TasksService.createTask(_todoId, title, deadline);

        res
            .status(201)
            .send(result);
    }

    async updateTask(req: Request, res: Response) {

        const {_todoId, _id} = req.params;

        const {title, completed, deadline, description} = req.body;

        const result = await TasksService.updateTask(_todoId, _id, title, completed, deadline, description);

        res
            .status(200)
            .send(result);
    }

    async deleteTask(req: Request, res: Response) {

        const {_todoId, _id} = req.params;

        const result = await TasksService.deleteTask(_todoId, _id);

        res
            .status(204)
            .send(result);
    }

}

export default new TasksController();