## Задание (Golang + PostgreSQL)
Написать сервис, который будет слушать входящие запросы по HTTP, 
преобразовывать их в запрос к соответствующей функции Postgres, 
выполнять запрос и возвращать ответ клиенту.

Дописать функции Postgres для сервиса.

/Скиллы: Golang, Postgres, regexp, строки,
работа с json в Golang и Postgres/

### 1. Web-сервис
- Написать сервис, который будет слушать входящие запросы по HTTP, преобразовывать их в запрос к соответствующей функции Postgres (по схеме трансляции, приведённой ниже), выполнять запрос и возвращать ответ клиенту.
  * Как плюс: ограничить максимальное количество одновременных коннектов к БД.
  * Как два плюса: добавить prometheus метрики на вызовы (количество вызовов, длительность выполнения).

- Настройки соединения с сервером Postgres читать из config файла:
port - (int) порт, на котором слушать запросы
endpoint - (string) название API
host - (string) hostname, где установлен Postgres
user - (string) имя пользователя Postgres
password - (string) пароль пользователя Postgres
schema - (string) схема в Postgres

- Трансляция запроса в вызов Postgres функции. Формат запроса к сервису:

HTTP_METHOD server:port/endpoint/vversion[/object/id ...]]/destination/[id] , где

HTTP_METHOD - одно из: GET, POST, PUT, DELETE
server - сервер, где запущен веб-сервис
port - порт
endpoint - значение из config-файла
version - номер версии API, число
/object/id - необязательный повторяющийся параметр, определяющий путь в иерархии объектов
/destination/ - конечный объект
id - id конечного объекта. Обязателен для методов PUT, DELETE, не указывается для POST. Для GET -- если указан, 
то возвращает элемент с данным id, если не указан, возвращает полный список элементов.
Правила трансляции
запрос в Postgres = select * from схема.[object1[_object2]...]_destination_method( [id1[, id2]... ,] id[, params])

В зависимости от HTTP метода к имени функции добавляется cуффикс method:

для GET - get
для POST - ins
для PUT - upd
для DELETE - del
В случае, если идентификатор объекта не указан, соответствующий элемент id в запросе должен быть равен нулю, на примерах:

для запроса GET http://localhost:80/api/v1/user/12/comment/34 
запрос в Postgres должен выглядеть так: select * from test.user_comment_get(12, 34) (комментарий c id=32 пользователя c id=12)

для запроса GET http://localhost:80/api/v1/user/12/comment/
запрос в Postgres должен выглядеть так: select * from test.user_comment_get(12, 0) (все комментарии пользователя 12)

для запроса GET http://localhost:80/api/v1/user/comment/
запрос в Postgres должен выглядеть так: select * from test.user_comment_get(0, 0) (все комментарии всех пользователей)

Для POST и PUT методов в теле запроса принимается JSON, который передаётся в Postgres в качестве параметра params.

Все методы должны возвращать результат работы соответствующей Postgres функции с ContentType = 'application/json'

### 2. PostgreSQL часть
Реализовать на стороне Postgres'а функции для работы с объектами
user : просмотр, добавление, редактирование, удаление (см. пример в sample.sql)
comment : просмотр, редактирование, удаление по id
user/XX/comment : просмотр комментариев пользователя XX, добавление комментария от пользователя XX

### 3. Примеры
* GET localhost:80/api/v1/user/34452
Транслируется в: select * from test.user_get(34452)
Физический смысл: Получить данные по пользователю 34452
Ответ сервиса: {"id":34452, "name":"Vasya", "email":"vasya@google"}
* GET localhost:80/api/v1/comment/456
Транслируется в: select * from test.comment_get(456)
Физический смысл: Получить комментарий с ID 456
Ответ сервиса: {"id":456, "id_user":34452, "txt":"My comment"}
* GET localhost:80/api/v1/user/34452/comment/
Транслируется в: select * from test.user_comment_get(34452, 0)
Физический смысл: Получить все комментарии пользователя 34452
Ответ сервиса: [{"id":456, "id_user":34452, "txt":"My comment"},{"id":460, "id_user":34452, "txt":"Foo!"}]
* GET localhost:80/api/v1/user/34452/comment/456
Транслируется в: select * from test.user_comment_get(34452, 456)
Физический смысл: Получить комментарий с ID 456 от пользователя 34452
Ответ сервиса: {"id":456, "id_user":34452, "txt":"My comment"}
* POST localhost:80/api/v1/user/34452/comment/
body: {"txt":"foo"}
Транслируется в: select * from test.user_comment_ins(34452, '{"txt":"foo"}')
Физический смысл: Создать комментарий от пользователя 34452
Ответ сервиса: {"id": 470}
* PUT localhost:80/api/v1/comment/460
body: {"txt":"bar"}
Транслируется в: select * from test.comment_upd(34452, 460, '{"txt":"bar"}')
Физический смысл: Изменить комментарий с ID 460
Ответ сервиса: {"id":460}
* DELETE localhost:80/api/v1/comment/460
Транслируется в: select * from test.user_comment_del(34452, 460)
Физический смысл: Удалить комментарий с ID 460
Ответ сервиса: {"id":460}

### sample.sql
```sql
create schema if not exists test;

create sequence if not exists test.seq_users;
create sequence if not exists test.seq_comments;

create table if not exists test.users
(
  id int not null default nextval('test.seq_users'::regclass),
  name varchar not null,
  email varchar not null,
  constraint "PK_users" primary key (id),
  constraint "UQ_users_email" unique (email),
  constraint "CHK_users_email" check (email like '%@%')
);

create table if not exists test.comments
(
  id int not null default nextval('test.seq_comments'::regclass),
  id_user int not null,
  txt varchar not null,
  constraint "PK_comments" primary key (id)
);


create or replace function test.user_get(_id integer) 
  returns json as
$BODY$
declare
  _ret json;
begin
  if _id = 0 then
    select array_to_json(array(
      select row_to_json(r)
      from (
        select u.id, u.name, u.email
        from test.users u
      ) r
    )) into _ret;
  else
    select row_to_json(r) into _ret
    from (
      select u.id, u.name, u.email
      from test.users u
      where id = _id
    ) r;
  end if;

  return _ret;

  exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;


create or replace function test.user_ins(_params json)
  returns json as
$BODY$
declare
  _newid integer;
begin
  _newid = 0;

  insert into test.users (name, email)
  select name, email
  from json_populate_record(null::test.users, _params)
  returning id into _newid;

  return json_build_object('id', _newid);

  exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;


create or replace function test.user_upd(_id integer, _params json)
  returns json as
$BODY$
begin
  update test.users set
    name = _params->>'name',
    email = _params->>'email'
  where id = _id;

  return json_build_object('id', _id);

  exception when others then

  return json_build_object('error', SQLERRM);
end
$BODY$
language plpgsql volatile cost 100;
