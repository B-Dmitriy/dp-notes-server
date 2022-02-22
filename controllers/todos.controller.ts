import {Request, Response} from 'express';
import TodosService from '../service/todos.service';

class TodosController {

    async getTodos(req: Request, res: Response) {

        const result = await TodosService.getTodos();

        if (result) {
            res
                .status(200)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Data missing from database');
            return;
        }
    }

    async createTodo(req: Request, res: Response) {

        const {title} = req.body;

        if (title === undefined) {
            res
                .status(400)
                .send('Request body missing property title');
            return;
        } else if (title === '') {
            res
                .status(400)
                .send('Title must contain at least 1 character ');
            return;
        }

        const result = await TodosService.createTodo(title);

        if (result) {
            res
                .status(201)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Error creating to-do list');
            return;
        }
    }

    async updateTodo(req: Request, res: Response) {

        const _id = req.params.id;

        const {title, description} = req.body;

        if (_id === undefined) {
            res
                .status(400)
                .send('_id parameter missing');
            return;
        } else if (title === undefined) {
            res
                .status(400)
                .send('Request body missing property title');
            return;
        } else if (description === undefined) {
            res
                .status(400)
                .send('Request body missing property description');
            return;
        }

        const result = await TodosService.updateTodo(_id, title, description);

        if (result) {
            res
                .status(200)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Error updating to-do list');
            return;
        }
    }

    async deleteTodo(req: Request, res: Response) {

        const _id = req.params.id;

        if (_id === undefined) {
            res
                .status(400)
                .send('_id parameter missing');
            return;
        }

        const result = await TodosService.deleteTodo(_id);

        if (result) {
            res
                .status(204)
                .send(result);
            return;
        } else {
            res
                .status(500)
                .send('Error deleting to-do list');
            return;
        }
    }
}

export default new TodosController();