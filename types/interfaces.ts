export interface ITodo {
    _id: string,
    title: string,
    description: string,
    tasks: ITask[]
}

export interface ITask {
    _id: string,
    _todoId: string,
    title: string,
    completed: boolean,
    created: Date,
    deadline: Date,
    description: string[]
}