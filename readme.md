# Easy API with Go

This is program I developed as a lerning process before creating an API in Golang.

I hope it helps peple in the same situation as me.

## Overview

It is an API that operates MySql tables with CRUD. The data use JSON.

The table infomation is as follows.

```
+-------+-------------+------+-----+---------+----------------+
| Field | Type        | Null | Key | Default | Extra          |
+-------+-------------+------+-----+---------+----------------+
| id    | int         | NO   | PRI | NULL    | auto_increment |
| name  | varchar(64) | NO   |     | NULL    |                |
| email | char(30)    | NO   |     | NULL    |                |
+-------+-------------+------+-----+---------+----------------+
```

## Description

I'm using the following pakages.

- github.com/go-sql-driver/mysql v1.6.0
- github.com/jinzhu/gorm v1.9.16
- github.com/labstack/echo v3.3.10
- gopkg.in/ini.v1 v1.62.0
- github.com/swaggo/swag/cmd/swag
- github.com/swaggo/echo-swagger

## Requirement

- Docker

## Instarll

```
$ docker-compose build
```

```
$ docker-compose up -d
```

```
$ docker-compose exec app realize start
```

## Using

### GET

```
$ curl "http://localhost:8081/users"

[{"id":1,"name":"Jupiter","email":"aaa@example.com"},{"id":2,"name":"Charlotte","email":"bbb@example.com"},{"id":3,"name":"Wing","email":"bbb@example.com"}]
```

```
$ curl "http://localhost:8081/user?name=Jupiter"

[{"id":1,"name":"Jupiter","email":"aaa@example.com"}]
```

### POST

```
$ curl -X POST -H "Content-Type: application/json" -d '{"name":"heytaro", "email":"eee@example.com"}' localhost:8081/create

"OK"
```

### PUT

```
 $ curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "name":"testaro", "email":"fjso@example.com"}' localhost:8081/update

"OK"
```

### DELETE

```
$ curl -X DELETE -H "Content-Type: application/json" -d '{"id": 1}' localhost:8081/delete

"OK"
```

### Swagger (API design documentation)

Access the following in your browser.

[http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[Aki Unleash - Akio Yano](https://github.com/AkiUnleash)
