# CalcService

## Описание
Этот проект представляет собой веб-сервис для вычисления арифметических выражений. Пользователь отправляет арифметическое выражение по HTTP и получает результат в ответ.

Сервис поддерживает следующие операции:
- Сложение: `+`
- Вычитание: `-`
- Умножение: `*`
- Деление: `/`
- Скобки для группировки выражений: `()`

---

### Инструкция по запуску проекта
1. Склонируйте проект с GitHub:
git clone https://github.com/Uneho/CalcService.git

2. cd calc_service

3. Запустите сервер:
go run ./cmd
---

## Примеры использования

### Успешный запрос:
Отправляем запрос с арифметическим выражением `2+2*2`.
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

### Ответ:
{
  "result": "6.00"
}


### Ошибка 422:
Этот код ошибки возвращается, если выражение содержит недопустимые символы (например, буквы или специальные символы, которые не являются операторами).

### Запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+a"
}'

### Ответ:
{
  "error": "Expression is not valid"
}

### Ошибка 500:
Этот код ошибки возвращается в случае внутренней ошибки сервера (например, деление на ноль).

### Запрос:
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2/(1-1)"
}'

### Ответ:
{
  "error": "Division by zero is not allowed"
}