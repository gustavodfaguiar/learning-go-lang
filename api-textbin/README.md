# TextBin

Create, get and delete notes.

### Describre architecture

- api-textbin
  - cmd
    - main.go
  - internal
    - backend
      - backend.go
    - db
      - db.go
  - it
    - integration_test.go
  - tests
    - unit_test.go
  - go.mod

### Commands

cache: docker run -p 6379:6379 --name some-redis -d redis

test: go test ./..

start: go run cmd/main.go

### Examples

POST

curl -i --header "Content-Type: application/json" --request POST --data '{"data" : "save this", "onetime" : false}' http://localhost:8080/api/note

onetime

curl -i --header "Content-Type: application/json" --request POST --data '{"data" : "save this one time", "onetime" : true}' http://localhost:8080/api/note


GET

curl -i  http://localhost:8080/api/note/75f62994-becd-4e09-bf42-145edd9e1d8c


### Pattern to create test

- Given
- When
- Then
