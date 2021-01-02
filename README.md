# Тестовое задание Merchant Experience

[![Build Status](https://travis-ci.com/architectv/merchantx-task.svg?token=fKJi943rZDkZWG4i8spC&branch=main)](https://travis-ci.com/architectv/merchantx-task)

<!-- ToC start -->
# Содержание

1. [Описание задачи](#Описание-задачи)
1. [Реализация](#Реализация)
1. [Endpoints](#Endpoints)
1. [Запуск](#Запуск)
1. [Тестирование](#Тестирование)
1. [Нагрузочное тестирование](#Нагрузочное-тестирование)
1. [Примеры](#Примеры)
<!-- ToC end -->

# Описание задачи

Разработать сервис, через который продавцы смогут передавать товары пачками в формате excel (xlsx).

Сервис принимает на вход ссылку на файл и id продавца, к чьему аккаунту будут привязаны загружаемые товары.
Сервис читает файл, сохраняет либо обновляет товары в БД. Обновление будет происходить, если пара (id продавца, offer_id) уже есть в базе.
В ответ на запрос выдаёт краткую статистику: количество созданных товаров, обновлённых, удалённых и количество строк с ошибками (например, цена отрицательная либо вообще не явялется числом).

Для проверки работоспособности сервиса нужно реализовать метод, с помощью которого можно будет достать список товаров из базы.
Метод должен принимать на вход id продавца, offer_id, подстроку названия товара (по подстроке "теле" возвращаются и "телефоны", и "телевизоры"). Ни один параметр не является обязательным, все указанные параметры применяются через логический оператор "AND".

В каждой строке скачанного файла содержится отдельный товар.
Колонки в файле и соответствующие значения полей товара следующие:

* **offer_id** - уникальный идентификатор товара в системе продавца,
* **name** - название товара,
* **price** - цена в рублях,
* **quantity** - количество товара на складе продавца,
* **available** - true/false (в случае false продавец хочет удалить товар из базы).

# Реализация

- Следование дизайну REST API.
- Подход "Чистой Архитектуры" и техника внедрения зависимости.
- Работа с фреймворком [fiber](https://github.com/gofiber/fiber).
- Работа с БД Postgres с использованием библиотеки [sqlx](https://github.com/jmoiron/sqlx) и написанием SQL запросов.
- Конфигурация приложения - библиотека [viper](https://github.com/spf13/viper).
- Реализация Graceful Shutdown.
- Запуск из Docker.
- Юнит-тестирование уровней бизнес-логики и взаимодействия с БД классическим способом и с помощью моков - библиотеки [testify](https://github.com/stretchr/testify), [mock](https://github.com/golang/mock).
- Сквозное (E2E) тестирование - BDD фреймворк [goconvey](https://github.com/smartystreets/goconvey).
- Непрерывная интеграция, запуск тестов в Travis CI.

**Структура проекта:**
```
.
├── pkg
│   ├── model       // основные структуры
│   ├── handler     // обработчики запросов
│   ├── service     // бизнес-логика
│   └── repository  // взаимодействие с БД
├── cmd             // точка входа в приложение
├── scripts         // SQL файлы с миграциями
├── configs         // файлы конфигурации
├── data            // директория для загружаемых файлов
├── test            // инициализация тестовой БД и набор тестовых файлов
└── e2e_test.go     // сквозной тест
```

# Endpoints

- GET /api/offers/
    - Параметры запроса:
        - id - идентификатор продавца,
        - offer_id - уникальный идентификатор товара в системе продавца,
        - substr - подстрока названия.
- PUT /api/offers/
    - Тело запроса:
        - id - идентификатор продавца,
        - link - ссылка на файл excel (xlsx).

# Запуск

```
make build
make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrateup
```

# Тестирование

Локальный запуск тестов:
```
make run_test
```

Для локального запуска тестов необходимо создать тестовую БД.
Это можно сделать, например, с помощью утилиты psql:
```
psql -c 'create database postgres_test;' -U postgres
psql -c "alter user postgres with password '1234';" -U postgres
```
# Нагрузочное тестирование

Нагрузочное тестирование проведено с помощью утилиты Apache Benchmark.
Результаты представлены в файле [ab_results.md](./ab_results.md)

# Примеры

[Ссылка на загрузку тестового файла.](https://docs.google.com/spreadsheets/d/e/2PACX-1vRmOaivfZYZqJCgnS6Dnjw8kLvRtgMELipP9r7m8nE_Te6N06glcNaGyNVw73f0VuKi8mgoErSploTZ/pub?output=xlsx)

Запросы сгенерированы из Postman для cURL.

### 1. PUT для _id=1_

**Запрос:**
```
curl --location --request PUT 'localhost:8001/api/offers/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "link": "https://docs.google.com/spreadsheets/d/e/2PACX-1vRmOaivfZYZqJCgnS6Dnjw8kLvRtgMELipP9r7m8nE_Te6N06glcNaGyNVw73f0VuKi8mgoErSploTZ/pub?output=xlsx"
}'
```
**Тело ответа:**
```
{
    "create_count": 5,
    "update_count": 1,
    "delete_count": 1,
    "error_count": 5
}
```

### 2. GET для _id=1_

**Запрос:**
```
$ curl GET localhost:8001/api/offers/?id=1
```
**Тело ответа:**
```
[
    {
        "id": 1,
        "offer_id": 1,
        "name": "телевизор",
        "price": 2000,
        "quantity": 10
    },
    {
        "id": 1,
        "offer_id": 2,
        "name": "телефон",
        "price": 4500,
        "quantity": 3
    },
    {
        "id": 1,
        "offer_id": 3,
        "name": "ноутбук",
        "price": 57000,
        "quantity": 45
    },
    {
        "id": 1,
        "offer_id": 4,
        "name": "книга",
        "price": 900,
        "quantity": 80
    }
]
```

### 3. GET для _id=1_ и _substr=теле_

**Запрос:**
```
$ curl GET localhost:8001/api/offers/?id=1&substr=теле
```
**Тело ответа:**
```
[
    {
        "id": 1,
        "offer_id": 1,
        "name": "телевизор",
        "price": 2000,
        "quantity": 10
    },
    {
        "id": 1,
        "offer_id": 2,
        "name": "телефон",
        "price": 4500,
        "quantity": 3
    }
]
```
