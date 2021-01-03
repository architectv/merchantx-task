# Тестовое задание Merchant Experience (Асинхронная схема работы)

# Endpoints

- GET /api/offers/ - получение товаров
    - Параметры запроса:
        - id - идентификатор продавца,
        - offer_id - уникальный идентификатор товара в системе продавца,
        - substr - подстрока названия.
- GET /api/stats/:id - получение статистики задания
    - Параметры пути:
        - id - идентификатор задания.
- PUT /api/offers/ - добавление товаров
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
    "stat_id": 1
}
```

### 2. GET для _stat_id=1_

**Запрос:**
```
curl --location --request GET 'localhost:8001/api/stats/1'
```
**Тело ответа:**
```
{
    "status": "done",
    "create_count": 5,
    "update_count": 1,
    "delete_count": 1,
    "error_count": 5
}
```

### 3. GET для _id=1_

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

### 4. GET для _id=1_ и _substr=теле_

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
