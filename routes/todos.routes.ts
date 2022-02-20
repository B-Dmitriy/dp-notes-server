import * as express from 'express';
import TodosController from '../controllers/todos.controller';

const routes = express.Router();

routes.get('/', TodosController.getTodos);
routes.post('/', TodosController.createTodo);
routes.put('/:id', TodosController.updateTodo);
routes.delete('/:id', TodosController.deleteTodo);

export default routes;