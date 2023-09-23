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

docker-compose-up

Все стартует:
![Screenshot from 2023-09-23 22-35-32](https://github.com/LittleMikle/sber_it/assets/101155101/76cafebe-6959-4ed3-b7c2-cac8622e4e65)

После запуска применить создание таблиц из файла scheme/scheme.sql

### Swagger:

![Screenshot 2023-09-23 at 22-32-12 Swagger UI](https://github.com/LittleMikle/sber_it/assets/101155101/dfecc032-20f0-4698-987c-dde8aaa7d8e3)


### Тесты:

go test -v ./...

![Screenshot from 2023-09-23 22-44-38](https://github.com/LittleMikle/sber_it/assets/101155101/0d1dc64b-32a4-4478-bdf6-a23bb7dfef23)

## Postman коллекция лежит в корневой папке

Создание списка ToDo (дату необходимо указывать в международном формате ISO 8601 ГГГГ.ММ.ДД):

![Screenshot from 2023-09-23 22-46-27](https://github.com/LittleMikle/sber_it/assets/101155101/c028f16a-0541-4a1f-9bf8-ccd13c3b0a48)

Список появился:

![Screenshot from 2023-09-23 22-47-29](https://github.com/LittleMikle/sber_it/assets/101155101/28b12066-1357-4109-a48f-11b5283e0765)

Изменение сегмента:

![Screenshot from 2023-09-23 22-49-12](https://github.com/LittleMikle/sber_it/assets/101155101/ad7e6ef1-72fd-4838-82fe-d9d6c56d8ca7)

Список изменился:

![Screenshot from 2023-09-23 22-48-50](https://github.com/LittleMikle/sber_it/assets/101155101/4bff62d4-28c8-430a-8d80-e0680f40faaf)

Для демонстрации пагинации добавим 5 новых записей (при пагинации выводит по 3 на страницу)

![Screenshot from 2023-09-23 22-53-34](https://github.com/LittleMikle/sber_it/assets/101155101/facecf9e-ff62-402e-9552-344b454c96d7)

Сделаем выборку по параметрам status и page (номер страницы, выводит по 3 записи)

![Screenshot from 2023-09-23 22-56-25](https://github.com/LittleMikle/sber_it/assets/101155101/4294328d-483d-4f87-b202-23c25f3ac065)

2 страница

![Screenshot from 2023-09-23 22-56-38](https://github.com/LittleMikle/sber_it/assets/101155101/de43c496-db16-4021-a53c-768a40dbf63c)

Выборка по дате (дату необходимо указывать в международном формате ISO 8601 ГГГГ.ММ.ДД):

![Screenshot from 2023-09-23 23-00-40](https://github.com/LittleMikle/sber_it/assets/101155101/00581ed1-d678-47f9-aa8f-9f8e158c1c8c)

![Screenshot from 2023-09-23 23-00-59](https://github.com/LittleMikle/sber_it/assets/101155101/5f23f5b7-22a5-41fd-8431-4fb2ca65010d)

Удаление списка, удалим список с id 1:

Было:

![Screenshot from 2023-09-23 23-04-19](https://github.com/LittleMikle/sber_it/assets/101155101/43db2af3-ad81-437a-a1ee-e29ab7e7c705)

![Screenshot from 2023-09-23 23-04-46](https://github.com/LittleMikle/sber_it/assets/101155101/2732f57c-e0c2-4929-89e7-1910aa63f952)

Стало:

![Screenshot from 2023-09-23 23-05-06](https://github.com/LittleMikle/sber_it/assets/101155101/5c8d07db-b055-496e-9c28-fc8b7fb39b26)



