package simple

type BarRepository struct {

}
// make provider-nya
func NewBarRepository() *BarRepository {
	return BarRepository{}	
}

type BarService struct{
	*BarRepository
}
// make provider-nya
func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{BarRepository: barRepository}
}