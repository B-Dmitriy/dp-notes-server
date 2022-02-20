import {Request, Response} from 'express';
import TodosService from '../service/todos.service';

class TodosController {

    async getTodos(req: Request, res: Response) {

        const result = await TodosService.getTodos();

        res
            .status(200)
            .send(result);
    }

    async createTodo(req: Request, res: Response) {

        const {title} = req.body;

        const result = await TodosService.createTodo(title);

        res
            .status(201)
            .send(result);
    }

    async updateTodo(req: Request, res: Response) {
        const _id = req.params.id;

        const {title, description} = req.body;

        const result = await TodosService.updateTodo(_id, title, description);

        res
            .status(200)
            .send(result);
    }

    async deleteTodo(req: Request, res: Response) {

        const _id = req.params.id;

        const result = await TodosService.deleteTodo(_id);
        console.log(result)
        res
            .status(204)
            .send(result);
    }

    async getTask(req: Request, res: Response) {

        const _id = req.params.id;

        const {title, description} = req.body;

        const result = await TodosService.updateTodo(_id, title, description);

        res
            .status(200)
            .send(result);
    }
}

export default new TodosController();