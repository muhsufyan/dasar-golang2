package simple

type FooRepository struct {
}

// make provider-nya
func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

// make provider-nya
func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{FooRepository: fooRepository}
}
