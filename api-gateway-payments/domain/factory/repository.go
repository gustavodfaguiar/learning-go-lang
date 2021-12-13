package factory

import "github.com/gustavodfaguiar/learning-go-lang/api-gateway-payments/domain/repository"

type RepositoryFactory interface {
	CreateTransicationRepository() repository.TransactionRepository
}
