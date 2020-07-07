.PHONY: help

PORT=8081

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

welcome:


setup: ## Constroi a imagem docker
	@echo 'Running on http://localhost:$(PORT)'

start: ## Roda aplicação com docker

stop: ## Parar de rodar aplicação com docker

test: ## Roda os testes da aplicação

test-docker: ## Roda os testes dentro do docker
