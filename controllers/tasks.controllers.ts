import {Request, Response} from 'express';
import TasksService from '../service/tasks.service';

class TasksController {

    async getTasks(req: Request, res: Response) {

        const _todoId = req.params.todoId;

        if (!_todoId) {
            res
                .status(400)
                .send('_todoId parameter missing');
            return;
        }

        const result = await TasksService.getTasks(_todoId);

        if (result) {
            res
                .status(200)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Tasks missing from database');
            return;
        }
    }

    async createTask(req: Request, res: Response) {

        const _todoId = req.params.todoId;

        const {title, deadline} = req.body;

        if (_todoId === undefined) {
            res
                .status(400)
                .send('_id parameter missing');
            return;
        } else if (title === undefined) {
            res
                .status(400)
                .send('Request body missing property title');
            return;
        } else if (title === '') {
            res
                .status(400)
                .send('Title must contain at least 1 character ');
            return;
        } else if (deadline === undefined) {
            res
                .status(400)
                .send('Request body missing property deadline');
            return;
        }

        const result = await TasksService.createTask(_todoId, title, deadline);

        if (result) {
            res
                .status(201)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Error creating task');
            return;
        }
    }

    async updateTask(req: Request, res: Response) {

        const {_todoId, _id} = req.params;

        const {title, completed, deadline, description} = req.body;

        if (_todoId === undefined) {
            res
                .status(400)
                .send('_todoId parameter missing');
            return;
        } else if (_id === undefined) {
            res
                .status(400)
                .send('_id parameter missing');
            return;
        } else if (title === undefined) {
            res
                .status(400)
                .send('Request body missing property title');
            return;
        } else if (completed === undefined) {
            res
                .status(400)
                .send('Request body missing property completed');
            return;
        } else if (deadline === undefined) {
            res
                .status(400)
                .send('Request body missing property deadline');
            return;
        } else if (description === undefined) {
            res
                .status(400)
                .send('Request body missing property description');
            return;
        }

        const result = await TasksService.updateTask(_todoId, _id, title, completed, deadline, description);

        if (result) {
            res
                .status(200)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Error updating tasks');
            return;
        }
    }

    async deleteTask(req: Request, res: Response) {

        const _id = req.params.id;
        const _todoId = req.params.todoId;

        if (_todoId === undefined) {
            res
                .status(400)
                .send('_todoId parameter missing');
            return;
        } else if (_id === undefined) {
            res
                .status(400)
                .send('_id parameter missing');
            return;
        }

        const result = await TasksService.deleteTask(_todoId, _id);

        if (result) {
            res
                .status(204)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Error deleting tasks');
            return;
        }
    }
}

export default new TasksController();