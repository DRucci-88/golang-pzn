package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

// Buat Function Provider (Constructor)
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{
		// Error: true,
	}
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	}

	return &SimpleService{
		SimpleRepository: repository,
	}, nil
}
