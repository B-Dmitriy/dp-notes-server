### Todos-server (Server api)

**Todos**

* ```GET``` /todos


* ```POST``` /todos
    * payload:
        * title: string


* ```PUT``` /todos/{todoId}
    * params:
        * todoId: string
    * payload:
        * title: string
        * description: string


* ```DELETE``` /todos/{todoId}
    * params:
        * todoId: string

**Tasks**

* ```GET``` /tasks/{todoId}
    * params:
        * todoId: string


* ```POST``` /tasks/{todoId}
    * params:
        * todoId: string
    * payload:
        * title: string
        * deadline: Date


* ```PUT``` /tasks/{todoId}/{id}
    * params:
        * todoId: string
        * id: string
    * payload:
        * title: string
        * completed: boolean
        * deadline: Date
        * description: string[]


* ```DELETE``` /tasks/{todoId}/{id}
    * params:
        * todoId: string
        * id: string