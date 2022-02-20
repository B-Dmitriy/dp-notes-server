import * as express from 'express';
import * as config from 'config';
import * as mongoose from 'mongoose';
import * as cors from 'cors';
import todosRoutes from './routes/todos.routes';
import tasksRoutes from './routes/tasks.routes';

const PORT = config.get('port') || 8088;
const dbUrl = config.get('dbUrl');
const app = express();

process.on('unhandledRejection', (err) => console.log(err));

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({extended: true}));

app.use('/todos', todosRoutes);
app.use('/tasks', tasksRoutes);

mongoose
    .connect(dbUrl)
    .then(() => {
        console.log('Database connection established')
    });

app.listen(PORT, () => {
    console.log(`Server started on ${PORT} port.`);
});
