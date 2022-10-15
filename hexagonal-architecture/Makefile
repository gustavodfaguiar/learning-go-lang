start:
	@docker-compose up

stop:
	@docker-compose down

build:
	@docker-compose build

open:
	@docker exec -it appproduct bash

test:
	@docker exec -it appproduct go test ./...

generate-mock:
	@mockgen -destination=application/mocks/application.go -source=application/product.go application
