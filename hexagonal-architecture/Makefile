start:
	@docker-compose up -d

test:
	@go test

generate-mock:
	@mockgen -destination=application/mocks/application.go -source=application/product.go application
