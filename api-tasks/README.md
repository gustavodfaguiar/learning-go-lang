## Api tasks

Cria uma api simples em golang

### Start

curl http://localhost:4000
```
go run web/main.go
```

GET
```bash
curl \
  -X GET \
  http://localhost:4000
```

GET
```bash
curl \
  -X GET \
  http://localhost:4000/task
```

POST
```bash
curl \
  -H 'Content-Type: application/json' \
  -X POST \
  -d '{"id": 123, "title": "Title", "description": "Description", "type": 1}' \
  http://localhost:4000
```
