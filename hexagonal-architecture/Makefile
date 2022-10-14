start:
	@docker-compose up -d

stop:
	@docker-compose down

open-container:
	@docker exec -it appproduct bash

test:
	@docker exec -it appproduct bash -- go test ./...

generate-mock:
	@mockgen -destination=application/mocks/application.go -source=application/product.go application
