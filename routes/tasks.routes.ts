import * as express from 'express';
import TasksController from '../controllers/tasks.controllers';

const routes = express.Router();

routes.get('/:todoId', TasksController.getTasks);
routes.post('/:todoId', TasksController.createTask);
routes.put('/:todoId/:id', TasksController.updateTask);
routes.delete('/:todoId/:id', TasksController.deleteTask);

export default routes;