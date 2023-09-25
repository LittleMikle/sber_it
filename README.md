# Тестовое задание Сбер_it
## Сервис ToDo list

**Выполнены все условия:**

**1) Пагинация**

**2) Swagger**

**3) Тесты**

**4) Docker**

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

Сервис выполнен в соответствии с чистой архитектурой.

### Запуск:

**docker-compose-up**

Все стартует:

![Screenshot from 2023-09-23 23-07-49](https://github.com/LittleMikle/sber_it/assets/101155101/d47a0cf4-1073-4c4f-95e6-c8b6631edf5d)


После запуска применить создание таблиц из файла scheme/scheme.sql

### Swagger:

![Screenshot 2023-09-23 at 22-32-12 Swagger UI](https://github.com/LittleMikle/sber_it/assets/101155101/ef11be63-4e62-4eff-b366-c9381401dcce)


### Тесты:

**go test -v ./...**

![Screenshot from 2023-09-23 22-44-38](https://github.com/LittleMikle/sber_it/assets/101155101/13fff0da-1bc1-4e6c-b5ba-80867e435ac0)


## Postman коллекция лежит в корневой папке

Создание списка ToDo (дату необходимо указывать в международном формате ISO 8601 ГГГГ.ММ.ДД):

![Screenshot from 2023-09-23 22-46-27](https://github.com/LittleMikle/sber_it/assets/101155101/67cd8332-456a-4daa-836b-f0897b9cf739)

Список появился:

![Screenshot from 2023-09-23 22-47-29](https://github.com/LittleMikle/sber_it/assets/101155101/0cb470a7-ffc8-40ea-b012-70b9b0743113)


**Присутствует валидация входных данных.**


При отсутствии названия списка:

![Screenshot from 2023-09-23 23-10-27](https://github.com/LittleMikle/sber_it/assets/101155101/0f009416-c76b-4554-9000-c4b37a0bb494)

При отстутствии даты:

![Screenshot from 2023-09-23 23-10-54](https://github.com/LittleMikle/sber_it/assets/101155101/8890b39a-8937-401d-b828-e78f8e0668f3)

Изменение сегмента:

![Screenshot from 2023-09-23 22-49-12](https://github.com/LittleMikle/sber_it/assets/101155101/2a1f51fb-ab69-4281-a85b-a3bafa81b07c)


Список изменился:

![Screenshot from 2023-09-23 22-48-50](https://github.com/LittleMikle/sber_it/assets/101155101/36f71261-96ea-47a7-99d4-e10914a6bf71)

Пагинация выполнена через offset, в реальном проекте нужно использовать другой метод (можно использовать в качестве page token id последней записи, возвращая ответ фронту можно отдельной переменной возвращать id последней записи и фронт сам будет присылать токен при следующем запросе. В результате получится конструкция вида ) 

SELECT id, title, description, date, status FROM todo_lists WHERE id > token (который вернет фронт) 
ORDER BY id
LIMIT 3;

Для демонстрации пагинации добавим 5 новых записей (при пагинации выводит по 3 на страницу)

![Screenshot from 2023-09-23 22-53-34](https://github.com/LittleMikle/sber_it/assets/101155101/6551a617-77d5-4c16-9571-07093f8499e9)

Сделаем выборку по параметрам status и page (номер страницы, выводит по 3 записи (в таблице 6 записей, 5 из них со статусом "undone" => должно вывести 3 на 1 странице и 2 на 2))

![Screenshot from 2023-09-23 22-56-25](https://github.com/LittleMikle/sber_it/assets/101155101/a0ebe30e-f3b4-4bdb-9898-796adba0adc0)


2 страница

![Screenshot from 2023-09-23 22-56-38](https://github.com/LittleMikle/sber_it/assets/101155101/a6cde31d-8a6a-4dbe-b8b1-b8b6a1c36514)


Выборка по дате (дату необходимо указывать в международном формате ISO 8601 ГГГГ.ММ.ДД), изменим дату в некоторых записях:

![Screenshot from 2023-09-23 23-00-40](https://github.com/LittleMikle/sber_it/assets/101155101/8afb9716-910b-4415-9ee8-f422104082d9)

![Screenshot from 2023-09-23 23-00-59](https://github.com/LittleMikle/sber_it/assets/101155101/cb4b6291-83f4-430f-82dd-63f945eb0f9f)


Удаление списка, удалим список с id 1:

Было:

![Screenshot from 2023-09-23 23-04-19](https://github.com/LittleMikle/sber_it/assets/101155101/70619a2d-43ed-4469-a4c7-c07cbe43aeb3)

![Screenshot from 2023-09-23 23-04-46](https://github.com/LittleMikle/sber_it/assets/101155101/683ea15f-aa94-432d-ad12-aa401641b5d5)

Стало:

![Screenshot from 2023-09-23 23-05-06](https://github.com/LittleMikle/sber_it/assets/101155101/a11f37d4-f553-4c67-86c2-813d21b1a865)




