# calc_service
Калькулятор арифметических выражений

# Установка и запуск
*все комманды ввводятся в git bash
(для удобства, комманды можно копировать и вставлять в git bash с помощью shift+insert)

```git clone https://github.com/egocentri/go-webcalc1.git```          - клонирование репозитория

```cd go-webcalc1```          - переход к репозиторию

```go run ./cmd/go-webcalc1/...```          - запуск сервиса

вот такой вывод вы должны получить
```Listening and serving HTTP on :8080```
# Если не запускается
Убедитесь, что модуль github.com/gin-gonic/gin установлен:

```go get -u github.com/gin-gonic/gin```
Если ошибка всё ещё появляется, выполните:
```
go clean -modcache
go mod tidy```

К сожалению иногда выполняется только через танцы с бубнами.

# Примеры использования 
*для ввода запросов необходимо открыть второе окно git bash

Успешный запрос:
```curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'```

Ответ:
```{"result": 6}```

Ошибка 422:
```curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "a"}'```

Ответ:
```{"error": ""}```
